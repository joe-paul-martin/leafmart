package logger

import (
	"leafmart/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLogger(config config.Config) {
	var logLevel zapcore.Level

	switch config.LogLevel {
	case "DEBUG":
		logLevel = zap.DebugLevel
	default:
		logLevel = zap.InfoLevel
	}

	loggerConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(logLevel),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, _ = loggerConfig.Build(zap.AddCaller())

	logger = logger.With(zap.String("service_name", config.ServiceName))
}

func Debug(msg string, args ...zapcore.Field) {
	logger.Debug(msg, args...)
}

func Info(msg string, args ...zapcore.Field) {
	logger.Info(msg, args...)
}

func Warn(msg string, args ...zapcore.Field) {
	logger.Warn(msg, args...)
}

func Panic(msg string, args ...zapcore.Field) {
	logger.Panic(msg, args...)
}

func Fatal(msg string, args ...zapcore.Field) {
	logger.Fatal(msg, args...)
}
