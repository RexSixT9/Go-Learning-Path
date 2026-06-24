package main

import (
	"fmt"
	"time"
)

func main() {
	// unbuffered channel is a channel that has no capacity to hold values. It requires both a sender and a receiver to be ready at the same time for communication to occur. handshake between sender and receiver. while buffered channel has a capacity to hold values, allowing the sender to send values without waiting for the receiver to be ready. buffered channels can store multiple values, allowing for asynchronous communication between goroutines.

	// Create a buffered channel with a capacity of 2, 2 means the channel can hold up to 2 values without blocking the sender. If the channel is full, the sender will block until there is space available in the channel.

	jobs := make(chan string, 3)

	go func() {
		fmt.Println("Sending jobs to channel...")
		fmt.Println("Sending job 1 to channel...")
		jobs <- "Job 1"
		fmt.Println("Sending job 2 to channel...")
		jobs <- "Job 2"
		fmt.Println("Sending job 3 to channel...")
		jobs <- "Job 3"
		fmt.Println("All jobs sent to channel.")
		close(jobs) // Close the channel after sending all jobs
	}()

	// Receive jobs from the channel
	for job := range jobs {
		fmt.Println("Received:", job)
		time.Sleep(300 * time.Millisecond) // Simulate time taken to process the job

		fmt.Println("Processing job:", job)
	}
	fmt.Println("All jobs processed.")
}
