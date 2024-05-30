package api

import (
	"log/slog"
	"net/http"
	"time"
)

// Extend the ResponseWriter interface to capture the status code for logging.
type wrappedResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *wrappedResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

// LogMiddleware logs incoming requests.
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		writer := &wrappedResponseWriter{ResponseWriter: w}
		next.ServeHTTP(writer, r)

		slog.Info("request", "method", r.Method, "path", r.URL.Path, "duration", time.Since(start))
	})
}
