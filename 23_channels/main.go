// package main

// import "fmt"

// func main() {
// 	messagesChannel := make(chan string) // Create a channel of type string

// 	messagesChannel <- "Hello, World!" // Send a message to the channel

// 	message := <-messagesChannel // Receive a message from the channel
// 	fmt.Println(message)         // Print the received message

// }

// In this example, we create a channel of type string using make(chan string). We then send a message "Hello, World!" to the channel using the <- operator. Finally, we receive the message from the channel and print it to the console. Channels are a powerful feature in Go that allow for communication and synchronization between goroutines. The program will deadlock because the main goroutine is trying to send a message to the channel before any goroutine is ready to receive it. To avoid this, you can use a separate goroutine to send the message or use buffered channels.

/*

package main

import (
	"fmt"
	"time"
)

// Example of using a separate goroutine to send a message to the channel
func processMessage(messagesChannel chan string) {
	for msg := range messagesChannel {
		fmt.Println(msg)
	}
}

// Example of using a channel to calculate the sum of numbers in a separate goroutine
func sum(numbers []int, resultsChannel chan int) {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	resultsChannel <- sum
}

// Example of using a channel to signal when a task is done
func task(done chan bool) {
	fmt.Println("Task starting...")
	defer func() {
		done <- true
	}()
}

// Example of using a channel to send emails in a separate goroutine
func emailSender(emailChannel chan string, done chan bool) {
	defer func() {
		done <- true
	}()
	for email := range emailChannel {
		fmt.Printf("Sending email to %s\n", email)
		time.Sleep(time.Millisecond * 500) // Simulate time taken to send an email
	}
	done <- true
}

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	// Start a separate goroutine to send a message to chan1
	go func() {
		chan1 <- 42
	}()

	// Start a separate goroutine to send a message to chan2
	go func() {
		chan2 <- "Hello, World!"
	}()

	// Receive messages from chan1 and chan2
	for i := 0; i < 2; i++ {
		select {
		case num := <-chan1:
			fmt.Printf("Received from chan1: %d\n", num)
		case msg := <-chan2:
			fmt.Printf("Received from chan2: %s\n", msg)
		}
	}

	// Example of using a channel to send emails in a separate goroutine
	emailChannel := make(chan string, 100) // Create a channel of type string
	done := make(chan bool)                // Create a channel of type bool

	for i := 1; i <= 5; i++ {
		emailChannel <- fmt.Sprintf("test%d@example.com", i)
	}
	close(emailChannel)
	go emailSender(emailChannel, done) // Start a goroutine to send emails

	fmt.Println("Emails queued for sending...")
	<-done // Wait for the email sender to finish

	// Example of using a channel to signal when a task is done
	done = make(chan bool) // Create a channel of type bool
	go task(done)          // Start a goroutine to run the task

	<-done // Wait for the task to complete

	// Example of using a channel to calculate the sum of numbers in a separate goroutine
	resultsChannel := make(chan int)             // Create a channel of type int
	go sum([]int{1, 2, 3, 4, 5}, resultsChannel) // Start a goroutine to calculate the sum
	result := <-resultsChannel                   // Receive the result from the channel
	fmt.Println("Sum:", result)

	// Example of using a separate goroutine to send a message to the channel
	messagesChannel := make(chan string) // Create a channel of type string
	go processMessage(messagesChannel)
	for {
		messagesChannel <- "Hello, World!" // Send a message to the channel
		time.Sleep(time.Second)            // Sleep for a second to allow the goroutine to finish
	}
	// time.Sleep(time.Second)            // Sleep for a second to allow the goroutine to finish
	// messagesChannel <- "Hello, World!" // Send a message to the channel

}
*/

package main

import (
	"fmt"
	"time"
)

type User struct {
	ID   int
	Name string
}

func main() {
	ch := make(chan User) // Create a channel of type User, unbuffered channel
	go func() {
		time.Sleep(200 * time.Millisecond) // Simulate some processing time
		fmt.Println("Sending user to channel...")
		// blocking send operation, will wait until the main goroutine receives the value
		ch <- User{ID: 1, Name: "John"} // Send a User struct to the channel
		fmt.Println("User sent to channel.")
	}()
	fmt.Println("Waiting to receive user from channel...")
	user := <-ch // blocking receive operation, will wait until the goroutine sends the value
	fmt.Printf("Received user: %+v\n", user)
}
