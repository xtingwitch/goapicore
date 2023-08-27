package router

import (
	"github.com/xtingwitch/GoApiCore/core/coreHttp"
	"github.com/xtingwitch/GoApiCore/core/events"
	"net/http"
	"regexp"
)

type Router struct {
	routes []Route
}

func (r *Router) addRoute(method, pattern string, handler CustomHandler) {
	variables := extractVariables(pattern)

	route := Route{
		Method:    method,
		Pattern:   regexp.MustCompile(convertPatternToRegex(pattern)),
		Variables: variables,
		Handler:   handler,
	}

	r.routes = append(r.routes, route)
}

func (r *Router) Get(pattern string, handler func(request *coreHttp.Request, response *coreHttp.Response) *coreHttp.Response) {
	r.addRoute("GET", pattern, handler)
}

func (r *Router) Post(pattern string, handler func(request *coreHttp.Request, response *coreHttp.Response) *coreHttp.Response) {
	r.addRoute("POST", pattern, handler)
}

func (r *Router) Patch(pattern string, handler func(request *coreHttp.Request, response *coreHttp.Response) *coreHttp.Response) {
	r.addRoute("PATCH", pattern, handler)
}

func (r *Router) Put(pattern string, handler func(request *coreHttp.Request, response *coreHttp.Response) *coreHttp.Response) {
	r.addRoute("PUT", pattern, handler)
}

func (r *Router) Delete(pattern string, handler func(request *coreHttp.Request, response *coreHttp.Response) *coreHttp.Response) {
	r.addRoute("DELETE", pattern, handler)
}

func (r *Router) Head(pattern string, handler func(request *coreHttp.Request, response *coreHttp.Response) *coreHttp.Response) {
	r.addRoute("HEAD", pattern, handler)
}

func (r *Router) Options(pattern string, handler func(request *coreHttp.Request, response *coreHttp.Response) *coreHttp.Response) {
	r.addRoute("OPTIONS", pattern, handler)
}

func (r *Router) Connect(pattern string, handler func(request *coreHttp.Request, response *coreHttp.Response) *coreHttp.Response) {
	r.addRoute("CONNECT", pattern, handler)
}

func (r *Router) Trace(pattern string, handler func(request *coreHttp.Request, response *coreHttp.Response) *coreHttp.Response) {
	r.addRoute("TRACE", pattern, handler)
}

func (r *Router) fireEvent(eventType string, data interface{}) {
	eventBus := events.GetGlobalEventBus()
	eventBus.Broadcast(events.Event{Name: eventType, Data: data})
}

func (r *Router) ServeHTTP(response *coreHttp.Response, req *http.Request) {
	r.fireEvent("Before Routing", req)
	r.fireEvent("Before Request Construction", req)
	customReq := coreHttp.NewRequest(req, nil)
	r.fireEvent("After Request Construction", customReq)

	for _, route := range r.routes {
		if req.Method == route.Method {
			matches := route.Pattern.FindStringSubmatch(req.URL.Path)
			if len(matches) > 0 {
				r.fireEvent("Route Found", route)

				vars := make(map[string]string)

				for i, name := range route.Variables {
					vars[name] = matches[i+1]
				}

				req = contextWithVars(req, vars)
				customReq.SetVars(vars)

				r.fireEvent("Before Route Handling", customReq)
				route.Handler(customReq, response)
				r.fireEvent("After Route Handling", customReq)
				r.fireEvent("Route Complete", customReq)
				return
			}
		}
	}

	r.fireEvent("Route Not Found", customReq)
	r.fireEvent("After Routing", customReq)
}
