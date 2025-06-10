// Description: This is the main entry point for our HTTP server application.
// It's responsible for setting up the router, registering our API endpoints (handlers),
// and starting the server.

package main

import (
	"log"
	"os"

	"github.com/hanzalaareeb/HTTPGolang/pkg/handlers"
	"github.com/hanzalaareeb/HTTPGolang/pkg/router"
	"github.com/hanzalaareeb/HTTPGolang/pkg/server"
)

// main is the function where the execution of the program begins.
func main() {
	// 1. Create a new instance of our custom router.
	// The router will be responsible for mapping incoming requests to the correct handler.
	log.Println("Initializing router...")
	r := router.New()

	// 2. Register application routes.
	// We delegate the registration of specific routes to the handlers package
	// to keep our main function clean and organized. This is a good practice
	// for modularity.
	log.Println("Registering application handlers...")
	handlers.RegisterRoutes(r)

	// 3. Create a new server instance.
	// We configure it to listen on port 8080 and use our custom router
	// to handle all incoming requests.
	// The server package abstracts away the details of the underlying http.Server.
	port := ":8080"
	s := server.New(port, r)

	// 4. Start the server.
	// We run this in a goroutine so it doesn't block the main thread.
	// This allows us to listen for shutdown signals gracefully.
	go func() {
		log.Printf("Server starting on port %s...", port)
		if err := s.Start(); err != nil {
			// If the server fails to start (e.g., port is already in use),
			// log the error and exit.
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// 5. Graceful Shutdown
	// The code below waits for a shutdown signal (like Ctrl+C).
	// Currently, our server doesn't have a Stop method, but this is where
	// you would call it. For now, we just block and wait.
	// In a real-world application, you would use a channel to listen for
	// signals like syscall.SIGINT and syscall.SIGTERM.
	log.Println("Application started. Press Ctrl+C to exit.")
	quit := make(chan os.Signal, 1)
	// In a complete implementation, you'd use signal.Notify(quit, os.Interrupt)
	// and then call a s.Stop() method when a signal is received.
	<-quit // Block until a signal is received.
	log.Println("Shutting down server...")

	// Here you would call s.Stop() to gracefully shut down the server.
	// more changes are required
}
