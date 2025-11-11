// Rate Limiting: Controlling request throughput
// Key concepts: Token bucket, leaky bucket, sliding window

package main

import (
	"fmt"
	"sync"
	"time"
)

// TokenBucket rate limiter
type TokenBucket struct {
	capacity     int
	tokens       int
	refillRate   int
	refillPeriod time.Duration
	lastRefill   time.Time
	mutex        sync.Mutex
}

func NewTokenBucket(capacity, refillRate int, refillPeriod time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:     capacity,
		tokens:       capacity,
		refillRate:   refillRate,
		refillPeriod: refillPeriod,
		lastRefill:   time.Now(),
	}
}

func (tb *TokenBucket) Allow() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	// Refill tokens
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)

	if elapsed >= tb.refillPeriod {
		tokensToAdd := int(elapsed/tb.refillPeriod) * tb.refillRate
		tb.tokens = min(tb.capacity, tb.tokens+tokensToAdd)
		tb.lastRefill = now
	}

	// Try to consume token
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}

	return false
}

func (tb *TokenBucket) Tokens() int {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	return tb.tokens
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Sliding window rate limiter
type SlidingWindow struct {
	limit    int
	window   time.Duration
	requests []time.Time
	mutex    sync.Mutex
}

func NewSlidingWindow(limit int, window time.Duration) *SlidingWindow {
	return &SlidingWindow{
		limit:  limit,
		window: window,
	}
}

func (sw *SlidingWindow) Allow() bool {
	sw.mutex.Lock()
	defer sw.mutex.Unlock()

	now := time.Now()
	cutoff := now.Add(-sw.window)

	// Remove old requests
	validRequests := sw.requests[:0]
	for _, reqTime := range sw.requests {
		if reqTime.After(cutoff) {
			validRequests = append(validRequests, reqTime)
		}
	}
	sw.requests = validRequests

	// Check if we can add new request
	if len(sw.requests) < sw.limit {
		sw.requests = append(sw.requests, now)
		return true
	}

	return false
}

func (sw *SlidingWindow) RequestCount() int {
	sw.mutex.Lock()
	defer sw.mutex.Unlock()
	return len(sw.requests)
}

// Channel-based rate limiter (simple)
type ChannelLimiter struct {
	ticker *time.Ticker
	tokens chan struct{}
}

func NewChannelLimiter(rps int, burstSize int) *ChannelLimiter {
	limiter := &ChannelLimiter{
		ticker: time.NewTicker(time.Second / time.Duration(rps)),
		tokens: make(chan struct{}, burstSize),
	}

	// Fill initial burst capacity
	for i := 0; i < burstSize; i++ {
		limiter.tokens <- struct{}{}
	}

	// Start refilling
	go limiter.refill()

	return limiter
}

func (cl *ChannelLimiter) refill() {
	for range cl.ticker.C {
		select {
		case cl.tokens <- struct{}{}:
		default:
			// Channel full, skip
		}
	}
}

func (cl *ChannelLimiter) Allow() bool {
	select {
	case <-cl.tokens:
		return true
	default:
		return false
	}
}

func (cl *ChannelLimiter) Wait() {
	<-cl.tokens
}

func (cl *ChannelLimiter) Close() {
	cl.ticker.Stop()
	close(cl.tokens)
}

// API server with rate limiting
func simulateAPIRequests(name string, limiter interface{}, requestCount int) {
	fmt.Printf("\n🌐 %s API Simulation (%d requests):\n", name, requestCount)

	allowed := 0
	denied := 0

	for i := 0; i < requestCount; i++ {
		var canProceed bool

		switch l := limiter.(type) {
		case *TokenBucket:
			canProceed = l.Allow()
		case *SlidingWindow:
			canProceed = l.Allow()
		case *ChannelLimiter:
			canProceed = l.Allow()
		}

		if canProceed {
			allowed++
			fmt.Printf("✅ Request %d: Allowed\n", i+1)
		} else {
			denied++
			fmt.Printf("❌ Request %d: Rate limited\n", i+1)
		}

		time.Sleep(50 * time.Millisecond) // Simulate request interval
	}

	fmt.Printf("📊 Summary - Allowed: %d, Denied: %d\n", allowed, denied)
}

// Distributed rate limiting simulation
func distributedRateLimit() {
	fmt.Println("\n=== Distributed Rate Limiting ===")

	// Shared rate limiter across multiple goroutines
	limiter := NewChannelLimiter(5, 3) // 5 RPS, burst of 3
	defer limiter.Close()

	var wg sync.WaitGroup

	// Simulate 3 concurrent clients
	for clientID := 0; clientID < 3; clientID++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for i := 0; i < 4; i++ {
				if limiter.Allow() {
					fmt.Printf("🟢 Client %d - Request %d: Allowed\n", id, i+1)
				} else {
					fmt.Printf("🔴 Client %d - Request %d: Rate limited\n", id, i+1)
				}
				time.Sleep(100 * time.Millisecond)
			}
		}(clientID)
	}

	wg.Wait()
}

func main() {
	fmt.Println("=== Rate Limiting Patterns ===")

	// Example 1: Token Bucket
	fmt.Println("\n1. Token Bucket Rate Limiter:")
	tokenBucket := NewTokenBucket(3, 1, time.Second) // 3 tokens, refill 1/second

	fmt.Printf("Initial tokens: %d\n", tokenBucket.Tokens())
	simulateAPIRequests("TokenBucket", tokenBucket, 8)

	// Wait and show refill
	time.Sleep(2 * time.Second)
	fmt.Printf("Tokens after 2s: %d\n", tokenBucket.Tokens())

	// Example 2: Sliding Window
	fmt.Println("\n2. Sliding Window Rate Limiter:")
	slidingWindow := NewSlidingWindow(4, 2*time.Second) // 4 requests per 2 seconds

	simulateAPIRequests("SlidingWindow", slidingWindow, 8)

	// Example 3: Channel-based
	fmt.Println("\n3. Channel-based Rate Limiter:")
	channelLimiter := NewChannelLimiter(3, 2) // 3 RPS, burst of 2
	defer channelLimiter.Close()

	simulateAPIRequests("ChannelLimiter", channelLimiter, 6)

	// Example 4: Distributed limiting
	distributedRateLimit()

	fmt.Println("\n✅ Rate limiting demo complete")
}
