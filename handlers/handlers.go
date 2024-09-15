package handlers

import (
	"encoding/json"
	"goLangExpress/auth"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// In a real application, validate the user credentials here
	if user.Username == "testuser" && user.Password == "password" {
		token, err := auth.GenerateToken(user.Username)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// In a stateless JWT implementation, logout is typically handled client-side
	// by removing the token. Here we'll just return a success message.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})
}

func Me(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("Username")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"username": username})
}

func Protected1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "This is protected route 1"})
}

func Protected2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "This is protected route 2"})
}

func Protected3(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "This is protected route 3"})
}

func Protected4(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "This is protected route 4"})
}

func Protected5(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "This is protected route 5"})
}