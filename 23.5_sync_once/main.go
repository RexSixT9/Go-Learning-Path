package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var once sync.Once

	workers := 5
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			// The Do method of the sync.Once type ensures that the function passed to it is executed only once, even if called from multiple goroutines. In this example, we are simulating a scenario where multiple goroutines are trying to perform some initialization work, but we want to ensure that the initialization is done only once.

			once.Do(setup)
			fmt.Printf("Worker %d completed\n", workerID)
		}(i)
	}
	wg.Wait()
	fmt.Println("All workers completed")
}

func setup() {
	// This function is called only once by the sync.Once type, even if multiple goroutines call it concurrently. It simulates some setup work that needs to be done only once.
	fmt.Println("Performing setup...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Setup completed")
}
