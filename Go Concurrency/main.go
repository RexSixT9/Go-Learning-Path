package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	URL string
	Err error
}

func worker(jobsChan chan string, wg *sync.WaitGroup, resultChan chan Result) {
	defer wg.Done()
	for job := range jobsChan {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("Worker processing image:", job)
		resultChan <- Result{URL: job, Err: nil}
	}
	fmt.Println("Worker finished processing jobs")
}

func main() {
	jobs := []string{
		"https://example.com/image1.jpg",
		"https://example.com/image2.jpg",
		"https://example.com/image3.jpg",
		"https://example.com/image4.jpg",
		"https://example.com/image5.jpg",
		"https://example.com/image6.jpg",
		"https://example.com/image7.jpg",
		"https://example.com/image8.jpg",
		"https://example.com/image9.jpg",
		"https://example.com/image10.jpg",
	}

	var wg sync.WaitGroup
	totalWorkers := 5

	resultChan := make(chan Result, 50)
	jobsChan := make(chan string, len(jobs))
	start := time.Now()

	for i := 1; i <= totalWorkers; i++ {
		wg.Add(1)
		go worker(jobsChan, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for i := 0; i < len(jobs); i++ {
		jobsChan <- jobs[i]
	}
	close(jobsChan)

	for result := range resultChan {
		fmt.Println("Result received:", result)
	}

	fmt.Println("Time taken:", time.Since(start))
}

// This program demonstrates the use of goroutines and channels in Go to process multiple images concurrently. Each worker simulates processing an image by sleeping for 50 milliseconds and then sends the result back through a channel. The main function waits for all workers to finish and collects their results.
// Worker Pooling is a common pattern in Go for managing concurrent tasks efficiently. In this example, we create a fixed number of workers (5) that process jobs from a channel. The main function sends jobs to the channel and collects results from another channel. This allows for efficient concurrent processing of tasks. Note that the program uses a WaitGroup to ensure that all workers have completed their tasks before closing the result channel.
