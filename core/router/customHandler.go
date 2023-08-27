package router

import (
	// ... (other imports)
	"github.com/xtingwitch/GoApiCore/core/coreHttp"
)

type CustomHandler func(request *coreHttp.Request, response *coreHttp.Response) *coreHttp.Response
