package main

// this is use to store the info client sent
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// this is for the server to send back the response

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
