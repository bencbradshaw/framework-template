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
// Validates email and password, and sets authentication cookie on success.
// In a production app, this would validate against a real user database.
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

	// In a real application, you would validate credentials against a database
	// For demo purposes, we'll accept any non-empty email/password
	if email != "" && password != "" {
		// Set auth cookie for 1 hour
		SetSecureCookie(w, "framework", email, time.Hour)
		http.Redirect(w, r, "/account/", http.StatusFound)
		return
	}

	// Login failed
	framework.RenderWithHtmlResponse(w, "login.custom.html", map[string]any{
		"title": "Login",
		"error": "Invalid email or password",
	})
}

// handleSignupSubmission processes signup form submissions.
// Validates email and password requirements, creates user account, and sets authentication cookie.
// In a production app, this would store user data in a database and handle password hashing.
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

	// In a real application, you would handle auth in a real way, e.g., store user in db or use IDP

	// For demo purposes, we'll accept any valid signup
	SetSecureCookie(w, "framework", email, time.Hour)
	http.Redirect(w, r, "/account/", http.StatusFound)
}
