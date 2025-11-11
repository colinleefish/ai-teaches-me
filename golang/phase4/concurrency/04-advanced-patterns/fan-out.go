// Fan-Out Pattern: Distributing work to multiple workers
// Key concepts: Load distribution, result collection, worker coordination

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job represents work to be done
type Job struct {
	ID   int
	Data string
}

type Result struct {
	JobID  int
	Output string
	Worker int
}

// Fan-out work to multiple workers
func fanOut(jobs <-chan Job, numWorkers int) <-chan Result {
	results := make(chan Result)

	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Close results when all workers done
	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

func worker(workerID int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		// Simulate work
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

		result := Result{
			JobID:  job.ID,
			Output: fmt.Sprintf("Processed: %s", job.Data),
			Worker: workerID,
		}

		results <- result
	}
}

// Fan-out with result ordering (collect results in job order)
func fanOutOrdered(jobs []Job, numWorkers int) []Result {
	jobChan := make(chan Job, len(jobs))
	resultChan := make(chan Result)

	// Send jobs
	for _, job := range jobs {
		jobChan <- job
	}
	close(jobChan)

	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobChan, resultChan, &wg)
	}

	// Collect results
	results := make([]Result, len(jobs))
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect and order results
	for result := range resultChan {
		results[result.JobID] = result
	}

	return results
}

func main() {
	fmt.Println("=== Fan-Out Pattern ===")

	// Create jobs
	jobs := make(chan Job, 10)
	for i := 0; i < 10; i++ {
		jobs <- Job{
			ID:   i,
			Data: fmt.Sprintf("task-%d", i),
		}
	}
	close(jobs)

	// Example 1: Basic fan-out
	fmt.Println("\n1. Basic Fan-Out (3 workers):")
	results := fanOut(jobs, 3)

	for result := range results {
		fmt.Printf("Job %d completed by worker %d: %s\n",
			result.JobID, result.Worker, result.Output)
	}

	// Example 2: Ordered results
	fmt.Println("\n2. Fan-Out with Ordered Results:")
	jobList := make([]Job, 8)
	for i := 0; i < 8; i++ {
		jobList[i] = Job{
			ID:   i,
			Data: fmt.Sprintf("ordered-task-%d", i),
		}
	}

	orderedResults := fanOutOrdered(jobList, 4)

	for i, result := range orderedResults {
		fmt.Printf("Position %d: Job %d by worker %d: %s\n",
			i, result.JobID, result.Worker, result.Output)
	}
}
