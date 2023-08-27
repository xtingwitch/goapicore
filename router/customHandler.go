package router

import (
	coreHttp2 "github.com/xtingwitch/GoApiCore/coreHttp"
)

type CustomHandler func(request *coreHttp2.Request, response *coreHttp2.Response) *coreHttp2.Response
