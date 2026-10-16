package main

import (
	"encoding/json"
	"net/http"
	"net/mail"
	"strconv"
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
func (s *UserStore) ListUserhandler(w http.ResponseWriter, r *http.Request) {
	users := s.List()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
func (s *UserStore) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// we are reading the value from the url path just like we read the r.Body
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}
	user, ok := s.GetById(id)

	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
func (s *UserStore) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input CreateUserRequest

	idstring := r.PathValue("id")
	id, err := strconv.Atoi(idstring)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid json body", http.StatusBadRequest)
		return
	}
	input.Name = strings.TrimSpace(input.Name)
	input.Email = strings.TrimSpace(input.Email)

	if input.Name == "" || input.Email == "" {
		http.Error(w, "Name and Email are required", http.StatusBadRequest)
		return
	}
	address, err := mail.ParseAddress(input.Email)
	if err != nil || address.Address != input.Email {
		http.Error(w, "User email is invalid", http.StatusBadRequest)
		return
	}
	user, ok := s.UpdateUser(id, input.Name, input.Email)
	if !ok {
		http.Error(w, "Unable to update the user because user not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}
