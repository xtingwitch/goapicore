// router/route.go

package router

import (
	"regexp"
)

type Route struct {
	Method    string
	Pattern   *regexp.Regexp
	Variables []string
	Handler   CustomHandler
}
