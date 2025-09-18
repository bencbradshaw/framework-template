package middleware

import (
	"context"
	"fmt"
	"net/http"
)

// Define a custom type for context keys to avoid collisions
type contextKey string

const userIDKey contextKey = "userID"

// AuthMiddleware checks for authentication cookie and adds user ID to request context.
// If no valid authentication cookie is found, redirects to the login page.
// For authenticated requests, adds the user ID to the request context for use by handlers.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("AUTH CHECK: [%s %s]\n", r.Method, r.URL.Path)

		// Attempt to retrieve the authentication cookie
		cookie, err := r.Cookie("framework")
		if err != nil {
			// No valid authentication cookie found - redirect to login page
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// Extract user ID from cookie value
		// In a production app, you'd decode/verify a JWT or look up a session
		// For this demo, the cookie value IS the user ID (email)
		userID := cookie.Value

		// Add user ID to request context for use by downstream handlers
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		r = r.WithContext(ctx)

		// Continue to the protected handler
		next.ServeHTTP(w, r)
	})
}
