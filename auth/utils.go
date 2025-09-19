package auth

import (
	"net/http"
	"os"
	"time"
)

// User represents a user in the system with role-based access
type User struct {
	Email    string
	Password string
	Role     string // "admin" or "user"
}

// Sample hardcoded users for testing
// In production, this would come from a database
var SampleUsers = map[string]User{
	"admin@example.com": {
		Email:    "admin@example.com",
		Password: "admin123",
		Role:     "admin",
	},
	"user@example.com": {
		Email:    "user@example.com",
		Password: "user123",
		Role:     "user",
	},
	"john@example.com": {
		Email:    "john@example.com",
		Password: "password",
		Role:     "user",
	},
	"sarah@example.com": {
		Email:    "sarah@example.com",
		Password: "password",
		Role:     "admin",
	},
}

// ValidateCredentials checks if the provided email and password are valid
// Returns the user object and success status
func ValidateCredentials(email, password string) (User, bool) {
	user, exists := SampleUsers[email]
	if !exists {
		return User{}, false
	}

	if user.Password == password {
		return user, true
	}

	return User{}, false
}

// GetRedirectURL returns the appropriate redirect URL based on user role
func GetRedirectURL(role string) string {
	if role == "admin" {
		return "/admin/"
	}
	return "/app/"
}

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

//	GetUserID extracts the user identifier from the authentication cookie.
//
// Returns the user ID and a boolean indicating if authentication was successful.
// This is a helper function for use in handlers after authentication middleware.
func GetUserID(r *http.Request) (string, bool) {
	cookie, err := r.Cookie("framework")
	if err != nil {
		return "", false
	}
	return cookie.Value, cookie.Value != ""
}

// GetUser extracts the full user information from the authentication cookie.
// Returns the user object and a boolean indicating if authentication was successful.
func GetUser(r *http.Request) (User, bool) {
	email, ok := GetUserID(r)
	if !ok {
		return User{}, false
	}

	user, exists := SampleUsers[email]
	if !exists {
		return User{}, false
	}

	return user, true
}
