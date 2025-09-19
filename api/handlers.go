package api

import (
	"encoding/json"
	"framework-template/auth"
	"net/http"
	"time"
)

// User represents a user account with basic profile information.
// Used for API responses and data transfer between frontend and backend.
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email,omitempty"`
	Name      string    `json:"name,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// APIResponse provides a consistent structure for all API responses.
// Includes success status, data payload, and error messages.
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// UserHandler demonstrates RESTful API patterns with proper error handling.
// Supports GET requests to retrieve user data and PUT requests to update user information.
// Requires authentication via the auth middleware.
func UserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			handleGetUser(w, r)
		case http.MethodPut:
			handleUpdateUser(w, r)
		default:
			w.Header().Set("Allow", "GET, PUT")
			writeErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// AllUsersHandler returns all users in the system (admin only).
// Only accessible by users with admin role for user management purposes.
func AllUsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodGet {
			w.Header().Set("Allow", "GET")
			writeErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Check if user is admin
		user, ok := auth.GetUser(r)
		if !ok {
			writeErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if user.Role != "admin" {
			writeErrorResponse(w, "Forbidden - Admin access required", http.StatusForbidden)
			return
		}

		// Get all users from sample data
		var allUsers []User
		for _, authUser := range auth.SampleUsers {
			apiUser := User{
				ID:        authUser.Email,
				Email:     authUser.Email,
				Name:      authUser.Email, // Using email as name for demo
				Role:      authUser.Role,
				CreatedAt: time.Now().AddDate(0, -1, 0), // Demo value - 1 month ago
			}
			allUsers = append(allUsers, apiUser)
		}

		writeSuccessResponse(w, allUsers, http.StatusOK)
	}
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	// Get full user information including role
	user, ok := auth.GetUser(r)
	if !ok {
		writeErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Convert auth.User to API User format
	apiUser := User{
		ID:        user.Email, // Using email as ID for this demo
		Email:     user.Email,
		Name:      user.Email, // Using email as name for demo
		Role:      user.Role,
		CreatedAt: time.Now(), // Demo value
	}

	writeSuccessResponse(w, apiUser, http.StatusOK)
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserID(r)
	if !ok {
		writeErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var updateData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		writeErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// In a real application, you would update in database
	// For demo purposes, just return success
	response := map[string]interface{}{
		"id":      userID,
		"updated": updateData,
	}

	writeSuccessResponse(w, response, http.StatusOK)
}

func writeSuccessResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Data:    data,
	})
}

func writeErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(APIResponse{
		Success: false,
		Error:   message,
	})
}
