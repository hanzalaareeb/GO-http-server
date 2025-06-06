// Description: This package defines a custom Context struct that wraps the standard
// http.ResponseWriter and *http.Request. It allows us to add helper methods
// for common tasks, like sending JSON responses.

package httpcontext

import (
	"encoding/json" // For encoding data into JSON format.
	"fmt"
	"net/http"
)

// Context wraps the standard http.ResponseWriter and *http.Request.
type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

// JSON is a helper method to send a JSON response.
// It takes a status code and a data payload (which can be any Go struct or map).
func (c *Context) JSON(statusCode int, data interface{}) {
	// Set the Content-Type header to indicate that the response body is JSON.
	c.Writer.Header().Set("Content-Type", "application/json")

	// Write the HTTP status code to the response header.
	// This must be done before writing the body.
	c.Writer.WriteHeader(statusCode)

	// Encode the data payload into JSON and write it to the response body.
	// json.NewEncoder is efficient as it writes directly to the writer's output stream.
	if err := json.NewEncoder(c.Writer).Encode(data); err != nil {
		// If an error occurs during JSON encoding, log it.
		// In a real app, you might have more robust error handling here.
		http.Error(c.Writer, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

// String is a helper method to send a plain text response.
func (c *Context) String(statusCode int, format string, values ...interface{}) {
	// Set the Content-Type header to plain text.
	c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// Write the HTTP status code.
	c.Writer.WriteHeader(statusCode)
	// Write the formatted string to the response body.
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// Status is a helper method to send a response with only a status code and no body.
func (c *Context) Status(statusCode int) {
	c.Writer.WriteHeader(statusCode)
}
