package coreHttp

import (
	"net/http"
)

type Response struct {
	ResponseWriter http.ResponseWriter
	Headers        Headers
	StatusCode     int
	Body           interface{}
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{
		ResponseWriter: w,
	}
}

func (r *Response) JSON(data interface{}, statusCode int) {
	r.Headers.Set("Content-Type", "application/json")
	r.SetStatusCode(statusCode)
	r.SetContent(data)
}

func (r *Response) SetContent(data interface{}) {
	r.Body = data
}

func (r *Response) SetHeader(key, value string) {
	r.Headers.Set(key, value)
}

func (r *Response) SetStatusCode(statusCode int) {
	r.StatusCode = statusCode
}
