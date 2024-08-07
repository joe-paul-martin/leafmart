package main

import (
	"leafmart/internal/config"
	"leafmart/internal/logger"
	"leafmart/internal/routers"
	"net/http"
)

func main() {
	config, _ := config.SetupConfig()

	logger.InitLogger(config)

	logger.Info(config.ServiceName)

	router := routers.InitRouter(config)

	// router.Use(middleware.LoggerMiddleware)

	http.ListenAndServe("localhost:8080", router)

}
