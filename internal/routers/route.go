package routers

import (
	"fmt"
	"leafmart/internal/routers/middleware"
	"net/http"
	"strings"
)

type Routes interface {
	matchPattern(string, string) (Route, map[string]string)
}

type Route struct {
	method      string
	pattern     string
	middlewares middleware.Middleware
	handler     http.Handler
}

func (router *Mux) matchPattern(method, pattern string) (Route, map[string]string) {
	segments := splitPath(pattern)
	fmt.Println(segments)
outer:
	for _, route := range router.routes {
		params := make(map[string]string)
		routeSegments := splitPath(route.pattern)
		if route.method == method && len(segments) == len(routeSegments) {
			for i, segment := range segments {
				routeSegment := routeSegments[i]
				if isDynamicSegment(routeSegment) {
					paramName := getParamName(routeSegment)
					paramValue := segment
					params[paramName] = paramValue
				} else if segment != routeSegment {
					continue outer
				}
			}
			return route, params
		}
	}
	return Route{handler: http.HandlerFunc(http.NotFound)}, nil
}

func splitPath(pattern string) []string {
	return strings.Split(pattern, "/")
}

func isDynamicSegment(segment string) bool {
	if segment == "" {
		return false
	}
	return string(segment[0]) == "{" && string(segment[len(segment)-1]) == "}"
}

func getParamName(segment string) string {
	return segment[1 : len(segment)-1]
}
