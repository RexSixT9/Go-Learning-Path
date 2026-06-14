package main

import "fmt"

// Functions are reusable blocks of code that perform a specific task. They can take parameters and return values. In Go, functions are defined using the `func` keyword, followed by the function name, parameters, and return type(s).
func add(a, b int) int {
	return a + b
}

// Functions can also return multiple values. In this example, the function returns both the sum and product of two integers.
func sumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

// Named Return Values: You can name the return values in the function signature. This allows you to return values without explicitly specifying them in the return statement.
func namedReturn(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return
}

// Variadic Functions: A variadic function can take a variable number of arguments. The parameter is defined with an ellipsis (...) before the type.
func sumAll(num ...int) int {
	total := 0
	for _, currentValue := range num {
		total += currentValue
	}
	return total
}

// Anonymous Functions: Functions can also be defined without a name and can be assigned to variables or used as arguments to other functions.
var res = func(a, b int) int {
	return a + b
}

// defer statement is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages. // Deferred function calls are executed in Last In First Out order after the surrounding function returns.
func deferExample() {
	defer fmt.Println("This will be printed last")
	defer fmt.Println("This will be printed second")
	fmt.Println("This will be printed first")
}

func main() {
	result := add(5, 10)
	fmt.Println("The sum is:", result)

	sum, product := sumAndProduct(5, 10)
	fmt.Println("The sum is:", sum)
	fmt.Println("The product is:", product)

	// Ignoring the product value
	onlySum, _ := sumAndProduct(5, 10)
	fmt.Println("The sum is:", onlySum)

	namedSum, namedProduct := namedReturn(5, 10)
	fmt.Println("The sum is:", namedSum)
	fmt.Println("The product is:", namedProduct)

	fmt.Println("The sum of all numbers is:", sumAll(5, 10, 15, 20))

	values := []int{1, 2, 3, 4, 5}
	fmt.Println("The sum of the slice is:", sumAll(values...))

	fmt.Println("The result of the anonymous function is:", res(5, 10))

	var resFunction = func(a, b int) int {
		return a * b
	}(5, 10) // Immediately Invoked Function Expression (IIFE)
	fmt.Println("The result of the IIFE is:", resFunction)

	deferExample()
}
