package router

// Description: This package implements a simple HTTP router. It maps HTTP methods
// and URL paths to specific handler functions.

import (
	"log"
	"net/http"
	"sync"

	// We import our custom context package. The router's job is to create
	// this context for each request and pass it to the handler.
	"github.com/hanzalaareeb/HTTPGolang/pkg/httpcontext"
)

// HandlerFunc defines the type for our custom handler functions.
// Instead of the standard `func(http.ResponseWriter, *http.Request)`,
// our handlers will accept a `*httpcontext.Context`, which provides useful helpers.
type HandlerFunc func(*httpcontext.Context)

// Router is our main router struct. It holds the routing rules.
type Router struct {
	// We use a sync.RWMutex to protect the routes map from concurrent access.
	// This is important because routes might be read (during request handling)
	// and written (during setup) at the same time in more complex scenarios.
	mu sync.RWMutex

	// routes is a map that stores the handlers. The structure is:
	// map[HTTP_METHOD]map[URL_PATH]HandlerFunc
	// For example: routes["GET"]["/users"] = GetUsersHandler
	routes map[string]map[string]HandlerFunc
}

// New creates and returns a new Router instance.
func New() *Router {
	return &Router{
		// Initialize the routes map. It's crucial to initialize nested maps as well.
		routes: make(map[string]map[string]HandlerFunc),
	}
}

// addRoute is an internal helper to add a new route to the map.
func (r *Router) addRoute(method, path string, handler HandlerFunc) {
	// Lock the mutex for writing to ensure thread safety.
	r.mu.Lock()
	defer r.mu.Unlock() // Ensure the mutex is unlocked when the function exits.

	// Check if the map for the given HTTP method exists.
	if r.routes[method] == nil {
		// If not, create it.
		r.routes[method] = make(map[string]HandlerFunc)
	}
	r.routes[method][path] = handler
	log.Printf("Registered route: %s %s", method, path)
}

// GET is a convenience method for registering a handler for the GET HTTP method.
func (r *Router) GET(path string, handler HandlerFunc) {
	r.addRoute("GET", path, handler)
}

// POST is a convenience method for registering a handler for the POST HTTP method.
func (r *Router) POST(path string, handler HandlerFunc) {
	r.addRoute("POST", path, handler)
}

// ServeHTTP makes our Router implement the `http.Handler` interface.
// This method is called for every incoming HTTP request.
// It's the heart of the router.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Lock the mutex for reading. A read lock allows multiple readers at the same time.
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Find the handlers for the request's method.
	pathHandlers, ok := r.routes[req.Method]
	if !ok {
		// If no handlers are registered for this HTTP method, send a 404 Not Found.
		http.NotFound(w, req)
		return
	}

	// Find the specific handler for the request's URL path.
	handler, ok := pathHandlers[req.URL.Path]
	if !ok {
		// If no handler is registered for this specific path, send a 404 Not Found.
		http.NotFound(w, req)
		return
	}

	// Create a new instance of our custom context for this request.
	// This context wraps the original ResponseWriter and Request.
	ctx := &httpcontext.Context{
		Writer:  w,
		Request: req,
	}

	// Call the matched handler function with the newly created context.
	handler(ctx)
}
