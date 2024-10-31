package middleware

import (
	"net/http"
	"time"

	"github.com/yosa12978/dazed/internal/logging"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logger(logger logging.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			writer := &wrappedWriter{
				ResponseWriter: w,
				statusCode:     200,
			}
			snapshot := time.Now()
			next.ServeHTTP(writer, r)
			latency := time.Since(snapshot).Microseconds()
			logger.Info(
				"incoming request",
				"endpoint", r.URL.Path,
				"method", r.Method,
				"status_code", writer.statusCode,
				"latency_us", latency,
			)
		})
	}
}
