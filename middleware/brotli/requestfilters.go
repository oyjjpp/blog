package brotli

import (
	"net/http"
	"strings"
)

// RequestFilter decide whether or not to compress response judging by request
type RequestFilter interface {
	// ShouldCompress decide whether or not to compress response,
	// judging by request
	ShouldCompress(req *http.Request) bool
}

// interface guards
var (
	_ RequestFilter = (*CommonRequestFilter)(nil)
)

// CommonRequestFilter judge via common easy criteria like
// http method, accept-encoding header, etc.
type CommonRequestFilter struct{}

// NewCommonRequestFilter ...
func NewCommonRequestFilter() *CommonRequestFilter {
	return &CommonRequestFilter{}
}

// ShouldCompress implements RequestFilter interface
func (c *CommonRequestFilter) ShouldCompress(req *http.Request) bool {
	return req.Method != http.MethodHead &&
		req.Method != http.MethodOptions &&
		req.Header.Get("Upgrade") == "" &&
		strings.Contains(req.Header.Get("Accept-Encoding"), "br")
}
