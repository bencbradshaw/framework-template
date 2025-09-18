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

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserID(r)
	print("UserID:", userID, "OK:", ok, "\n")
	if !ok {
		writeErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// In a real application, you would fetch from database
	user := User{
		ID:        userID,
		Email:     "user@example.com",
		Name:      "Example User",
		CreatedAt: time.Now().AddDate(0, -1, 0), // 1 month ago
	}

	writeSuccessResponse(w, user, http.StatusOK)
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
