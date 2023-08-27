package router

import (
	"context"
	"net/http"
	"regexp"
)

func extractVariables(pattern string) []string {
	r := regexp.MustCompile(`{([^}]+)}`)
	matches := r.FindAllStringSubmatch(pattern, -1)
	var variables []string
	for _, match := range matches {
		variables = append(variables, match[1])
	}
	return variables
}

func convertPatternToRegex(pattern string) string {
	return regexp.QuoteMeta(pattern)
}

func contextWithVars(req *http.Request, vars map[string]string) *http.Request {
	ctx := req.Context()
	for key, value := range vars {
		ctx = context.WithValue(ctx, key, value)
	}
	return req.WithContext(ctx)
}
