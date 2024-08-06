package routers

import (
	"context"
	"leafmart/internal/config"
	"leafmart/internal/routers/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Router interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
	HandleFunc(string, string, http.HandlerFunc, ...middleware.Middleware)
	Use(middleware.Middleware)
}

type Route struct {
	method      string
	pattern     string
	middlewares middleware.MiddlewareChain
	handler     http.HandlerFunc
}

type Mux struct {
	routes            []Route
	commonmiddlewares middleware.MiddlewareChain
}

func NewRouter() *Mux {
	return &Mux{}
}

func (router *Mux) Use(middleware middleware.Middleware) {
	router.commonmiddlewares = append(router.commonmiddlewares, middleware)
}

func (router *Mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	combinedMiddlewares := []middleware.Middleware{}
	if len(router.commonmiddlewares) > 0 {
		combinedMiddlewares = append(combinedMiddlewares, router.commonmiddlewares...)
	}

	route, params := router.matchPattern(req.Method, req.URL.Path)
	handler := route.handler
	ctx, cancel := context.WithTimeout(req.Context(), 2*time.Second)
	defer cancel()

	if len(params) > 0 {
		for key, value := range params {
			ctx = context.WithValue(ctx, key, value)
		}
	}

	combinedMiddlewares = append(combinedMiddlewares, route.middlewares...)
	for i := len(combinedMiddlewares) - 1; i >= 0; i-- {
		handler = combinedMiddlewares[i](handler)
	}

	req = req.WithContext(ctx)
	handler.ServeHTTP(w, req)
	return
}

func (router *Mux) HandleFunc(method, pattern string, handler http.HandlerFunc, middlewares ...middleware.Middleware) {

	route := Route{
		method:      method,
		pattern:     pattern,
		middlewares: middlewares,
		handler:     handler,
	}
	router.routes = append(router.routes, route)
}

func InitRouter(config config.Config) *Mux {

	if config.Env == "DEBUG" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := NewRouter()
	router.RegisterRoutes()

	return router

}
