// Description: This package contains the application's request handlers.
// Each handler is responsible for the business logic of a specific API endpoint.

package handlers

import (
	"net/http" // Provides HTTP status constants like http.StatusOK.

	"github.com/hanzalaareeb/HTTPGolang/pkg/httpcontext"
	"github.com/hanzalaareeb/HTTPGolang/pkg/router"
)

// User represents a user in our system.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// RegisterRoutes is a function that registers all the application's routes
// with the provided router. This keeps the route setup organized and separate
// from the main application startup logic.
func RegisterRoutes(r *router.Router) {
	r.GET("/health", HealthCheckHandler)
	r.GET("/users", GetUsersHandler)
	r.POST("/users", CreateUserHandler)
}

// HealthCheckHandler handles the /health endpoint.
// It's a simple handler to check if the service is running.
func HealthCheckHandler(c *httpcontext.Context) {
	// Use the JSON helper from our custom context to send a response.
	// The `map[string]string` will be automatically encoded into a JSON object.
	c.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"service": "HTTPGolang_Server",
	})
}

// GetUsersHandler handles requests to retrieve a list of users.
func GetUsersHandler(c *httpcontext.Context) {
	// In a real application, you would fetch this data from a database.
	// Here, we're just using a static list for demonstration.
	users := []User{
		{ID: 1, Name: "Hanzala"},
		{ID: 2, Name: "Areeb"},
	}

	// Send the list of users as a JSON array.
	c.JSON(http.StatusOK, users)
}

// CreateUserHandler handles requests to create a new user.
func CreateUserHandler(c *httpcontext.Context) {
	// For a POST request, you would typically decode the request body.
	// For example:
	// var newUser User
	// if err := json.NewDecoder(c.Request.Body).Decode(&newUser); err != nil {
	//     c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	//     return
	// }
	//
	// log.Printf("Created new user: %v", newUser)

	// For this example, we'll just return a success message.
	c.JSON(http.StatusCreated, map[string]string{
		"status": "user created successfully",
	})
}
