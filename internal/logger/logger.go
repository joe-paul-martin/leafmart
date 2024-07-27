package logger

import (
	"leafmart/internal/config"

	"go.uber.org/zap"
)

var logger *zap.Logger

func InitLogger(config config.Config) {
	logger, _ = zap.NewProduction()
}

func Info(msg string) {
	logger.Info(msg)
}
