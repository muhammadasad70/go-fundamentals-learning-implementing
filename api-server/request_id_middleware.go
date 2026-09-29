package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// we are using context to store the request id

// we are creating our oown type based on string
type contextKey string

// we are creating and using our own type to avoid the collision like we can use the string type but if the two package use the string type for contextKey then conflict can be occure

const requestIDKey contextKey = "requestID"

// we have created the const to store and reterive the request id from context
// we are creating an middleware which accept and handler wwhich is next http.Handler and we attach our logic with it and it return an handler
// in this next http.Handler mean which handler will run when middleware will finish like in this case hello handler
// returning handler mean our middleware itself become something who can handle http
func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			// by using the sprintf we are converting the generated id by time .now().UnixNano() into string
			requestID = fmt.Sprintf("%d", time.Now().UnixNano())
		}
		w.Header().Set("X-Request-ID", requestID)
		// r.Context() mean this is an existing context and in go each http request already has a context
		// we are just adding the request id to the context
		// mean we are replacing the existing context with new key value pair like requestIDKey and requestID
		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		// context.WithValue will return new context
		// we are adding the request in the context so with this requestid will move freely in the entire application and can b use any wear like for logs benchmarking and others factor
		r = r.WithContext(ctx)
		// with r.WithContext() we are attaching request with context
		next.ServeHTTP(w, r)
		// we are passing control to the next handler like when middleware done request will not stope it goes to the other handlers
	})

}
