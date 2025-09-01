package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// statusRecorder captures the status code for logging
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(statusCode int) {
	r.status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

// LoggingMiddleware logs HTTP requests with method, status, duration, and path
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
