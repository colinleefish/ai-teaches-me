// Circuit Breaker: Fail-fast for unhealthy services
// Key concepts: State management, failure detection, recovery

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type State int

const (
	StateClosed State = iota
	StateOpen
	StateHalfOpen
)

func (s State) String() string {
	switch s {
	case StateClosed:
		return "CLOSED"
	case StateOpen:
		return "OPEN"
	case StateHalfOpen:
		return "HALF-OPEN"
	default:
		return "UNKNOWN"
	}
}

// CircuitBreaker implementation
type CircuitBreaker struct {
	state            State
	failureCount     int
	successCount     int
	failureThreshold int
	successThreshold int
	timeout          time.Duration
	lastFailureTime  time.Time
	mutex            sync.RWMutex
}

func NewCircuitBreaker(failureThreshold, successThreshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:            StateClosed,
		failureThreshold: failureThreshold,
		successThreshold: successThreshold,
		timeout:          timeout,
	}
}

// Call executes a function with circuit breaker protection
func (cb *CircuitBreaker) Call(fn func() error) error {
	if !cb.allowRequest() {
		return errors.New("circuit breaker is OPEN")
	}

	err := fn()
	cb.recordResult(err)
	return err
}

func (cb *CircuitBreaker) allowRequest() bool {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	switch cb.state {
	case StateClosed:
		return true
	case StateOpen:
		// Check if timeout has passed
		if time.Since(cb.lastFailureTime) > cb.timeout {
			cb.state = StateHalfOpen
			cb.successCount = 0
			fmt.Println("🟡 Circuit breaker: OPEN → HALF-OPEN")
			return true
		}
		return false
	case StateHalfOpen:
		return true
	default:
		return false
	}
}

func (cb *CircuitBreaker) recordResult(err error) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if err != nil {
		cb.failureCount++
		cb.lastFailureTime = time.Now()

		if cb.state == StateHalfOpen {
			cb.state = StateOpen
			fmt.Println("🔴 Circuit breaker: HALF-OPEN → OPEN (failure)")
		} else if cb.state == StateClosed && cb.failureCount >= cb.failureThreshold {
			cb.state = StateOpen
			fmt.Printf("🔴 Circuit breaker: CLOSED → OPEN (failures: %d)\n", cb.failureCount)
		}
	} else {
		cb.successCount++

		if cb.state == StateHalfOpen && cb.successCount >= cb.successThreshold {
			cb.state = StateClosed
			cb.failureCount = 0
			fmt.Printf("🟢 Circuit breaker: HALF-OPEN → CLOSED (successes: %d)\n", cb.successCount)
		}
	}
}

func (cb *CircuitBreaker) GetState() State {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	return cb.state
}

func (cb *CircuitBreaker) GetStats() (State, int, int) {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	return cb.state, cb.failureCount, cb.successCount
}

// Simulated external service
type ExternalService struct {
	failureRate float64
	latency     time.Duration
}

func NewExternalService(failureRate float64, latency time.Duration) *ExternalService {
	return &ExternalService{
		failureRate: failureRate,
		latency:     latency,
	}
}

func (es *ExternalService) Call() error {
	// Simulate latency
	time.Sleep(es.latency)

	// Simulate failures
	if rand.Float64() < es.failureRate {
		return errors.New("external service failure")
	}

	return nil
}

// Service with circuit breaker protection
type ProtectedService struct {
	circuitBreaker *CircuitBreaker
	externalSvc    *ExternalService
	callCount      int
}

func NewProtectedService(cb *CircuitBreaker, externalSvc *ExternalService) *ProtectedService {
	return &ProtectedService{
		circuitBreaker: cb,
		externalSvc:    externalSvc,
	}
}

func (ps *ProtectedService) ProcessRequest() error {
	ps.callCount++

	err := ps.circuitBreaker.Call(func() error {
		return ps.externalSvc.Call()
	})

	state, failures, successes := ps.circuitBreaker.GetStats()

	if err != nil {
		fmt.Printf("❌ Request %d: FAILED - %v [%s F:%d S:%d]\n",
			ps.callCount, err, state, failures, successes)
	} else {
		fmt.Printf("✅ Request %d: SUCCESS [%s F:%d S:%d]\n",
			ps.callCount, state, failures, successes)
	}

	return err
}

// Advanced circuit breaker with metrics
type MetricsCircuitBreaker struct {
	*CircuitBreaker
	totalRequests   int
	totalFailures   int
	totalSuccesses  int
	blockedRequests int
}

func NewMetricsCircuitBreaker(failureThreshold, successThreshold int, timeout time.Duration) *MetricsCircuitBreaker {
	return &MetricsCircuitBreaker{
		CircuitBreaker: NewCircuitBreaker(failureThreshold, successThreshold, timeout),
	}
}

func (mcb *MetricsCircuitBreaker) Call(fn func() error) error {
	mcb.totalRequests++

	if !mcb.allowRequest() {
		mcb.blockedRequests++
		return errors.New("circuit breaker is OPEN")
	}

	err := fn()
	mcb.recordResult(err)

	if err != nil {
		mcb.totalFailures++
	} else {
		mcb.totalSuccesses++
	}

	return err
}

func (mcb *MetricsCircuitBreaker) GetMetrics() (int, int, int, int) {
	mcb.mutex.RLock()
	defer mcb.mutex.RUnlock()
	return mcb.totalRequests, mcb.totalSuccesses, mcb.totalFailures, mcb.blockedRequests
}

func simulateServiceCalls(service *ProtectedService, count int) {
	fmt.Printf("\n🚀 Simulating %d service calls:\n", count)

	for i := 0; i < count; i++ {
		service.ProcessRequest()
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("=== Circuit Breaker Pattern ===")

	// Example 1: Basic circuit breaker
	fmt.Println("\n1. Basic Circuit Breaker:")

	// Create unstable service (70% failure rate)
	unstableService := NewExternalService(0.7, 50*time.Millisecond)

	// Circuit breaker: 3 failures to open, 2 successes to close, 2s timeout
	cb := NewCircuitBreaker(3, 2, 2*time.Second)

	protectedService := NewProtectedService(cb, unstableService)

	// This should trigger the circuit breaker to open
	simulateServiceCalls(protectedService, 10)

	fmt.Println("\n⏱️  Waiting for circuit breaker timeout...")
	time.Sleep(2500 * time.Millisecond)

	// Create stable service for recovery testing
	stableService := NewExternalService(0.1, 50*time.Millisecond) // 10% failure rate
	protectedService.externalSvc = stableService

	fmt.Println("\n🔄 Testing recovery with stable service:")
	simulateServiceCalls(protectedService, 5)

	// Example 2: Metrics circuit breaker
	fmt.Println("\n2. Circuit Breaker with Metrics:")

	metricsService := NewExternalService(0.4, 30*time.Millisecond)
	metricsCB := NewMetricsCircuitBreaker(2, 3, 1*time.Second)

	// Make calls and track metrics
	for i := 0; i < 15; i++ {
		err := metricsCB.Call(func() error {
			return metricsService.Call()
		})

		total, success, failures, blocked := metricsCB.GetMetrics()
		state, _, _ := metricsCB.GetStats()

		status := "✅"
		if err != nil {
			status = "❌"
		}

		fmt.Printf("%s Call %d [%s] - Total:%d Success:%d Failed:%d Blocked:%d\n",
			status, i+1, state, total, success, failures, blocked)

		time.Sleep(150 * time.Millisecond)
	}

	// Final metrics
	total, success, failures, blocked := metricsCB.GetMetrics()
	fmt.Printf("\n📊 Final Metrics:\n")
	fmt.Printf("   Total Requests: %d\n", total)
	fmt.Printf("   Successful: %d (%.1f%%)\n", success, float64(success)/float64(total)*100)
	fmt.Printf("   Failed: %d (%.1f%%)\n", failures, float64(failures)/float64(total)*100)
	fmt.Printf("   Blocked: %d (%.1f%%)\n", blocked, float64(blocked)/float64(total)*100)

	fmt.Println("\n✅ Circuit breaker demo complete")
}
