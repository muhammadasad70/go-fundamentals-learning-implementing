package main

import (
	"encoding/json"
	"net/http"
)

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
