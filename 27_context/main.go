// Context is a powerful tool in Go for managing deadlines, cancellation signals, and request-scoped values across API boundaries and between processes. It allows you to control the lifecycle of operations and manage resources effectively. In this example, we will demonstrate how to use context to manage a long-running operation and handle cancellation.

package main

import (
	"context"
	"fmt"
	"time"
)

// simulateLongRunningTask simulates a long-running operation that checks for cancellation signals from the provided context. It prints "Working..." every second until the context is cancelled.
func simulateLongRunningTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task cancelled:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {

	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure that the cancel function is called to release resources when the context is no longer needed

	go simulateLongRunningTask(ctx) // Start the long-running task in a separate goroutine

	time.Sleep(2 * time.Second) // Simulate some work in the main function

	// Cancel the context to stop the long-running task
	fmt.Println("Cancelling context...")
	cancel()

	// Wait for a moment to observe the cancellation effect
	time.Sleep(1 * time.Second)
}
