// panic and recover are built-in functions in Go that allow you to handle unexpected errors and recover from panics. A panic is a runtime error that stops the normal execution of a program, while recover allows you to regain control and continue execution after a panic has occurred. In this example, we will demonstrate how to use panic and recover to handle errors gracefully.

/*
package main

import "fmt"

func function1() {
	defer func() {
		fmt.Println("Function1 defer called")
	}()
}

func function2() {
	defer func() {
		fmt.Println("Function2 defer called")
	}()
	panic("A panic occurred in function2")
}

func main() {
	function1()
	function2()
}
*/

// In this example, we have two functions, function1 and function2. The main function calls both of these functions. The function2 will panic, which will stop the normal execution of the program. However, the defer statements in both functions will still be executed, allowing us to see the output from the defer statements before the program terminates due to the panic.

// To handle the panic and recover from it, we can use the recover function inside a deferred function. This allows us to catch the panic and continue execution of the program instead of terminating it. Here's how we can modify the previous example to include panic recovery:

package main

import "fmt"

func panicFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("A panic occurred in panicFunction")
}

func main() {
	fmt.Println("Program starts")
	panicFunction()
	fmt.Println("Program continues after recovering from panic")
}
