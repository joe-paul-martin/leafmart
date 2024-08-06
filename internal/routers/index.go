package routers

import (
	"leafmart/internal/handlers"
	"leafmart/internal/routers/middleware"
)

func (router *Mux) RegisterRoutes() {
	router.HandleFunc("GET", "/", handlers.HomePage, middleware.LoggerMiddleware)
}
