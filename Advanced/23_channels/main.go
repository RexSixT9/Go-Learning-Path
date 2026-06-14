// package main

// import "fmt"

// func main() {
// 	messagesChannel := make(chan string) // Create a channel of type string

// 	messagesChannel <- "Hello, World!" // Send a message to the channel

// 	message := <-messagesChannel // Receive a message from the channel
// 	fmt.Println(message)         // Print the received message

// }

// In this example, we create a channel of type string using make(chan string). We then send a message "Hello, World!" to the channel using the <- operator. Finally, we receive the message from the channel and print it to the console. Channels are a powerful feature in Go that allow for communication and synchronization between goroutines.

// Note: In this example, the program will deadlock because the main goroutine is trying to send a message to the channel before any goroutine is ready to receive it. To avoid this, you can use a separate goroutine to send the message or use buffered channels.

// Here's an example of using a separate goroutine to send a message to the channel:

package main

import "fmt"

func main() {
	messagesChannel := make(chan string) // Create a channel of type string
	go func() {
		messagesChannel <- "Hello, World!" // Send a message to the channel
	}()
	message := <-messagesChannel // Receive a message from the channel
	fmt.Println(message)         // Print the received message
}
