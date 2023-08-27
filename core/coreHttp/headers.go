package coreHttp

import (
	// ... (other imports)
	"net/http"
)

type Headers struct {
	rawHeaders http.Header
}

func NewHeaders(headers http.Header) *Headers {
	return &Headers{
		rawHeaders: headers,
	}
}

func (h *Headers) Get(key string, defaultValue interface{}) string {
	value := h.rawHeaders.Get(key)

	if value == "" && defaultValue != nil {
		return defaultValue.(string)
	}

	return value
}

func (h *Headers) Set(key, value string) {
	h.rawHeaders.Set(key, value)
}

func (h *Headers) Add(key, value string) {
	h.rawHeaders.Add(key, value)
}

func (h *Headers) Del(key string) {
	h.rawHeaders.Del(key)
}

func (h *Headers) All() http.Header {
	return h.rawHeaders
}

func (h *Headers) Clone() *Headers {
	clonedHeaders := make(http.Header)

	for key, values := range h.rawHeaders {
		clonedHeaders[key] = append([]string{}, values...)
	}

	return NewHeaders(clonedHeaders)
}

func (h *Headers) Exists(key string) bool {
	_, exists := h.rawHeaders[key]
	return exists
}
