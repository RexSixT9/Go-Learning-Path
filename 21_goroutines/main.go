// This program demonstrates the use of goroutines in Go. It starts multiple tasks concurrently using goroutines. Each task prints its ID to the console. However, since the main function does not wait for the goroutines to finish, the program may exit before all tasks have completed. To ensure that all tasks complete before the program exits, you can use synchronization techniques such as WaitGroups or channels. In this example, we simply start the goroutines without waiting for them, which may lead to some tasks not being printed if the main function exits too quickly.

package main

import (
	"fmt"
	// "time"
)

func task(id int) {
	fmt.Printf("Task %d is running\n", id)
}

func main() {
	for i := 1; i <= 5; i++ {
		go task(i) // Start a new goroutine for each task
	}
	// time.Sleep(1 * time.Second) // Sleep for a second to allow goroutines to finish (not a recommended practice)
}
