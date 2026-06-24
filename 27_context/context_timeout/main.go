package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Context Timeout is a powerful tool in Go for managing deadlines, cancellation signals, and request-scoped values across API boundaries and between processes. It allows you to control the lifecycle of operations and manage resources effectively. In this example, we will demonstrate how to use context with a timeout to manage a long-running operation and handle cancellation.

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel() // Ensure that the cancel function is called to release resources when the context is no longer needed

	go slowWork(ctx)

	<-ctx.Done() // Wait for the context to be done (either timeout or cancellation)
	fmt.Println("Main function exiting:", ctx.Err())
	fmt.Println("Main function completed.")
}

func slowWork(ctx context.Context) {
	select {
	case <-time.After(1 * time.Second): // Simulate a slow operation that takes 1 second to complete
		return
	case <-ctx.Done(): // Check if the context has been cancelled or timed out
		return
	}

}
