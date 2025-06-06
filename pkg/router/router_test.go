// Description: This file contains tests for our custom router.
// It ensures that routes are registered correctly and that requests are
// dispatched to the appropriate handlers. It also tests the "not found" case.

package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hanzalaareeb/HTTPGolang/pkg/httpcontext"
)

// TestRouter_ServeHTTP_Found tests if a registered route is correctly handled.
func TestRouter_ServeHTTP_Found(t *testing.T) {
	// 1. Setup: Create a new router and register a test route.
	r := New()
	// This flag will be set to true inside our handler if it gets called.
	handlerCalled := false
	testHandler := func(c *httpcontext.Context) {
		handlerCalled = true // Mark that the handler was called.
		c.Status(http.StatusOK)
	}
	r.GET("/test", testHandler)

	// 2. Create a request and recorder.
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	rr := httptest.NewRecorder()

	// 3. Execute: Call the router's ServeHTTP method.
	r.ServeHTTP(rr, req)

	// 4. Assert: Check if the handler was called and if the status code is correct.
	if !handlerCalled {
		t.Error("expected the test handler to be called, but it was not")
	}
	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %d, but got %d", http.StatusOK, rr.Code)
	}
}

// TestRouter_ServeHTTP_NotFound tests if the router correctly returns a 404
// for an unregistered route.
func TestRouter_ServeHTTP_NotFound(t *testing.T) {
	// 1. Setup: Create a new router but don't register the route we're going to request.
	r := New()
	r.GET("/exists", func(c *httpcontext.Context) {})

	// 2. Create a request to a non-existent path.
	req, err := http.NewRequest("GET", "/does-not-exist", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	rr := httptest.NewRecorder()

	// 3. Execute: Call the router's ServeHTTP method.
	r.ServeHTTP(rr, req)

	// 4. Assert: Check if the status code is 404 Not Found.
	if rr.Code != http.StatusNotFound {
		t.Errorf("expected status code %d for an undefined route, but got %d", http.StatusNotFound, rr.Code)
	}
}

// TestRouter_ServeHTTP_MethodNotAllowed tests if the router correctly returns a 404
// when the path exists but the method does not. Our simple router returns 404,
// a more advanced one might return 405 Method Not Allowed.
func TestRouter_ServeHTTP_MethodNotAllowed(t *testing.T) {
	// 1. Setup: Register a GET route.
	r := New()
	r.GET("/test", func(c *httpcontext.Context) {})

	// 2. Create a POST request to the same path.
	req, err := http.NewRequest("POST", "/test", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	rr := httptest.NewRecorder()

	// 3. Execute: Call ServeHTTP.
	r.ServeHTTP(rr, req)

	// 4. Assert: Check for 404 Not Found.
	if rr.Code != http.StatusNotFound {
		t.Errorf("expected status code %d for wrong method, but got %d", http.StatusNotFound, rr.Code)
	}
}
