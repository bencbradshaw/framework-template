package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// statusRecorder wraps http.ResponseWriter to capture the HTTP status code.
// This allows the logging middleware to log the actual response status.
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(statusCode int) {
	r.status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

// LoggingMiddleware logs HTTP requests with method, status code, duration, and path.
// Outputs both start and completion logs for request tracing.
// Measures and reports the total request processing time.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		fmt.Printf("START: [%s %s]\n", r.Method, r.URL.Path)
		recorder := &statusRecorder{
			ResponseWriter: w,
			status:         http.StatusOK,
		}
		next.ServeHTTP(recorder, r)
		finishTime := time.Now()
		totalTime := finishTime.Sub(startTime)
		fmt.Printf("[%s %d %v] %s\n", r.Method, recorder.status, totalTime, r.URL.Path)
	})
}
