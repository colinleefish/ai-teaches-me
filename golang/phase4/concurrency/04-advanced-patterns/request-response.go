// Request-Response Pattern: Bidirectional communication
// Key concepts: Reply channels, correlation, async responses

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Request with reply channel
type Request struct {
	ID      int
	Data    string
	ReplyTo chan<- Response
}

type Response struct {
	RequestID int
	Result    string
	Error     error
}

// Simple request-response server
func startServer(requests <-chan Request) {
	go func() {
		for req := range requests {
			// Simulate processing
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

			response := Response{
				RequestID: req.ID,
				Result:    fmt.Sprintf("Processed: %s", req.Data),
			}

			// Send response back
			req.ReplyTo <- response
		}
	}()
}

// Async client that sends requests and collects responses
func asyncClient(requests chan<- Request, numRequests int) <-chan Response {
	responses := make(chan Response, numRequests)

	go func() {
		defer close(responses)
		var wg sync.WaitGroup

		for i := 0; i < numRequests; i++ {
			wg.Add(1)

			// Create response channel for this request
			replyChan := make(chan Response, 1)

			// Send request
			request := Request{
				ID:      i,
				Data:    fmt.Sprintf("request-%d", i),
				ReplyTo: replyChan,
			}

			requests <- request

			// Handle response asynchronously
			go func() {
				defer wg.Done()
				response := <-replyChan
				responses <- response
			}()
		}

		wg.Wait()
	}()

	return responses
}

// Request-response with correlation map (for multiplexed connections)
type CorrelatedServer struct {
	requests  chan Request
	responses map[int]chan<- Response
	mutex     sync.RWMutex
}

func NewCorrelatedServer() *CorrelatedServer {
	server := &CorrelatedServer{
		requests:  make(chan Request),
		responses: make(map[int]chan<- Response),
	}

	// Start processing
	go server.processRequests()

	return server
}

func (s *CorrelatedServer) processRequests() {
	for req := range s.requests {
		go func(r Request) {
			// Process request
			time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)

			// Send response using correlation
			s.mutex.RLock()
			if replyChan, exists := s.responses[r.ID]; exists {
				replyChan <- Response{
					RequestID: r.ID,
					Result:    fmt.Sprintf("Correlated result for: %s", r.Data),
				}
			}
			s.mutex.RUnlock()
		}(req)
	}
}

func (s *CorrelatedServer) SendRequest(id int, data string) <-chan Response {
	replyChan := make(chan Response, 1)

	// Register response channel
	s.mutex.Lock()
	s.responses[id] = replyChan
	s.mutex.Unlock()

	// Send request
	s.requests <- Request{
		ID:   id,
		Data: data,
	}

	return replyChan
}

// Pipeline with request-response stages
func pipelineWithRequests() {
	fmt.Println("\nPipeline with Request-Response:")

	// Stage 1: Input processing
	stage1Requests := make(chan Request)
	startServer(stage1Requests)

	// Stage 2: Further processing
	stage2Requests := make(chan Request)
	startServer(stage2Requests)

	// Coordinator that chains the pipeline
	finalResults := make(chan string)

	go func() {
		defer close(finalResults)

		for i := 0; i < 5; i++ {
			// Stage 1
			stage1Reply := make(chan Response, 1)
			stage1Requests <- Request{
				ID:      i,
				Data:    fmt.Sprintf("input-%d", i),
				ReplyTo: stage1Reply,
			}

			stage1Response := <-stage1Reply

			// Stage 2
			stage2Reply := make(chan Response, 1)
			stage2Requests <- Request{
				ID:      i,
				Data:    stage1Response.Result,
				ReplyTo: stage2Reply,
			}

			stage2Response := <-stage2Reply
			finalResults <- fmt.Sprintf("Final: %s", stage2Response.Result)
		}
	}()

	// Collect final results
	for result := range finalResults {
		fmt.Println(result)
	}
}

func main() {
	fmt.Println("=== Request-Response Pattern ===")

	// Example 1: Basic request-response
	fmt.Println("\n1. Basic Request-Response:")
	requests := make(chan Request)
	startServer(requests)

	responses := asyncClient(requests, 5)

	for response := range responses {
		fmt.Printf("Response %d: %s\n", response.RequestID, response.Result)
	}

	// Example 2: Correlated server
	fmt.Println("\n2. Correlated Request-Response:")
	server := NewCorrelatedServer()

	// Send multiple requests
	var responseChannels []<-chan Response
	for i := 0; i < 4; i++ {
		respChan := server.SendRequest(i, fmt.Sprintf("correlated-data-%d", i))
		responseChannels = append(responseChannels, respChan)
	}

	// Collect responses (may arrive out of order)
	for i, respChan := range responseChannels {
		response := <-respChan
		fmt.Printf("Correlated response %d: %s\n", i, response.Result)
	}

	// Example 3: Pipeline with request-response
	pipelineWithRequests()
}
