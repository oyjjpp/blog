package brotli

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/andybalholm/brotli"
)

// writerWrapper wraps the originalHandler
// to test whether to brotli and brotli the body if applicable.
type writerWrapper struct {
	// header filter are applied by its sequence
	Filters []ResponseHeaderFilter
	// min content length to enable compress
	MinContentLength int64
	OriginWriter     http.ResponseWriter
	// use initBrotliWriter() to init brotliWriter when in need
	GetBrotliWriter func() *brotli.Writer
	// must close brotli writer and put it back to pool
	PutBrotliWriter func(*brotli.Writer)

	// internal below
	// *** WARNING ***
	// *writerWrapper.Reset() method must be updated
	// upon following field changing

	// compress or not
	// default to true
	shouldCompress bool
	// whether body is large enough
	bodyBigEnough bool
	// is header already flushed?
	headerFlushed         bool
	responseHeaderChecked bool
	statusCode            int
	// how many raw bytes has been written
	size         int
	brotliWriter *brotli.Writer
	bodyBuffer   []byte
}

// interface guard
var _ http.ResponseWriter = (*writerWrapper)(nil)
var _ http.Flusher = (*writerWrapper)(nil)

func newWriterWrapper(filters []ResponseHeaderFilter, minContentLength int64, originWriter http.ResponseWriter, getBrotliWriter func() *brotli.Writer, putBrotliWriter func(*brotli.Writer)) *writerWrapper {
	return &writerWrapper{
		shouldCompress:   true,
		bodyBuffer:       make([]byte, 0, minContentLength),
		Filters:          filters,
		MinContentLength: minContentLength,
		OriginWriter:     originWriter,
		GetBrotliWriter:  getBrotliWriter,
		PutBrotliWriter:  putBrotliWriter,
	}
}

// Reset the wrapper into a fresh one,
// writing to originWriter
func (w *writerWrapper) Reset(originWriter http.ResponseWriter) {
	w.OriginWriter = originWriter

	// internal below

	// reset status with caution
	// all internal fields should be taken good care
	w.shouldCompress = true
	w.headerFlushed = false
	w.responseHeaderChecked = false
	w.bodyBigEnough = false
	w.statusCode = 0
	w.size = 0

	if w.brotliWriter != nil {
		w.PutBrotliWriter(w.brotliWriter)
		w.brotliWriter = nil
	}
	if w.bodyBuffer != nil {
		w.bodyBuffer = w.bodyBuffer[:0]
	}
}

func (w *writerWrapper) Status() int {
	return w.statusCode
}

func (w *writerWrapper) Size() int {
	return w.size
}

func (w *writerWrapper) Written() bool {
	return w.headerFlushed || len(w.bodyBuffer) > 0
}

func (w *writerWrapper) WriteHeaderCalled() bool {
	return w.statusCode != 0
}

func (w *writerWrapper) initBrotliWriter() {
	w.brotliWriter = w.GetBrotliWriter()
	w.brotliWriter.Reset(w.OriginWriter)
}

// Header implements http.ResponseWriter
func (w *writerWrapper) Header() http.Header {
	return w.OriginWriter.Header()
}

// Write implements http.ResponseWriter
func (w *writerWrapper) Write(data []byte) (int, error) {
	w.size += len(data)

	if !w.WriteHeaderCalled() {
		w.WriteHeader(http.StatusOK)
	}

	if !w.shouldCompress {
		return w.OriginWriter.Write(data)
	}
	if w.bodyBigEnough {
		return w.brotliWriter.Write(data)
	}

	// fast check
	if !w.responseHeaderChecked {
		w.responseHeaderChecked = true

		header := w.Header()
		for _, filter := range w.Filters {
			w.shouldCompress = filter.ShouldCompress(header)
			if !w.shouldCompress {
				w.WriteHeaderNow()
				return w.OriginWriter.Write(data)
			}
		}

		if w.enoughContentLength() {
			w.bodyBigEnough = true
			w.WriteHeaderNow()
			w.initBrotliWriter()
			return w.brotliWriter.Write(data)
		}
	}

	if !w.writeBuffer(data) {
		w.bodyBigEnough = true

		// detect Content-Type if there's none
		if header := w.Header(); header.Get("Content-Type") == "" {
			header.Set("Content-Type", http.DetectContentType(w.bodyBuffer))
		}

		w.WriteHeaderNow()
		w.initBrotliWriter()
		if len(w.bodyBuffer) > 0 {
			written, err := w.brotliWriter.Write(w.bodyBuffer)
			if err != nil {
				err = fmt.Errorf("w.brotliWriter.Write: %w", err)
				return written, err
			}
		}
		return w.brotliWriter.Write(data)
	}

	return len(data), nil
}

func (w *writerWrapper) writeBuffer(data []byte) (fit bool) {
	if int64(len(data)+len(w.bodyBuffer)) > w.MinContentLength {
		return false
	}

	w.bodyBuffer = append(w.bodyBuffer, data...)
	return true
}

func (w *writerWrapper) enoughContentLength() bool {
	contentLength, err := strconv.ParseInt(w.Header().Get("Content-Length"), 10, 64)
	if err != nil {
		return false
	}
	if contentLength != 0 && contentLength >= w.MinContentLength {
		return true
	}

	return false
}

// WriteHeader implements http.ResponseWriter
//
// WriteHeader does not really calls originalHandler's WriteHeader,
// and the calling will actually be handler by WriteHeaderNow().
//
// http.ResponseWriter does not specify clearly whether permitting
// updating status code on second call to WriteHeader(), and it's
// conflicting between http and gin's implementation.
// Here, brotli consider second(and furthermore) calls to WriteHeader()
// valid. WriteHeader() is disabled after flushing header.
// Do note setting status code to 204 or 304 marks content uncompressable,
// and a later status code change does not revert this.
func (w *writerWrapper) WriteHeader(statusCode int) {
	if w.headerFlushed {
		return
	}

	w.statusCode = statusCode

	if !w.shouldCompress {
		return
	}

	if statusCode == http.StatusNoContent ||
		statusCode == http.StatusNotModified {
		w.shouldCompress = false
		return
	}
}

// WriteHeaderNow Forces to write the http header (status code + headers).
//
// WriteHeaderNow must always be called and called after
// WriteHeader() is called and
// w.shouldCompress is decided.
//
// This method is usually called by gin's AbortWithStatus()
func (w *writerWrapper) WriteHeaderNow() {
	if w.headerFlushed {
		return
	}

	// if neither WriteHeader() or Write() are called,
	// do nothing
	if !w.WriteHeaderCalled() {
		return
	}

	if w.shouldCompress {
		header := w.Header()
		header.Del("Content-Length")
		header.Set("Content-Encoding", "br")
		header.Add("Vary", "Accept-Encoding")
		originalEtag := w.Header().Get("ETag")
		if originalEtag != "" && !strings.HasPrefix(originalEtag, "W/") {
			w.Header().Set("ETag", "W/"+originalEtag)
		}
	}

	w.OriginWriter.WriteHeader(w.statusCode)

	w.headerFlushed = true
}

// FinishWriting flushes header and closed brotli writer
//
// Write() and WriteHeader() should not be called
// after FinishWriting()
func (w *writerWrapper) FinishWriting() {
	// still buffering
	if w.shouldCompress && !w.bodyBigEnough {
		w.shouldCompress = false
		w.WriteHeaderNow()
		if len(w.bodyBuffer) > 0 {
			_, _ = w.OriginWriter.Write(w.bodyBuffer)
		}
	}

	w.WriteHeaderNow()
	if w.brotliWriter != nil {
		w.PutBrotliWriter(w.brotliWriter)
		w.brotliWriter = nil
	}
}

// Flush implements http.Flusher
func (w *writerWrapper) Flush() {
	w.FinishWriting()

	if flusher, ok := w.OriginWriter.(http.Flusher); ok {
		flusher.Flush()
	}
}
