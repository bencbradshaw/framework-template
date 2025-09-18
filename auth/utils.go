package auth

import (
	"net/http"
	"os"
	"time"
)

// SetSecureCookie sets an HTTP cookie with secure defaults based on the environment.
// In production (APP_ENV=production), cookies are marked as Secure.
// All cookies are HttpOnly and use SameSite=Strict for security.
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

// ClearCookie clears an HTTP cookie by setting it to expire immediately.
// Uses the same security settings as SetSecureCookie for consistency.
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

// GetUserID extracts the user identifier from the authentication cookie.
// Returns the user ID and a boolean indicating if authentication was successful.
// This is a helper function for use in handlers after authentication middleware.
func GetUserID(r *http.Request) (string, bool) {
	cookie, err := r.Cookie("framework")
	if err != nil {
		return "", false
	}
	return cookie.Value, cookie.Value != ""
}
