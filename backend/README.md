## Backend Directory Structure

### Overview
The [`backend`] directory contains all the server-side code for the Real-Time Forum Project. It is organized into several subdirectories, each responsible for different aspects of the backend functionality.

### Directory Structure
```
backend/
	.env
	cmd/
		app/
			...
		bootservices/
	go.mod
	go.sum
	orm/
		delete.go
		insert.go
		models.go
		orm.go
		scan.go
		sql.go
		tags.go
		types.go
		update.go
	server/
		gateway/
		microservices/
		middleware/
		router/
		server.go
	services/
		auth/
		chat/
		notification/
		posts/
	user.json
	utils/
		jwt/
		key/
		utils.go
		Validation/
```

### [`.env`]
Contains environment variables for configuring the backend services. Example:
```
AUTH_SERVICE="http://localhost:8080"
CHAT_SERVICE="http://localhost:9090"
POST&COMMENT_SERVICE="http://localhost:8181"
NOTIFICATION_SERVICE="http://localhost:9191"
```

### [`cmd/`]
Contains the entry points for the application.

#### [`cmd/app/`]
- **Purpose**: This directory contains the main application entry point.
- **Files**: 
  - [`main.go`]: Initializes and starts the server.

#### [`cmd/bootservices/`]
- **Purpose**: Contains bootstrapping code for initializing services.
- **Files**: 
  - [`main.go`]: Bootstraps and initializes all services.

### [`go.mod`] and [`go.sum`]
- **Purpose**: These files manage the dependencies for the Go modules used in the project.

### [`orm/`]
Contains the Object-Relational Mapping (ORM) layer for database interactions.

- **Files**:
  - `delete.go`: Handles deletion of records.
  - `insert.go`: Handles insertion of records.
  - `models.go`: Defines the data models.
  - `orm.go`: Core ORM functionalities.
  - `scan.go`: Scans database rows into Go structs.
  - `sql.go`: SQL query building utilities.
  - `tags.go`: Handles struct tags for ORM.
  - `types.go`: Defines custom types for ORM.
  - `update.go`: Handles updating of records.

### [`server/`]
Contains server configuration and routing.

#### [`server/gateway/`]
- **Purpose**: Manages the API Gateway, which routes requests to appropriate microservices.

#### [`server/microservices/`]
- **Purpose**: Contains implementations of various microservices.

#### [`server/middleware/`]
- **Purpose**: Contains middleware components for request processing.

#### [`server/router/`]
- **Purpose**: Defines the routes for the application.

#### [`server.go`]
- **Purpose**: Main server configuration and initialization.

### [`services/`]
Contains the business logic and microservices.

#### [`services/auth/`]
- **Purpose**: Manages user authentication.
- **Subdirectories**:
  - [`controllers/`]: Request handlers for authentication endpoints.
  - [`database/`]: Database migrations and models related to authentication.
  - [`models/`]: Data models for authentication.

#### [`services/chat/`]
- **Purpose**: Manages private messaging between users.
- **Subdirectories**:
  - [`controllers/`]: Request handlers for chat endpoints.
  - [`database/`]: Database migrations and models related to chat.
  - [`models/`]: Data models for chat.

#### [`services/notification/`]
- **Purpose**: Manages real-time notifications.
- **Subdirectories**:
  - [`controllers/`]: Request handlers for notification endpoints.
  - [`database/`]: Database migrations and models related to notifications.
  - [`models/`]: Data models for notifications.

#### [`services/posts/`]
- **Purpose**: Manages post creation and comments.
- **Subdirectories**:
  - [`controllers/`]: Request handlers for post endpoints.
  - [`database/`]: Database migrations and models related to posts.
  - [`models/`]: Data models for posts.

### [`user.json`]
- **Purpose**: Contains user-related data in JSON format.

### [`utils/`]
Contains utility functions and helpers.

#### [`utils/jwt/`]
- **Purpose**: Utilities for handling JSON Web Tokens (JWT).

#### [`utils/key/`]
- **Purpose**: Key management utilities.

#### [`utils/utils.go`]
- **Purpose**: General utility functions.

#### [`utils/Validation/`]
- **Purpose**: Input validation utilities.

## Detailed Documentation for Key Files

### [`cmd/app/main.go`]
```go
package main

import (
	"log"
	"net/http"
	"backend/server"
)

func main() {
	// Initialize the server
	srv := server.NewServer()

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", srv.Router))
}
```
- **Purpose**: Entry point for the application. Initializes and starts the server.
- **Key Functions**:
  - [`main()`]: Initializes the server and starts listening on port 8080.

### [`orm/models.go`]
```go
package orm

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
```
- **Purpose**: Defines the data models for the application.
- **Key Models**:
  - [`User`]: Represents a user in the system.
  - [`Post`]: Represents a post created by a user.

### [`server/server.go`]
```go
package server

import (
	"net/http"
	"backend/server/router"
)

type Server struct {
	Router *http.ServeMux
}

func NewServer() *Server {
	r := router.NewRouter()
	return &Server{Router: r}
}
```
- **Purpose**: Configures and initializes the server.
- **Key Functions**:
  - `NewServer()`: Creates a new server instance with the configured router.

### [`services/auth/controllers/login.go`]
```go
package controllers

import (
	"net/http"
	"encoding/json"
	"backend/services/auth/models"
	"backend/utils/jwt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Authenticate user
	token, err := jwt.GenerateToken(creds.Username)
	if err != nil {
		http.Error(w, "Authentication failed", http.StatusUnauthorized)
		return
	}

	// Return JWT token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
```
- **Purpose**: Handles user login requests.
- **Key Functions**:
  - `LoginHandler()`: Authenticates the user and returns a JWT token.

### [`utils/jwt/jwt.go`]
```go
package jwt

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   username,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
```
- **Purpose**: Provides utilities for generating and validating JWT tokens.
- **Key Functions**:
  - `GenerateToken()`: Generates a JWT token for a given username.

