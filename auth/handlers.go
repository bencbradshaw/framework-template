package auth

import (
	"net/http"
	"time"

	"github.com/bencbradshaw/framework"
)

// LoginHandler handles both GET (show form) and POST (process login)
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

// SignupHandler handles both GET (show form) and POST (process signup)
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

// LogoutHandler clears the auth cookie and redirects to home
func LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ClearCookie(w, "framework")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

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
		SetSecureCookie(w, "framework", "demo-session", time.Hour)
		http.Redirect(w, r, "/account/", http.StatusFound)
		return
	}

	// Login failed
	framework.RenderWithHtmlResponse(w, "login.custom.html", map[string]any{
		"title": "Login",
		"error": "Invalid email or password",
	})
}

func handleSignupSubmission(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	// Basic validation
	if name == "" || email == "" || password == "" || confirmPassword == "" {
		framework.RenderWithHtmlResponse(w, "signup.custom.html", map[string]any{
			"title": "Sign Up",
			"error": "All fields are required",
		})
		return
	}

	if password != confirmPassword {
		framework.RenderWithHtmlResponse(w, "signup.custom.html", map[string]any{
			"title": "Sign Up",
			"error": "Passwords do not match",
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

	// In a real application, you would:
	// 1. Check if email already exists
	// 2. Hash the password
	// 3. Store user in database
	// 4. Send verification email

	// For demo purposes, we'll accept any valid signup
	SetSecureCookie(w, "framework", "demo-session", time.Hour)
	http.Redirect(w, r, "/account/", http.StatusFound)
}
