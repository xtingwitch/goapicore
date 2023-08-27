package router

import (
	"github.com/xtingwitch/GoApiCore/coreHttp"
)

type CustomHandler func(request *coreHttp.Request, response *coreHttp.Response) *coreHttp.Response
