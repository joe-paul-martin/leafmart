package middleware

import (
	"leafmart/internal/logger"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func LoggerMiddleware(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		logger.Info("starting the request", zap.Any("time", start.Format(time.RFC3339)))

		next.ServeHTTP(w, req)

		logger.Info("ending the request", zap.Any("time", start.Format(time.RFC3339)), zap.Any("elapsed time", time.Since(start).Milliseconds()))
	}
	return http.HandlerFunc(fn)
}
