package routers

import (
	"leafmart/internal/handlers"
	"leafmart/internal/routers/middleware"
)

func (router *Mux) RegisterRoutes() {
	router.HandleFunc("GET", "/", handlers.HomePage, middleware.LoggerMiddleware)

	router.HandleFunc("GET", "/{id}", handlers.SecondPage)

	router.Route("/admin", func(r Router) {
		r.HandleFunc("GET", "/home", handlers.AdminPage)
	})
}
