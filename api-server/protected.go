package main

import (
	"encoding/json"
	"net/http"
)

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value(usernameKey).(string)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}
	respone := map[string]string{
		"message":  "Protected routes accessed successfully",
		"username": username,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respone)
}
