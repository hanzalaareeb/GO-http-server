// Description: This package provides a wrapper around the standard Go http.Server,
// making it easier to configure and manage.

package server

import (
	"context"
	"net/http" // The core Go package for HTTP servers and clients.
	"time"
)

// Server holds the details for our HTTP server.
type Server struct {
	httpServer *http.Server
}

// New creates and configures a new Server instance.
// It takes a listening address (e.g., ":8080") and an http.Handler (our router) as arguments.
// An http.Handler is an interface that responds to an HTTP request. Our router will implement this.
func New(addr string, handler http.Handler) *Server {
	// We create an instance of the standard http.Server.
	// It's good practice to configure timeouts to prevent resource exhaustion
	// from slow or malicious clients.
	srv := &http.Server{
		Addr:         addr,              // The address to listen on.
		Handler:      handler,           // The handler to delegate requests to (our router).
		ReadTimeout:  5 * time.Second,   // Max time to read the entire request.
		WriteTimeout: 10 * time.Second,  // Max time to write the response.
		IdleTimeout:  120 * time.Second, // Max time for a connection to be idle.
	}

	return &Server{
		httpServer: srv,
	}
}

// Start makes the server begin listening for and serving HTTP requests.
// It's a blocking call.
func (s *Server) Start() error {
	// ListenAndServe starts the server and blocks until the server is shut down
	// or an error occurs. The error is returned, except for http.ErrServerClosed,
	// which indicates a graceful shutdown.
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Stop provides a way to gracefully shut down the server.
// It allows active connections to finish before closing.
func (s *Server) Stop(ctx context.Context) error {
	// Shutdown gracefully shuts down the server without interrupting any
	// active connections. It waits for them to finish up to the context deadline.
	return s.httpServer.Shutdown(ctx)
}
