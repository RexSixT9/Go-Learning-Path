// package main

// import (
// 	"fmt"
// 	"log"
// 	"strconv"
// )

// func main() {
// 	if err := run(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func run() error {
// 	input := "3"
// 	level, err := parseLevel(input)

// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("Parsed level:", level)
// 	return nil

// }

// func parseLevel(level string) (int, error) {
// 	n, err := strconv.Atoi(level)
// 	if err != nil {
// 		return 0, fmt.Errorf("invalid level: %w", err)
// 	}
// 	if n < 1 || n > 10 {
// 		return 0, fmt.Errorf("level out of range: %d", n)
// 	}
// 	return n, nil
// }

// Error handling is a crucial aspect of programming in Go. It allows developers to gracefully handle unexpected situations and provide meaningful feedback to users or other parts of the system. In this example, we will demonstrate how to handle errors effectively in Go by using error wrapping and custom error messages.

// Error Wrapping is a technique that allows you to add context to an error while preserving the original error. This is particularly useful when you want to provide more information about where and why an error occurred, making it easier to debug and understand the root cause of the issue.

// package main

// import "fmt"

// func firstFunction() error {
// 	return fmt.Errorf("an error occurred in firstFunction")
// }

// func secondFunction() error {
// 	firstErr := firstFunction()
// 	if firstErr != nil {
// 		return fmt.Errorf("secondFunction encountered an error: %w", firstErr)
// 	}
// 	return nil
// }

// func main() {
// 	err := secondFunction()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// }

// Unwrapping an error allows you to retrieve the original error that was wrapped. This can be done using the errors.Unwrap function or by using the errors.Is and errors.As functions to check for specific error types.

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func firstFunction() error {
// 	return fmt.Errorf("an error occurred in firstFunction")
// }

// func secondFunction() error {
// 	firstErr := firstFunction()
// 	if firstErr != nil {
// 		return fmt.Errorf("secondFunction encountered an error: %w", firstErr)
// 	}
// 	return nil
// }

// func main() {
// 	err := secondFunction()
// 	fmt.Println("Original Error:", err)

// 	innerErr := errors.Unwrap(err)
// 	if innerErr != nil {
// 		fmt.Println("Unwrapped Error:", innerErr)
// 	}

// }

// CustomError is a user-defined error type that wraps another error and provides additional context. It implements the error interface by defining the Error() method, which returns a formatted string containing the custom message and the wrapped error. The Unwrap() method allows access to the underlying error, enabling error unwrapping and inspection.

package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Message string
	Wrapped error
}

func (e CustomError) Error() string {
	return fmt.Sprintf("CustomError: %s, Wrapped: %v", e.Message, e.Wrapped)
}

func (e CustomError) Unwrap() error {
	return e.Wrapped
}

func someFunction() error {
	return CustomError{
		Message: "Something went wrong",
		Wrapped: fmt.Errorf("original error"),
	}
}

func main() {
	err := someFunction()
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = errors.Unwrap(err)
	innerErr := fmt.Errorf("Innermost: %w", err)
	fmt.Println("Inner Error:", innerErr)
}
