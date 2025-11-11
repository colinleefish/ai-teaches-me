// Week 11: Web Server Practice
// Build production-ready web applications and APIs

package main

import (
	"fmt"
)

// TODO: Exercise 1 - RESTful API Server
func exercise1_RESTfulAPI() {
	fmt.Println("=== Exercise 1: RESTful API Server ===")

	// TODO: Build a complete REST API for a blog system:
	// 1. CRUD operations for posts (GET, POST, PUT, DELETE)
	// 2. JSON request/response handling
	// 3. URL routing with path parameters
	// 4. Input validation and error handling
	// 5. Middleware for logging and CORS

	// Example endpoints:
	// GET    /api/posts       - List all posts
	// GET    /api/posts/:id   - Get specific post
	// POST   /api/posts       - Create new post
	// PUT    /api/posts/:id   - Update post
	// DELETE /api/posts/:id   - Delete post

	fmt.Println("Exercise 1 completed!")
}

// TODO: Exercise 2 - Authentication Middleware
func exercise2_AuthenticationMiddleware() {
	fmt.Println("\n=== Exercise 2: Authentication Middleware ===")

	// TODO: Implement JWT-based authentication:
	// 1. User registration and login endpoints
	// 2. JWT token generation and validation
	// 3. Authentication middleware
	// 4. Protected routes
	// 5. Role-based access control

	fmt.Println("Exercise 2 completed!")
}

// TODO: Exercise 3 - File Upload Service
func exercise3_FileUploadService() {
	fmt.Println("\n=== Exercise 3: File Upload Service ===")

	// TODO: Build a file upload and management service:
	// 1. File upload with size limits
	// 2. File type validation
	// 3. Storage organization
	// 4. File serving with proper headers
	// 5. Thumbnail generation for images

	fmt.Println("Exercise 3 completed!")
}

func main() {
	fmt.Println("🌐 Welcome to Web Server Practice! 🌐")

	exercise1_RESTfulAPI()
	// exercise2_AuthenticationMiddleware()
	// exercise3_FileUploadService()

	fmt.Println("\n🎉 Ready for database integration!")
}
