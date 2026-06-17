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

package main

import (
	"fmt"
	"time"
)

func main() {
	messageCh := make(chan string)

	go func() {
		time.Sleep(10 * time.Second)
		messageCh <- "Hello!"
	}()

	select {
	case msg := <-messageCh:
		fmt.Println("Received:", msg)

	case <-time.After(5 * time.Second):
		fmt.Println("User timed out")
	}
}
