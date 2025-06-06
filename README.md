# Go HTTP Server From Scratch

This project is an educational example of how to build a modular, production-ready HTTP server in Go using only the standard library. It demonstrates fundamental concepts like routing, request handling, and structuring a web service application without relying on third-party frameworks like Gin or Echo.

The goal is to provide a clear and well-commented codebase that explains the "why" behind the structure, making it a great starting point for anyone looking to understand the mechanics of Go's `net/http` package.

## Project Structure

The project is organized into several packages to promote modularity and separation of concerns.

``` bash
HTTPGolang/
├── cmd/
│   └── server/
│       └── main.go         # Main application entry point. Initializes and starts the server.
├── pkg/
│   ├── server/
│   │   └── server.go       # A wrapper around the standard http.Server for easy configuration.
│   ├── router/
│   │   ├── router.go       # A custom HTTP router to map requests to handlers.
│   │   └── router_test.go  # Tests for the router.
│   ├── httpcontext/
│   │   └── context.go      # A custom context with helper functions for handlers (e.g., sending JSON).
│   └── handlers/
│       ├── handlers.go     # Application-specific business logic (API handlers).
│       └── handlers_test.go# Tests for the API handlers.
├── go.mod                  # Defines the Go module and its properties.
├── .gitignore              # Specifies files for Git to ignore.
└── README.md               # This file.
```

## Getting Started

### Prerequisites

- Go version 1.22 or higher.

### Installation

1. Clone the repository:

    ```bash
        git clone https://github.com/hanzalaareeb/GO-http-server.git
        cd HTTPGolang
    ```

2. Tidy dependencies:
    This will ensure your go.mod file is in sync. Since there are no external dependencies, it will just verify the setup.

    ```bash
        go mod tidy
    ```

## Usage

### Running the Server

To start the HTTP server, run the `main.go` file from the root of the project:

```bash
    go run cmd/server/main.go
```

The server will start and listen on port 8080.

```bash
    2024/06/07 12:00:00 Initializing router...
    2024/06/07 12:00:00 Registered route: GET /health
    2024/06/07 12:00:00 Registered route: GET /users
    2024/06/07 12:00:00 Registered route: POST /users
    2024/06/07 12:00:00 Registering application handlers...
    2024/06/07 12:00:00 Server starting on port :8080...
    2024/06/07 12:00:00 Application started. Press Ctrl+C to exit.
```

### Running the Tests

To run the unit tests for all packages, execute the following command from the root of the project:

```bash
    go test ./... -v
```

The `-v` flag enables verbose output, showing the status of each test.

## API Endpoints

The server exposes the following endpoints:

| Method | Path | Description | Example curl Command |
|----------|----------|----------|----------|
| GET | /health | Checks the health of the service. | curl <http://localhost:8080/health> |
| GET | /users | Retrieves a static list of users. | curl <http://localhost:8080/users> |
| POST | /users | Simulates the creation of a new user. | curl -X POST <http://localhost:8080/users> |
| ANY | /anything | (Non-existent) Returns a 404 Not Found response. | curl <http://localhost:8080/non-existent-route> |
