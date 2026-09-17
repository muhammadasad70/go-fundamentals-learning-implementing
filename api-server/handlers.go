package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// in this we are comapring method allowed to the user mean this handler is only for get purpose so user can not try the post method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	// fprintln() mean write at destination like we use the fmt.println() to show the result on the terminal so F write to the destination which is w mean in response body
	fmt.Fprintln(w, "OK")
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// as in the http response header we also mention the Content-Type like json html image or any form we have to mention the Content_type
	w.Header().Set("Content_Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// we can use this
	// fmt.Fprintln(w, `message:Hello`)
	// But for proper go code we should not build the json string manually instead we encode the go data into json using the encode json package
	response := map[string]string{
		"message": "Hello",
	}
	// above we have created the go data mean by using the maps (key : value )
	json.NewEncoder(w).Encode(response)
	// above we conver the go data into json and writes it to the http response
}
