package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.Handle("/protected", authMiddleware(http.HandlerFunc(protectedHandler)))
	mux.HandleFunc("/login", loginHandler)
	fmt.Println("server running on port :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error from server", err)
	}
}
