// Graceful Shutdown: Clean service termination
// Key concepts: Signal handling, resource cleanup, coordinated shutdown

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Service represents a long-running service
type Service struct {
	name     string
	shutdown chan struct{}
	done     chan struct{}
}

func NewService(name string) *Service {
	return &Service{
		name:     name,
		shutdown: make(chan struct{}),
		done:     make(chan struct{}),
	}
}

func (s *Service) Start() {
	go func() {
		defer close(s.done)
		fmt.Printf("🚀 %s started\n", s.name)

		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Printf("📊 %s working...\n", s.name)
			case <-s.shutdown:
				fmt.Printf("🛑 %s shutting down...\n", s.name)
				time.Sleep(200 * time.Millisecond) // Simulate cleanup
				fmt.Printf("✅ %s stopped\n", s.name)
				return
			}
		}
	}()
}

func (s *Service) Stop() {
	close(s.shutdown)
	<-s.done
}

// Worker pool with graceful shutdown
type WorkerPool struct {
	workers  int
	jobs     chan int
	shutdown chan struct{}
	wg       sync.WaitGroup
}

func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		workers:  workers,
		jobs:     make(chan int, 10),
		shutdown: make(chan struct{}),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()

	for {
		select {
		case job, ok := <-wp.jobs:
			if !ok {
				fmt.Printf("💼 Worker %d: job channel closed\n", id)
				return
			}

			// Process job
			fmt.Printf("💼 Worker %d processing job %d\n", id, job)
			time.Sleep(100 * time.Millisecond)

		case <-wp.shutdown:
			fmt.Printf("💼 Worker %d shutting down\n", id)
			return
		}
	}
}

func (wp *WorkerPool) AddJob(job int) {
	select {
	case wp.jobs <- job:
	case <-wp.shutdown:
		fmt.Printf("❌ Cannot add job %d - shutting down\n", job)
	}
}

func (wp *WorkerPool) Shutdown(timeout time.Duration) {
	fmt.Println("🛑 Worker pool shutdown initiated...")

	// Stop accepting new jobs
	close(wp.shutdown)

	// Close job channel after brief pause
	go func() {
		time.Sleep(50 * time.Millisecond)
		close(wp.jobs)
	}()

	// Wait for workers with timeout
	done := make(chan struct{})
	go func() {
		wp.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("✅ All workers stopped gracefully")
	case <-time.After(timeout):
		fmt.Println("⏰ Worker shutdown timeout - forcing stop")
	}
}

// Context-based coordinated shutdown
func coordinatedShutdown() {
	fmt.Println("\n=== Coordinated Shutdown ===")

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// Start multiple services
	services := []string{"Database", "API Server", "Message Queue"}

	for _, serviceName := range services {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()

			ticker := time.NewTicker(300 * time.Millisecond)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					fmt.Printf("🔄 %s running\n", name)
				case <-ctx.Done():
					fmt.Printf("🛑 %s received shutdown signal\n", name)
					time.Sleep(100 * time.Millisecond) // Cleanup
					fmt.Printf("✅ %s shutdown complete\n", name)
					return
				}
			}
		}(serviceName)
	}

	// Simulate running for a bit
	time.Sleep(1 * time.Second)

	// Coordinate shutdown
	fmt.Println("📢 Initiating coordinated shutdown...")
	cancel()

	// Wait for all services
	wg.Wait()
	fmt.Println("🎯 Coordinated shutdown complete")
}

func main() {
	fmt.Println("=== Graceful Shutdown Patterns ===")

	// Example 1: Signal handling
	fmt.Println("\n1. Signal-based Shutdown (Press Ctrl+C to test):")

	// Create services
	webServer := NewService("WebServer")
	database := NewService("Database")

	// Start services
	webServer.Start()
	database.Start()

	// Setup signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Wait for signal or timeout
	select {
	case sig := <-sigChan:
		fmt.Printf("\n📨 Received signal: %v\n", sig)
		fmt.Println("🛑 Initiating graceful shutdown...")

		// Shutdown services in reverse order
		database.Stop()
		webServer.Stop()

		fmt.Println("✅ Graceful shutdown complete")

	case <-time.After(3 * time.Second):
		fmt.Println("\n⏰ Demo timeout - shutting down anyway")
		database.Stop()
		webServer.Stop()
	}

	// Example 2: Worker pool shutdown
	fmt.Println("\n2. Worker Pool Graceful Shutdown:")

	pool := NewWorkerPool(3)
	pool.Start()

	// Add some jobs
	for i := 0; i < 8; i++ {
		pool.AddJob(i)
		time.Sleep(50 * time.Millisecond)
	}

	// Graceful shutdown with timeout
	pool.Shutdown(2 * time.Second)

	// Example 3: Coordinated shutdown
	coordinatedShutdown()
}
