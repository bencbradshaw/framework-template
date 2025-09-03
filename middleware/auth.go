package middleware

import (
	"context"
	"fmt"
	"net/http"
)

// Define a custom type for context keys to avoid collisions
type contextKey string

const userIDKey contextKey = "userID"

// AuthMiddleware checks for authentication cookie and adds user ID to context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("AUTH CHECK: [%s %s]\n", r.Method, r.URL.Path)
		cookie, err := r.Cookie("framework")
		if err != nil {
			// Redirect to login page instead of showing 401 error
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// Extract user ID from cookie value
		// In a real app, you'd decode/verify a JWT or look up a session
		// For now, assuming the cookie value IS the user ID
		userID := cookie.Value

		// Add user ID to request context
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
