package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// this is the DTO Concept

// we are using this struct because the incomming data will be in struct form mean our struct act as the schema or shape of the incoming request from the client
type HelloRequest struct {
	Name string `json:"name"`
}
type HelloResponse struct {
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// we are checking the method allowed to the user or not mean might be user can write /Use GET instead of POST so we are checking the allowed method
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// in this if client sent the request id then we fetch it using the r.Header because req id is the meta data of the request so it must be in header so req id created by client so we get it using this

	requestID := r.Header.Get("X-Request-ID")

	// but if client has not sent the request then our server apiGateway /proxy server or load balancer create it to trace the request

	if requestID == "" {
		requestID = fmt.Sprintf("%d", time.Now().UnixNano())
	}
	// As data from the user will be in json form so first we will decode the data into Go data likw we are using struct so it will fill our struct
	var req HelloRequest
	// we are also check the behaviour of the user mean is body invalid or empty mean user not semnd the name in the body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid body request", http.StatusBadRequest)
		return
	}
	// we also return the error if the name is empty from the client side mean request body has no name
	if req.Name == "" {
		http.Error(w, "Name must required", http.StatusBadRequest)
		return
	}
	// by using the Sprintf we are formating the respose and we are writing the response in the struct
	response := HelloResponse{
		Message:   fmt.Sprintf("Hello %s", req.Name),
		RequestID: requestID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Request-ID", requestID)
	w.WriteHeader(http.StatusOK)
	// in this we are encoding the go struct data into json
	json.NewEncoder(w).Encode(response)

}
