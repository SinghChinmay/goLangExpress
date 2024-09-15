package middleware

import (
	"encoding/json"
	"net/http"
	"strings"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Message string `json:"error"`
}

// Implement the error interface for ErrorResponse
func (e *ErrorResponse) Error() string {
	return e.Message
}

// Helper function for JSON error responses
func jsonError(w http.ResponseWriter, errMsg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Message: errMsg})
}

func ValidateUserInput(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			jsonError(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := validateUser(user); err != nil {
			jsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Create a new request with the validated user data
		newReq := r.Clone(r.Context())
		newReq.Body = &dummyReadCloser{strings.NewReader(encodeUser(user))}
		
		next.ServeHTTP(w, newReq)
	}
}

func validateUser(user User) error {
	if strings.TrimSpace(user.Username) == "" {
		return &ErrorResponse{Message: "Username cannot be empty"}
	}
	if len(user.Username) < 3 || len(user.Username) > 50 {
		return &ErrorResponse{Message: "Username must be between 3 and 50 characters"}
	}
	if strings.TrimSpace(user.Password) == "" {
		return &ErrorResponse{Message: "Password cannot be empty"}
	}
	if len(user.Password) < 6 {
		return &ErrorResponse{Message: "Password must be at least 6 characters long"}
	}
	return nil
}

func encodeUser(user User) string {
	b, _ := json.Marshal(user)
	return string(b)
}

type dummyReadCloser struct {
	*strings.Reader
}

func (dummyReadCloser) Close() error {
	return nil
}