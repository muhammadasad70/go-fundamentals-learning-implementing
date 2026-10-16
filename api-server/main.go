package main

import (
	"fmt"
	"net/http"
)

func main() {
	store := NewUserStore()
	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", store.CreateUserhandler)
	mux.HandleFunc("GET /users", store.ListUserhandler)
	mux.HandleFunc("GET /users/{id}", store.GetUserHandler)
	mux.HandleFunc("PUT /users/{id}", store.UpdateUserHandler)
	mux.HandleFunc("DELETE /users/{id}", store.DeleteUserHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/hello", helloHandler)
	fmt.Println("server running on port :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error from server", err)
	}
}
