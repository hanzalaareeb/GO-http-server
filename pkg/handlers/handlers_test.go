// Description: This file contains tests for our API handlers.
// We use Go's standard `testing` package along with `net/http/httptest`
// to simulate HTTP requests and inspect the responses.

package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/hanzalaareeb/HTTPGolang/pkg/httpcontext"
	"github.com/hanzalaareeb/HTTPGolang/pkg/router"
)

// TestHealthCheckHandler tests the /health endpoint.
func TestHealthCheckHandler(t *testing.T) {
	// 1. Create a new request for the /health endpoint.
	// We don't need to provide a request body for a GET request.
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// 2. Create a ResponseRecorder.
	// This is a special type from httptest that acts as an http.ResponseWriter
	// but records the result (status code, headers, body) for inspection.
	rr := httptest.NewRecorder()

	// 3. Create a new router and register the handler.
	// We need a router to dispatch the request to the correct handler.
	r := router.New()
	r.GET("/health", HealthCheckHandler)

	// 4. Serve the HTTP request.
	// This will execute the handler associated with the request's path and method.
	r.ServeHTTP(rr, req)

	// 5. Check the status code.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// 6. Check the response body.
	// We expect a specific JSON response.
	expected := map[string]string{
		"status":  "ok",
		"service": "HTTPGolang_Server",
	}
	var actual map[string]string
	// Unmarshal the JSON from the response body into our `actual` map.
	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatalf("could not unmarshal response body: %v", err)
	}

	// Use reflect.DeepEqual for a robust comparison of maps/structs.
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}
}

// TestGetUsersHandler tests the /users endpoint for GET requests.
func TestGetUsersHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handlerCtx := &httpcontext.Context{Writer: rr, Request: req}

	// We can call the handler directly with a mocked context.
	GetUsersHandler(handlerCtx)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := []User{
		{ID: 1, Name: "Hanzala"},
		{ID: 2, Name: "Areeb"},
	}
	var actual []User
	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatalf("could not unmarshal response body: %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}
}

// TestCreateUserHandler tests the /users endpoint for POST requests.
func TestCreateUserHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/users", nil) // Body is nil for this simple case.
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	r := router.New()
	r.POST("/users", CreateUserHandler)

	r.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check response body
	expected := map[string]string{"status": "user created successfully"}
	var actual map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatalf("could not unmarshal response body: %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}
}
