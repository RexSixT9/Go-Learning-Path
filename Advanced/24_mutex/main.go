package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex // Mutex to protect the views field
}

func (p *post) incrementViews(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()     // Signal that this goroutine is done
		p.mu.Unlock() // Unlock the mutex after incrementing views
	}()

	p.mu.Lock() // Lock the mutex before accessing the views field
	p.views += 1
}

func main() {

	var wg sync.WaitGroup // WaitGroup to wait for all goroutines to finish
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1) // Increment the WaitGroup counter for each goroutine
		myPost.incrementViews(&wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Printf("Post has %d views\n", myPost.views)
}

// Mutexes are used to protect shared data from being accessed by multiple goroutines at the same time. In this example, we have a simple post struct that has a views field. We increment the views field in a loop, but since we are not using any synchronization mechanism, this code is not safe for concurrent access. If multiple goroutines were to call incrementViews() simultaneously, it could lead to a race condition and incorrect results. To fix this, we would need to use a mutex to lock the incrementViews() method, ensuring that only one goroutine can access it at a time.

// Detailed Desc about working pattern and workflow of the code:
// 1. We define a struct called post that has an integer field views and a mutex field mu. The mutex is used to protect the views field from concurrent access.
// 2. We define a method incrementViews() on the post struct that takes a pointer to a WaitGroup as an argument. This method is responsible for incrementing the views field safely.
// 3. Inside the incrementViews() method, we first lock the mutex to ensure that only one goroutine can access the views field at a time. We then increment the views field and unlock the mutex after we are done. We also call wg.Done() to signal that this goroutine has finished its work.
// 4. In the main function, we create a WaitGroup and an instance of the post struct. We then start a loop that spawns 100 goroutines, each calling the incrementViews() method on the post instance.
// 5. After starting all the goroutines, we call wg.Wait() to block the main goroutine until all the incrementViews() goroutines have finished executing.
// 6. Finally, we print the total number of views, which should be 100 if everything is working correctly.
