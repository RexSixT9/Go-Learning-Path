// Goroutines are lightweight threads managed by the Go runtime. They allow concurrent execution of functions, making it easier to write concurrent programs. In this example, we will demonstrate how to create and run goroutines in Go. Go provides the `go` keyword to start a new goroutine. When you call a function with the `go` keyword, it runs concurrently with the calling function.

// package main

// import (
// 	"fmt"
// 	// "time"
// )

// func task(id int) {
// 	fmt.Printf("Task %d is running\n", id)
// }

// func main() {
// 	for i := 1; i <= 5; i++ {
// 		go task(i) // Start a new goroutine for each task
// 	}
// 	// time.Sleep(1 * time.Second) // Sleep for a second to allow goroutines to finish (not a recommended practice)
// }

// Go Concurrency vs Parallelism := Concurrency is the ability of a program to manage multiple tasks at the same time, while parallelism is the ability to execute multiple tasks simultaneously. In Go, goroutines provide concurrency, allowing you to write programs that can handle multiple tasks concurrently without blocking the main execution flow.

package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	go func() {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("Goroutine 1 finished", time.Since(start))
	}()

	go func() {
		time.Sleep(150 * time.Millisecond)
		fmt.Println("Goroutine 2 finished", time.Since(start))
	}()

	fmt.Println("Main function finished", time.Since(start))

	fmt.Println("Waiting...", time.Since(start))
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Waiting...", time.Since(start))
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Waiting...", time.Since(start))
	time.Sleep(500 * time.Millisecond)

	fmt.Println("All goroutines finished", time.Since(start))
}
