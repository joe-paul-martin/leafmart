package routers

import (
	"leafmart/internal/handlers"
	"leafmart/internal/routers/middleware"
)

func (router *Router) RegisterRoutes() {
	router.HandleFunc("GET", "/", handlers.HomePage, middleware.LoggerMiddleware)
}
