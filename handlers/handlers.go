package handlers

import (
	"encoding/json"
	"goLangExpress/auth"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func jsonError(w http.ResponseWriter, errMsg string, status int) {
	jsonResponse(w, map[string]string{"error": errMsg}, status)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		jsonError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = auth.CreateUser(user.Username, user.Password)
	if err != nil {
		jsonError(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, map[string]string{"message": "User created successfully"}, http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		jsonError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if auth.AuthenticateUser(user.Username, user.Password) {
		token, err := auth.GenerateToken(user.Username)
		if err != nil {
			jsonError(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		jsonResponse(w, map[string]string{"token": token}, http.StatusOK)
	} else {
		jsonError(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// In a stateless JWT implementation, logout is typically handled client-side
	// by removing the token. Here we'll just return a success message.
	jsonResponse(w, map[string]string{"message": "Logged out successfully"}, http.StatusOK)
}

func Me(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("Username")
	jsonResponse(w, map[string]string{"username": username}, http.StatusOK)
}

func Protected1(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, map[string]string{"message": "This is protected route 1"}, http.StatusOK)
}

func Protected2(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, map[string]string{"message": "This is protected route 2"}, http.StatusOK)
}

func Protected3(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, map[string]string{"message": "This is protected route 3"}, http.StatusOK)
}

func Protected4(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, map[string]string{"message": "This is protected route 4"}, http.StatusOK)
}

func Protected5(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, map[string]string{"message": "This is protected route 5"}, http.StatusOK)
}