// WaitGroup is a synchronization primitive in Go that allows you to wait for a collection of goroutines to finish. It provides a way to block the main goroutine until all the other goroutines have completed their tasks. Here's an example of how to use WaitGroup in Go:
// In this example, we create a WaitGroup and launch multiple goroutines to perform tasks concurrently. Each goroutine calls wg.Done() when it finishes its task, and the main goroutine waits for all of them to complete using wg.Wait().
package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Doing task %d\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()
}
