package auth

import (
	"net/http"
	"time"

	"github.com/bencbradshaw/framework"
)

// LoginHandler handles both GET (show form) and POST (process login).
// GET requests render the login form template.
// POST requests validate credentials and set authentication cookies.
func LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			framework.RenderWithHtmlResponse(w, "login.custom.html", map[string]any{
				"title": "Login",
			})
		case http.MethodPost:
			handleLoginSubmission(w, r)
		default:
			w.Header().Set("Allow", "GET, POST")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// SignupHandler handles both GET (show form) and POST (process signup).
// GET requests render the signup form template.
// POST requests validate input, create user account, and set authentication cookies.
func SignupHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			framework.RenderWithHtmlResponse(w, "signup.custom.html", map[string]any{
				"title": "Sign Up",
			})
		case http.MethodPost:
			handleSignupSubmission(w, r)
		default:
			w.Header().Set("Allow", "GET, POST")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// LogoutHandler clears the authentication cookie and redirects to home page.
// This handler only accepts GET requests and immediately logs out the user.
func LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ClearCookie(w, "framework")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// handleLoginSubmission processes login form submissions.
// Validates email and password against hardcoded users, and sets authentication cookie on success.
// Redirects to appropriate dashboard based on user role (admin -> /admin/, user -> /app/).
func handleLoginSubmission(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Basic validation
	if email == "" || password == "" {
		framework.RenderWithHtmlResponse(w, "login.custom.html", map[string]any{
			"title": "Login",
			"error": "Email and password are required",
		})
		return
	}

	// Validate credentials against sample users
	user, valid := ValidateCredentials(email, password)
	if !valid {
		framework.RenderWithHtmlResponse(w, "login.custom.html", map[string]any{
			"title": "Login",
			"error": "Invalid email or password",
		})
		return
	}

	// Set auth cookie for 1 hour
	SetSecureCookie(w, "framework", email, time.Hour)

	// Redirect based on user role
	redirectURL := GetRedirectURL(user.Role)
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

// handleSignupSubmission processes signup form submissions.
// Validates email and password requirements, creates user account, and sets authentication cookie.
// For demo purposes, new users are assigned "user" role and redirected to /app/.
// In production, this would store user data in a database with proper password hashing.
func handleSignupSubmission(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Basic validation
	if email == "" || password == "" {
		framework.RenderWithHtmlResponse(w, "signup.custom.html", map[string]any{
			"title": "Sign Up",
			"error": "All fields are required",
		})
		return
	}

	if len(password) < 6 {
		framework.RenderWithHtmlResponse(w, "signup.custom.html", map[string]any{
			"title": "Sign Up",
			"error": "Password must be at least 6 characters long",
		})
		return
	}

	// Check if user already exists
	if _, exists := SampleUsers[email]; exists {
		framework.RenderWithHtmlResponse(w, "signup.custom.html", map[string]any{
			"title": "Sign Up",
			"error": "User with this email already exists",
		})
		return
	}

	// For demo purposes, add new user to sample users with "user" role
	// In production, this would be stored in a database
	SampleUsers[email] = User{
		Email:    email,
		Password: password,
		Role:     "user",
	}

	// Set auth cookie for 1 hour
	SetSecureCookie(w, "framework", email, time.Hour)

	// New users get "user" role, redirect to /app/
	http.Redirect(w, r, "/app/", http.StatusFound)
}
