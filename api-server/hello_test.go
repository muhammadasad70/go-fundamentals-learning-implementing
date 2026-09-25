package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandlerSuccess(t *testing.T) {
	// creating fake body for fake request
	body := `{"name":"Asad"}`
	// by using NewRequest we create the fake request like method handler and the body and using the NewReader we are reading the fake body because when client send the request so we read the stream so by using the
	// Reade we are reading fake body
	req := httptest.NewRequest(
		http.MethodPost, "/hello", strings.NewReader(body),
	)
	// setting the fake content type and the fake request-id
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Request-ID", "req-123")
	// NewRecorder by using this we are creating the fake server which behave like the real server
	recorder := httptest.NewRecorder()
	// we are calling the handler directly and passing the recorder which is for fake response and for the request which we pass in the real server
	helloHandler(recorder, req)
	// now we are testing all thr cases 1 by 1
	// test case for the success
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status %d and got %d", http.StatusOK, recorder.Code)
	}
	// test case for the body
	expectedBody := `{"message":"Hello Asad","request_id":"req-123"}`

	actualBody := strings.TrimSpace(recorder.Body.String())
	if actualBody != expectedBody {
		t.Errorf("expected body %s and got %s", expectedBody, actualBody)
	}

	// test case for the request id

	if recorder.Header().Get("X-Request-ID") != "req-123" {
		t.Errorf("expected request_id req-123,got %s", recorder.Header().Get("X-Request-ID"))
	}

}
