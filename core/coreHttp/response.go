package coreHttp

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	http.ResponseWriter
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{
		ResponseWriter: w,
	}
}

func (r *Response) JSON(data interface{}, statusCode int) error {
	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(statusCode)

	return json.NewEncoder(r).Encode(data)
}

func (r *Response) SetHeader(key, value string) {
	r.Header().Set(key, value)
}

func (r *Response) SetStatusCode(statusCode int) {
	r.WriteHeader(statusCode)
}
