package main

import "fmt"

// A closure is a function that captures and retains access to variables from its surrounding scope, even after that scope has finished executing. In Go, closures are created when you define a function inside another function and the inner function references variables from the outer function.

func counter() func() int {
	count := 0
	return func() int { // This inner function is a closure that captures the 'count' variable from the outer function.
		count += 1
		return count
	}
}

func main() {
	countFunc := counter() // 'countFunc' is now a closure that retains access to the 'count' variable defined in 'counter'.
	fmt.Println(countFunc())
	fmt.Println(countFunc())
	fmt.Println(countFunc())
}
