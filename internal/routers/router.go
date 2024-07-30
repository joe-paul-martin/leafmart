package routers

import (
	"context"
	"leafmart/internal/config"
	"leafmart/internal/routers/middleware"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Route struct {
	method          string
	pattern         string
	middlewareChain middleware.MiddlewareChain
	handler         http.HandlerFunc
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{}
}

// func (router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	for _, route := range router.routes {
// 		if req.Method == route.method && strings.HasPrefix(req.URL.Path, route.pattern) {
//
// 			route.handler.ServeHTTP(w, req)
// 			return
// 		}
// 	}
// 	http.NotFound(w, req)
// }

func (router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range router.routes {
		if req.Method == route.method && strings.HasPrefix(req.URL.Path, route.pattern) {
			ctx, cancel := context.WithTimeout(req.Context(), 2*time.Second)
			defer cancel()

			req = req.WithContext(ctx)

			handler := route.handler
			for i := len(route.middlewareChain) - 1; i >= 0; i-- {
				handler = route.middlewareChain[i](handler)
			}
			handler.ServeHTTP(w, req)
			return
		}
	}
	http.NotFound(w, req)
}

// func (router *Router) HandleFunc(method, pattern string, handler http.HandlerFunc) {
// 	router.routes = append(router.routes, Route{method: method, pattern: pattern, handler: handler})
// }

func (router *Router) HandleFunc(method, pattern string, handler http.HandlerFunc, middlewares ...middleware.Middleware) {

	route := Route{
		method:          method,
		pattern:         pattern,
		middlewareChain: middlewares,
		handler:         handler,
	}
	router.routes = append(router.routes, route)
}

func InitRouter(config config.Config) *Router {

	if config.Env == "DEBUG" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := NewRouter()
	router.RegisterRoutes()

	return router

}
