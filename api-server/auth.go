package main

import (
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Pasword  string `json:"password"`
}
type LoginResponse struct {
	Token string `json:"token"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not alowed", http.StatusMethodNotAllowed)
		return
	}
	var req LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return

	}
	// temporary hardcoded credential for learning purpose
	if req.Username != "asad" || req.Pasword != "1234" {
		http.Error(w, "invalid credential", http.StatusUnauthorized)
		return
	}

	token, err := generateToken(req.Username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}
	response := LoginResponse{
		Token: token,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
