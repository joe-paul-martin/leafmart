package routers

import (
	"fmt"
	"net/http"
	"strings"
)

func (router *Mux) matchPattern(method, pattern string) (Route, map[string]interface{}) {
	segments := splitPath(pattern)
	fmt.Println(segments)
	for _, route := range router.routes {
		params := make(map[string]interface{})
		routeSegments := splitPath(route.pattern)
		if route.method == method && len(segments) == len(routeSegments) {
			for i, segment := range segments {
				routeSegment := routeSegments[i]
				if isDynamicSegment(routeSegment) {
					paramName := getParamName(routeSegment)
					paramValue := segment
					params[paramName] = paramValue
				} else if segment != routeSegment {
					break
				}
			}
			return route, params
		}
	}
	return Route{handler: http.NotFound}, nil
}

func splitPath(pattern string) []string {
	return strings.Split(pattern, "/")
}

func isDynamicSegment(segment string) bool {
	return string(segment[0]) == "{" && string(segment[len(segment)-1]) == "}"
}

func getParamName(segment string) string {
	return segment[1 : len(segment)-1]
}
