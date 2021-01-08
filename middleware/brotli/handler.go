package brotli

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"

	"github.com/andybalholm/brotli"
	"github.com/gin-gonic/gin"
)

const (
	BestSpeed          = brotli.BestSpeed
	BestCompression    = brotli.BestCompression
	DefaultCompression = brotli.DefaultCompression
	DefalutContentLen  = 1024
)

// Config is used in Handler initialization
type Config struct {
	// 压缩等级
	// 可选择的值: 0 => 11.
	// see https://pkg.go.dev/github.com/andybalholm/brotli#pkg-constants
	CompressionLevel int
	// 响应内容长度
	MinContentLength int64
	// 根据请求校验是否过滤
	RequestFilter []RequestFilter
	// 根据响应校验是否过滤
	ResponseHeaderFilter []ResponseHeaderFilter
}

// Handler implement brotli compression for gin
type Handler struct {
	compressionLevel     int
	minContentLength     int64
	requestFilter        []RequestFilter
	responseHeaderFilter []ResponseHeaderFilter
	brotliWriterPool     sync.Pool
	wrapperPool          sync.Pool
}

func NewHandler(config Config) *Handler {
	if config.CompressionLevel < BestSpeed || config.CompressionLevel > BestCompression {
		panic(fmt.Sprintf("brotli: invalid CompressionLevel: %d", config.CompressionLevel))
	}
	if config.MinContentLength <= 0 {
		panic(fmt.Sprintf("brotli: invalid MinContentLength: %d", config.MinContentLength))
	}

	handler := Handler{
		compressionLevel:     config.CompressionLevel,
		minContentLength:     config.MinContentLength,
		requestFilter:        config.RequestFilter,
		responseHeaderFilter: config.ResponseHeaderFilter,
	}

	// brotli writer
	handler.brotliWriterPool.New = func() interface{} {
		return brotli.NewWriterLevel(ioutil.Discard, handler.compressionLevel)
	}
	handler.wrapperPool.New = func() interface{} {
		return newWriterWrapper(handler.responseHeaderFilter, handler.minContentLength, nil, handler.getBrotliWriter, handler.putBrotliWriter)
	}

	return &handler
}

// 默认配置
var defaultConfig = Config{
	CompressionLevel: DefaultCompression,
	MinContentLength: DefalutContentLen,
	RequestFilter: []RequestFilter{
		NewCommonRequestFilter(),
	},
	ResponseHeaderFilter: []ResponseHeaderFilter{
		DefaultContentTypeFilter(),
	},
}

// DefaultHandler 创建一个默认handler
func DefaultHandler() *Handler {
	return NewHandler(defaultConfig)
}

// getBrotliWriter 获取一个brotli writer
func (h *Handler) getBrotliWriter() *brotli.Writer {
	return h.brotliWriterPool.Get().(*brotli.Writer)
}

// putBrotliWriter 回收BrotliWriter
func (h *Handler) putBrotliWriter(w *brotli.Writer) {
	if w == nil {
		return
	}

	_ = w.Close()
	w.Reset(ioutil.Discard)
	h.brotliWriterPool.Put(w)
}

func (h *Handler) getWriteWrapper() *writerWrapper {
	return h.wrapperPool.Get().(*writerWrapper)
}

func (h *Handler) putWriteWrapper(w *writerWrapper) {
	if w == nil {
		return
	}

	// 回收资源
	w.FinishWriting()
	w.OriginWriter = nil
	h.wrapperPool.Put(w)
}

type ginGzipWriter struct {
	wrapper      *writerWrapper
	originWriter gin.ResponseWriter
}

type ginBrotliWriter struct {
	wrapper      *writerWrapper
	originWriter gin.ResponseWriter
}

// 校验接口是否被实现
var _ gin.ResponseWriter = &ginBrotliWriter{}

func (g *ginBrotliWriter) WriteHeaderNow() {
	g.wrapper.WriteHeaderNow()
}

func (g *ginBrotliWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return g.originWriter.Hijack()
}

func (g *ginBrotliWriter) CloseNotify() <-chan bool {
	return g.originWriter.CloseNotify()
}

func (g *ginBrotliWriter) Status() int {
	return g.wrapper.Status()
}

func (g *ginBrotliWriter) Size() int {
	return g.wrapper.Size()
}

func (g *ginBrotliWriter) Written() bool {
	return g.wrapper.Written()
}

func (g *ginBrotliWriter) Pusher() http.Pusher {
	// TODO: not sure how to implement gzip for HTTP2
	return nil
}

// WriteString implements interface gin.ResponseWriter
func (g *ginBrotliWriter) WriteString(s string) (int, error) {
	return g.wrapper.Write([]byte(s))
}

// Write implements interface gin.ResponseWriter
func (g *ginBrotliWriter) Write(data []byte) (int, error) {
	return g.wrapper.Write(data)
}

// WriteHeader implements interface gin.ResponseWriter
func (g *ginBrotliWriter) WriteHeader(code int) {
	g.wrapper.WriteHeader(code)
}

// WriteHeader implements interface gin.ResponseWriter
func (g *ginBrotliWriter) Header() http.Header {
	return g.wrapper.Header()
}

// Flush implements http.Flusher
func (g *ginBrotliWriter) Flush() {
	g.wrapper.Flush()
}

// Gin implement gin's middleware
func (h *Handler) Gin(ctx *gin.Context) {
	var shouldCompress = true

	// 根据请求信息校验是否进行压缩
	for _, filter := range h.requestFilter {
		shouldCompress = filter.ShouldCompress(ctx.Request)
		if !shouldCompress {
			break
		}
	}

	if shouldCompress {
		wrapper := h.getWriteWrapper()
		wrapper.Reset(ctx.Writer)

		originWriter := ctx.Writer
		ctx.Writer = &ginBrotliWriter{
			originWriter: ctx.Writer,
			wrapper:      wrapper,
		}
		defer func() {
			// 资源回收
			h.putWriteWrapper(wrapper)
			ctx.Writer = originWriter
		}()
	}

	ctx.Next()
}
