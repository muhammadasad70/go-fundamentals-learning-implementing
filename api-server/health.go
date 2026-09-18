package main

import (
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
