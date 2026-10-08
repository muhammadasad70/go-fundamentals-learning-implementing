package main

import (
	"encoding/json"
	"net/http"
	"net/mail"
	"strings"
)

func (s *UserStore) CreateUserhandler(w http.ResponseWriter, r *http.Request) {
	var input CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid Json Body", http.StatusBadRequest)
		return
	}
	input.Name = strings.TrimSpace(input.Name)
	input.Email = strings.TrimSpace(input.Email)

	if input.Name == "" || input.Email == "" {
		http.Error(w, "Name and Email can't be null", http.StatusBadRequest)
		return
	}
	address, err := mail.ParseAddress(input.Email)
	if err != nil || address.Address != input.Email {
		http.Error(w, "Invalid user email", http.StatusBadRequest)
		return

	}
	user := s.CreateUser(input.Name, input.Email)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}
