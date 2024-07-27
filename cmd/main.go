package main

import (
	"leafmart/internal/config"
	"leafmart/internal/logger"
)

func main() {
	config, _ := config.SetupConfig()

	logger.InitLogger(config)

	logger.Info(config.ServiceName)
}
