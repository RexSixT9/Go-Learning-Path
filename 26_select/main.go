// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	alice := make(chan string)
// 	bob := make(chan string)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		alice <- "Alice calling"
// 	}()

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		bob <- "Bob calling"
// 	}()

// 	select {
// 	case msg := <-alice:
// 		fmt.Println(msg)

// 	case msg := <-bob:
// 		fmt.Println(msg)
// 	}
// }

// ANOTHER EXAMPLE

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	messageCh := make(chan string)

// 	go func() {
// 		time.Sleep(10 * time.Second)
// 		messageCh <- "Hello!"
// 	}()

// 	select {
// 	case msg := <-messageCh:
// 		fmt.Println("Received:", msg)

// 	case <-time.After(5 * time.Second):
// 		fmt.Println("User timed out")
// 	}
// }

// Another Example

package main

import (
	"fmt"
	"time"
)

func main() {
	// select -> allows you to wait on multiple channel operations
	// It blocks until one of the channels is ready for communication and then executes the corresponding case statement. If multiple channels are ready, one of them is chosen at random.
	// It is often used in concurrent programming to handle multiple channels and perform different actions based on which channel receives data first. It can also be used to implement timeouts and cancellation of operations.

	resultChannel := make(chan string)
	go func() {
		time.Sleep(100 * time.Millisecond)
		resultChannel <- "Result"
	}()

	// time.AfterFunc is a function that takes a duration and a function as arguments. It waits for the specified duration and then executes the provided function. In this case, it waits for 1 second and then sends "Timeout" to the resultChannel.

	timeOut := time.After(250 * time.Millisecond)

	select {
	case result := <-resultChannel:
		fmt.Println("Received:", result)
	case <-timeOut:
		fmt.Println("Operation timed out")
	}

	fmt.Println("Main function completed.")
}
