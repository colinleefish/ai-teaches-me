// Week 9: HTTP Client Operations
// This file demonstrates net/http package for making HTTP requests

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Sample data structures for HTTP examples
type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// TODO: Demonstrate basic HTTP GET requests
func demonstrateBasicGET() {
	fmt.Println("=== Basic HTTP GET Requests ===")
	
	// TODO: Simple GET request with http.Get
	url := "https://jsonplaceholder.typicode.com/users/1"
	
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return
	}
	defer resp.Body.Close() // Always close the response body
	
	// TODO: Check response status
	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Status: %s\n", resp.Status)
	
	// TODO: Read response headers
	fmt.Println("Response Headers:")
	for key, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", key, value)
		}
	}
	
	// TODO: Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}
	
	fmt.Printf("Response Body: %s\n", string(body))
	
	// TODO: Parse JSON response
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}
	
	fmt.Printf("Parsed User: %+v\n", user)
}

// TODO: Demonstrate custom HTTP requests
func demonstrateCustomRequests() {
	fmt.Println("\n=== Custom HTTP Requests ===")
	
	// TODO: Create custom request with http.NewRequest
	url := "https://jsonplaceholder.typicode.com/posts"
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}
	
	// TODO: Add custom headers
	req.Header.Set("User-Agent", "Go-HTTP-Client/1.0")
	req.Header.Set("Accept", "application/json")
	req.Header.Add("X-Custom-Header", "custom-value")
	
	// TODO: Add query parameters
	q := req.URL.Query()
	q.Add("userId", "1")
	q.Add("_limit", "3")
	req.URL.RawQuery = q.Encode()
	
	fmt.Printf("Request URL: %s\n", req.URL.String())
	
	// TODO: Execute request with custom client
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error executing request: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	// TODO: Handle response
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
		return
	}
	
	var posts []Post
	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}
	
	fmt.Printf("Retrieved %d posts\n", len(posts))
	for _, post := range posts {
		fmt.Printf("Post %d: %s\n", post.ID, post.Title)
	}
}

// TODO: Demonstrate POST requests with JSON data
func demonstratePOSTRequests() {
	fmt.Println("\n=== POST Requests with JSON ===")
	
	// TODO: Prepare JSON data
	newPost := Post{
		UserID: 1,
		Title:  "My New Post",
		Body:   "This is the content of my new post.",
	}
	
	jsonData, err := json.Marshal(newPost)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}
	
	// TODO: Create POST request
	url := "https://jsonplaceholder.typicode.com/posts"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error creating POST request: %v\n", err)
		return
	}
	
	// TODO: Set appropriate headers for JSON
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	
	// TODO: Execute request
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error executing POST request: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	fmt.Printf("POST Response Status: %d\n", resp.StatusCode)
	
	// TODO: Handle response
	var createdPost Post
	err = json.NewDecoder(resp.Body).Decode(&createdPost)
	if err != nil {
		fmt.Printf("Error decoding response: %v\n", err)
		return
	}
	
	fmt.Printf("Created Post: %+v\n", createdPost)
}

// TODO: Demonstrate form data requests
func demonstrateFormRequests() {
	fmt.Println("\n=== Form Data Requests ===")
	
	// TODO: URL-encoded form data
	formData := url.Values{}
	formData.Set("name", "John Doe")
	formData.Set("email", "john@example.com")
	formData.Add("interests", "programming")
	formData.Add("interests", "music")
	
	// TODO: Create POST request with form data
	req, err := http.NewRequest("POST", "https://httpbin.org/post", 
		strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Printf("Error creating form request: %v\n", err)
		return
	}
	
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	// TODO: Execute request
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error executing form request: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	// TODO: Read and display response
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Form Response (first 200 chars): %.200s...\n", string(body))
}

// TODO: Demonstrate HTTP client configuration
func demonstrateClientConfiguration() {
	fmt.Println("\n=== HTTP Client Configuration ===")
	
	// TODO: Custom client with timeouts
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			// TODO: Connection pooling settings
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 2,
			IdleConnTimeout:     30 * time.Second,
			// TODO: TLS settings
			// TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	
	// TODO: Test with configured client
	testURL := "https://jsonplaceholder.typicode.com/users"
	
	start := time.Now()
	resp, err := client.Get(testURL)
	elapsed := time.Since(start)
	
	if err != nil {
		fmt.Printf("Error with configured client: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	fmt.Printf("Request completed in: %v\n", elapsed)
	fmt.Printf("Response status: %s\n", resp.Status)
	
	// TODO: Connection reuse demonstration
	fmt.Println("Making second request to test connection reuse...")
	start = time.Now()
	resp2, err := client.Get(testURL)
	elapsed2 := time.Since(start)
	
	if err != nil {
		fmt.Printf("Error with second request: %v\n", err)
		return
	}
	defer resp2.Body.Close()
	
	fmt.Printf("Second request completed in: %v\n", elapsed2)
	fmt.Printf("Speed improvement: %v\n", elapsed-elapsed2)
}

// TODO: Demonstrate context with HTTP requests
func demonstrateHTTPWithContext() {
	fmt.Println("\n=== HTTP Requests with Context ===")
	
	// TODO: Request with timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	url := "https://jsonplaceholder.typicode.com/posts"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request with context: %v\n", err)
		return
	}
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// TODO: Check for context timeout
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Request timed out")
		} else {
			fmt.Printf("Request error: %v\n", err)
		}
		return
	}
	defer resp.Body.Close()
	
	fmt.Printf("Request completed successfully: %s\n", resp.Status)
	
	// TODO: Cancellable request
	fmt.Println("Testing request cancellation...")
	
	cancelCtx, cancelFunc := context.WithCancel(context.Background())
	
	// Cancel after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		cancelFunc()
	}()
	
	longReq, _ := http.NewRequestWithContext(cancelCtx, "GET", 
		"https://httpbin.org/delay/5", nil)
	
	_, err = client.Do(longReq)
	if err != nil {
		if cancelCtx.Err() == context.Canceled {
			fmt.Println("Request was cancelled")
		} else {
			fmt.Printf("Request error: %v\n", err)
		}
	}
}

// TODO: Demonstrate error handling and retries
func demonstrateErrorHandling() {
	fmt.Println("\n=== Error Handling and Retries ===")
	
	// TODO: Retry logic with exponential backoff
	retryRequest := func(url string, maxRetries int) (*http.Response, error) {
		client := &http.Client{Timeout: 5 * time.Second}
		
		for attempt := 0; attempt < maxRetries; attempt++ {
			resp, err := client.Get(url)
			
			if err == nil && resp.StatusCode == http.StatusOK {
				return resp, nil
			}
			
			if resp != nil {
				resp.Body.Close()
			}
			
			// Calculate backoff delay
			delay := time.Duration(attempt+1) * time.Second
			fmt.Printf("Attempt %d failed, retrying in %v...\n", attempt+1, delay)
			time.Sleep(delay)
		}
		
		return nil, fmt.Errorf("failed after %d attempts", maxRetries)
	}
	
	// TODO: Test with unreliable endpoint
	unreliableURL := "https://httpbin.org/status/500"
	
	resp, err := retryRequest(unreliableURL, 3)
	if err != nil {
		fmt.Printf("All retry attempts failed: %v\n", err)
	} else {
		defer resp.Body.Close()
		fmt.Printf("Request succeeded: %s\n", resp.Status)
	}
	
	// TODO: Different error types
	demonstrateErrorTypes := func() {
		// Network error
		_, err := http.Get("http://invalid-domain-12345.com")
		if err != nil {
			fmt.Printf("Network error: %v\n", err)
		}
		
		// Timeout error
		client := &http.Client{Timeout: 1 * time.Millisecond}
		_, err = client.Get("https://httpbin.org/delay/1")
		if err != nil {
			fmt.Printf("Timeout error: %v\n", err)
		}
	}
	
	demonstrateErrorTypes()
}

// TODO: Demonstrate authentication
func demonstrateAuthentication() {
	fmt.Println("\n=== HTTP Authentication ===")
	
	// TODO: Basic Authentication
	basicAuthExample := func() {
		req, _ := http.NewRequest("GET", "https://httpbin.org/basic-auth/user/pass", nil)
		req.SetBasicAuth("user", "pass")
		
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Basic auth error: %v\n", err)
			return
		}
		defer resp.Body.Close()
		
		fmt.Printf("Basic auth response: %s\n", resp.Status)
	}
	
	// TODO: Bearer Token Authentication
	bearerTokenExample := func() {
		req, _ := http.NewRequest("GET", "https://httpbin.org/bearer", nil)
		req.Header.Set("Authorization", "Bearer your-token-here")
		
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Bearer token error: %v\n", err)
			return
		}
		defer resp.Body.Close()
		
		fmt.Printf("Bearer token response: %s\n", resp.Status)
	}
	
	// TODO: API Key Authentication
	apiKeyExample := func() {
		req, _ := http.NewRequest("GET", "https://httpbin.org/get", nil)
		req.Header.Set("X-API-Key", "your-api-key")
		
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("API key error: %v\n", err)
			return
		}
		defer resp.Body.Close()
		
		fmt.Printf("API key response: %s\n", resp.Status)
	}
	
	basicAuthExample()
	bearerTokenExample()
	apiKeyExample()
}

// TODO: Demonstrate cookie handling
func demonstrateCookieHandling() {
	fmt.Println("\n=== Cookie Handling ===")
	
	// TODO: Client with cookie jar
	jar := &http.CookieJar{}
	client := &http.Client{
		Jar: jar,
	}
	
	// TODO: Set cookies
	setCookieURL := "https://httpbin.org/cookies/set/session/abc123"
	resp, err := client.Get(setCookieURL)
	if err != nil {
		fmt.Printf("Error setting cookie: %v\n", err)
		return
	}
	resp.Body.Close()
	
	fmt.Println("Cookie set successfully")
	
	// TODO: Get cookies
	getCookieURL := "https://httpbin.org/cookies"
	resp, err = client.Get(getCookieURL)
	if err != nil {
		fmt.Printf("Error getting cookies: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Cookies response: %s\n", string(body))
	
	// TODO: Manual cookie handling
	manualCookieExample := func() {
		req, _ := http.NewRequest("GET", "https://httpbin.org/cookies", nil)
		
		// Add cookies manually
		cookie := &http.Cookie{
			Name:  "manual-cookie",
			Value: "manual-value",
		}
		req.AddCookie(cookie)
		
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Manual cookie error: %v\n", err)
			return
		}
		defer resp.Body.Close()
		
		fmt.Printf("Manual cookie response: %s\n", resp.Status)
	}
	
	manualCookieExample()
}

// TODO: Demonstrate HTTP client patterns
func demonstrateHTTPPatterns() {
	fmt.Println("\n=== HTTP Client Patterns ===")
	
	// TODO: API Client wrapper
	type APIClient struct {
		BaseURL    string
		HTTPClient *http.Client
		APIKey     string
	}
	
	func NewAPIClient(baseURL, apiKey string) *APIClient {
		return &APIClient{
			BaseURL: baseURL,
			APIKey:  apiKey,
			HTTPClient: &http.Client{
				Timeout: 10 * time.Second,
			},
		}
	}
	
	func (c *APIClient) makeRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
		url := c.BaseURL + endpoint
		req, err := http.NewRequest(method, url, body)
		if err != nil {
			return nil, err
		}
		
		// Add common headers
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "Go-API-Client/1.0")
		
		return c.HTTPClient.Do(req)
	}
	
	func (c *APIClient) GetUsers() ([]User, error) {
		resp, err := c.makeRequest("GET", "/users", nil)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		
		var users []User
		err = json.NewDecoder(resp.Body).Decode(&users)
		return users, err
	}
	
	// TODO: Test API client
	client := NewAPIClient("https://jsonplaceholder.typicode.com", "test-key")
	users, err := client.GetUsers()
	if err != nil {
		fmt.Printf("API client error: %v\n", err)
	} else {
		fmt.Printf("Retrieved %d users via API client\n", len(users))
	}
}

// Helper function for timing HTTP operations
func timeHTTPOperation(name string, fn func()) {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	fmt.Printf("%s took: %v\n", name, elapsed)
}

func main() {
	fmt.Println("🌐 Welcome to HTTP Client Operations! 🌐")
	fmt.Println("This file teaches you how to make HTTP requests in Go")
	
	// TODO: Implement each demonstration function
	// Start with basic GET requests and progress to advanced patterns
	
	demonstrateBasicGET()
	// demonstrateCustomRequests()
	// demonstratePOSTRequests()
	// demonstrateFormRequests()
	// demonstrateClientConfiguration()
	// demonstrateHTTPWithContext()
	// demonstrateErrorHandling()
	// demonstrateAuthentication()
	// demonstrateCookieHandling()
	// demonstrateHTTPPatterns()
	
	fmt.Println("\n🎉 Congratulations! You've mastered HTTP client operations!")
	fmt.Println("Next: Learn context usage in context_usage.go")
}

/*
🔍 Key Concepts to Remember:

1. **Basic Requests**: http.Get(), http.Post(), http.NewRequest()
2. **Response Handling**: Always close response body, check status codes
3. **Client Configuration**: Timeouts, connection pooling, custom transport
4. **Context**: Use context for cancellation and timeouts
5. **Authentication**: Basic auth, bearer tokens, API keys
6. **Error Handling**: Network errors, timeouts, retries
7. **Performance**: Connection reuse, proper client configuration

📋 Essential HTTP Operations:
```go
// Basic GET
resp, err := http.Get(url)
defer resp.Body.Close()

// Custom request
req, _ := http.NewRequest("POST", url, body)
req.Header.Set("Content-Type", "application/json")
client := &http.Client{Timeout: 10 * time.Second}
resp, err := client.Do(req)

// With context
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
```

🚨 Common Mistakes:
- Forgetting to close response body (resource leak)
- Not checking HTTP status codes
- Using default client without timeouts
- Not handling context cancellation
- Ignoring connection reuse opportunities
- Poor error handling and retry logic

🎯 Next Steps:
- Learn context package for cancellation patterns
- Practice with real APIs and authentication
- Build robust HTTP client libraries
- Master timeout and retry strategies
*/
