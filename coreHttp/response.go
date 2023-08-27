package coreHttp

import (
	"encoding/json"
	"github.com/xtingwitch/GoApiCore/responseModels"
	"net/http"
)

type Response struct {
	responseWriter http.ResponseWriter
	headers        Headers
	statusCode     int
	body           interface{}
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{
		responseWriter: w,
	}
}

func (r *Response) SetContent(data interface{}) {
	r.body = data
}

func (r *Response) SetHeader(key, value string) {
	r.headers.Set(key, value)
}

func (r *Response) SetStatusCode(statusCode int) {
	r.statusCode = statusCode
}

func (r *Response) GetBody() interface{} {
	return r.body
}

func (r *Response) GetHeaders() Headers {
	return r.headers
}

func (r *Response) GetRawHeaders() http.Header {
	return r.GetHeaders().rawHeaders
}

func (r *Response) GetStatusCode() int {
	return r.statusCode
}

func (r *Response) GetBodyBytes() ([]byte, error) {
	bytes, err := json.Marshal(r.GetBody())

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (r *Response) MessageResponse(message string) {
	responseData := responseModels.MessageResponse{}
	responseData.Data.Message = message

	r.body = responseData
	r.SetHeader("Content-Type", "application/vnd.api+json")
	r.SetStatusCode(http.StatusOK)
}
