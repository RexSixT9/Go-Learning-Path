/*
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

	for i := 0; i < 200; i++ {
		wg.Add(1) // Increment the WaitGroup counter for each goroutine
		go myPost.incrementViews(&wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Printf("Post has %d views\n", myPost.views)
}
*/

// Mutexes are used to protect shared data from being accessed by multiple goroutines at the same time. In this example, we have a simple post struct that has a views field. We increment the views field in a loop, but since we are not using any synchronization mechanism, this code is not safe for concurrent access. If multiple goroutines were to call incrementViews() simultaneously, it could lead to a race condition and incorrect results. To fix this, we would need to use a mutex to lock the incrementViews() method, ensuring that only one goroutine can access it at a time.

// Detailed Desc about working pattern and workflow of the code:
// 1. We define a struct called post that has an integer field views and a mutex field mu. The mutex is used to protect the views field from concurrent access.
// 2. We define a method incrementViews() on the post struct that takes a pointer to a WaitGroup as an argument. This method is responsible for incrementing the views field safely.
// 3. Inside the incrementViews() method, we first lock the mutex to ensure that only one goroutine can access the views field at a time. We then increment the views field and unlock the mutex after we are done. We also call wg.Done() to signal that this goroutine has finished its work.
// 4. In the main function, we create a WaitGroup and an instance of the post struct. We then start a loop that spawns 100 goroutines, each calling the incrementViews() method on the post instance.
// 5. After starting all the goroutines, we call wg.Wait() to block the main goroutine until all the incrementViews() goroutines have finished executing.
// 6. Finally, we print the total number of views, which should be 100 if everything is working correctly.

// RWMutexes are a type of mutex that allows multiple readers to access a resource simultaneously, but only one writer can access the resource at a time. This is useful when you have a scenario where there are many read operations and few write operations, as it can improve performance by allowing multiple readers to access the resource concurrently while still ensuring that writers have exclusive access when needed. In Go, you can use the sync.RWMutex type to implement read-write locks. The RWMutex has two main methods: RLock() and RUnlock() for read operations, and Lock() and Unlock() for write operations. When a goroutine calls RLock(), it can access the resource as long as there are no writers holding the lock. If a goroutine calls Lock(), it will block until all readers have released their locks and no other writers are holding the lock. This allows for efficient concurrent access to shared resources while maintaining data integrity.

package main

import (
	"fmt"
	"sync"
)

type Post struct {
	views int
	likes int
	mu    sync.RWMutex
}

func (p *Post) View(wg *sync.WaitGroup) {
	defer wg.Done()

	p.mu.Lock()
	defer p.mu.Unlock()

	p.views++
}

func (p *Post) Like(wg *sync.WaitGroup) {
	defer wg.Done()

	p.mu.Lock()
	defer p.mu.Unlock()

	p.likes++
}

func (p *Post) Stats(wg *sync.WaitGroup) {
	defer wg.Done()

	p.mu.RLock()
	defer p.mu.RUnlock()

	fmt.Printf("Views: %d, Likes: %d\n", p.views, p.likes)
}

func main() {
	var wg sync.WaitGroup

	post := Post{}

	for i := 0; i < 200; i++ {
		wg.Add(1)
		go post.View(&wg)
	}

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go post.Like(&wg)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go post.Stats(&wg)
	}

	wg.Wait()

	fmt.Println("Final Stats")

	post.mu.RLock()
	fmt.Printf("Views: %d, Likes: %d\n", post.views, post.likes)
	post.mu.RUnlock()
}
