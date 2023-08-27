package coreHttp

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"
)

type Request struct {
	rawRequest *http.Request
	vars       map[string]string
	headers    *Headers
}

func NewRequest(r *http.Request, vars map[string]string) *Request {
	return &Request{
		rawRequest: r,
		vars:       vars,
		headers:    NewHeaders(r.Header),
	}
}

func (r *Request) SetVars(vars map[string]string) {
	r.vars = vars
}

func (r *Request) Method() string {
	return r.rawRequest.Method
}

func (r *Request) Path() string {
	return r.rawRequest.URL.Path
}

func (r *Request) IsMethod(method string) bool {
	return r.rawRequest.Method == method
}

func (r *Request) IsGet() bool {
	return r.IsMethod("GET")
}

func (r *Request) IsPost() bool {
	return r.IsMethod("POST")
}

func (r *Request) IsPut() bool {
	return r.IsMethod("PUT")
}

func (r *Request) IsDelete() bool {
	return r.IsMethod("DELETE")
}

func (r *Request) IsPatch() bool {
	return r.IsMethod("PATCH")
}

func (r *Request) IsOptions() bool {
	return r.IsMethod("OPTIONS")
}

func (r *Request) IsAjax() bool {
	return strings.ToLower(r.rawRequest.Header.Get("X-Requested-With")) == "xmlhttprequest"
}

func (r *Request) Param(key string) string {
	return r.vars[key]
}

func (r *Request) Params() map[string]string {
	return r.vars
}

func (r *Request) QueryParam(key string) string {
	return r.rawRequest.URL.Query().Get(key)
}

func (r *Request) QueryParams() map[string][]string {
	return r.rawRequest.URL.Query()
}

func (r *Request) Body() (string, error) {
	bodyBytes, err := io.ReadAll(r.rawRequest.Body)

	if err != nil {
		return "", err
	}

	defer func() {
		closeErr := r.rawRequest.Body.Close()
		if closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	return string(bodyBytes), err
}

func (r *Request) BodyAsJSON(v interface{}) error {
	bodyBytes, err := io.ReadAll(r.rawRequest.Body)
	if err != nil {
		return err
	}

	defer func() {
		closeErr := r.rawRequest.Body.Close()
		if closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	err = json.Unmarshal(bodyBytes, v)

	if err != nil {
		return err
	}

	return nil
}

func (r *Request) GetJSONField(fieldName string, result interface{}, defaultValue interface{}) error {
	var jsonData map[string]interface{}
	err := r.BodyAsJSON(&jsonData)
	if err != nil {
		return err
	}

	fieldValue, found := getNestedField(jsonData, strings.Split(fieldName, "."))
	if !found {
		reflect.ValueOf(result).Elem().Set(reflect.ValueOf(defaultValue))
		return nil
	}

	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(fieldValue))
	return nil
}

func getNestedField(data map[string]interface{}, fieldNames []string) (interface{}, bool) {
	if len(fieldNames) == 0 {
		return nil, false
	}

	fieldName := fieldNames[0]
	value, found := data[fieldName]
	if !found {
		return nil, false
	}

	if len(fieldNames) == 1 {
		return value, true
	}

	if nestedData, ok := value.(map[string]interface{}); ok {
		return getNestedField(nestedData, fieldNames[1:])
	}

	return nil, false
}
