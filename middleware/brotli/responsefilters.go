package brotli

import (
	"net/http"
	"strings"
)

// 根据响应进行过滤
type ResponseHeaderFilter interface {
	ShouldCompress(header http.Header) bool
}

// 校验接口是否被实现
var (
	_ ResponseHeaderFilter = (*SkipCompressedFilter)(nil)
	_ ResponseHeaderFilter = (*ContentTypeFilter)(nil)
)

// SkipCompressedFilter judges whether content has been
// already compressed
type SkipCompressedFilter struct{}

// NewSkipCompressedFilter ...
func NewSkipCompressedFilter() *SkipCompressedFilter {
	return &SkipCompressedFilter{}
}

// ShouldCompress implements ResponseHeaderFilter interface
//
// Content-Encoding: https://tools.ietf.org/html/rfc2616#section-3.5
func (s *SkipCompressedFilter) ShouldCompress(header http.Header) bool {
	return header.Get("Content-Encoding") == "" && header.Get("Transfer-Encoding") == ""
}

// ContentTypeFilter 响应类型是否支持压缩
type ContentTypeFilter struct {
	contentType []string
}

// NewContentTypeFilter ...
func NewContentTypeFilter(types []string) *ContentTypeFilter {
	return &ContentTypeFilter{types}
}

// ShouldCompress implements RequestFilter interface
func (e *ContentTypeFilter) ShouldCompress(header http.Header) bool {
	contentType := header.Get("Content-Type")
	if contentType == "" {
		return false
	}

	for _, item := range e.contentType {
		if strings.Contains(contentType, item) {
			return true
		}
	}
	return false
}

// defaultContentType
var defaultContentType = []string{
	"text/plain",
	"application/x-javascript",
	"application/javascript",
	"application/json",
	"text/css",
	"application/xml",
}

// DefaultContentTypeFilter
func DefaultContentTypeFilter() *ContentTypeFilter {
	return NewContentTypeFilter(defaultContentType)
}
