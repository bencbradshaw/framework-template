package auth

import (
	"net/http"
	"os"
	"time"
)

// SetSecureCookie sets a cookie with secure defaults based on environment
func SetSecureCookie(w http.ResponseWriter, name, value string, maxAge time.Duration) {
	isProd := os.Getenv("APP_ENV") == "production"

	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		MaxAge:   int(maxAge.Seconds()),
		HttpOnly: true,
		Secure:   isProd,
		SameSite: http.SameSiteStrictMode,
	})
}

// ClearCookie clears a cookie by setting it to expire immediately
func ClearCookie(w http.ResponseWriter, name string) {
	isProd := os.Getenv("APP_ENV") == "production"

	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   isProd,
		SameSite: http.SameSiteStrictMode,
	})
}

// Define a custom type for context keys to avoid collisions (matching middleware)
type contextKey string

const userIDKey contextKey = "userID"

// GetUserID extracts user ID from request context
// This is a helper function for use in handlers after authentication middleware
func GetUserID(r *http.Request) (string, bool) {
	userID, ok := r.Context().Value(userIDKey).(string)
	return userID, ok && userID != ""
}
