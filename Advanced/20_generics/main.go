// Generics are a way to write flexible and reusable code in Go. They allow you to define functions, types, and data structures that can work with different data types without sacrificing type safety. Generics were introduced in Go 1.18, and they provide a way to create functions and types that can operate on any type that satisfies certain constraints.
// T is a type parameter that can be any type (indicated by 'any'). The function takes a slice of type T as an argument.

package main

import "fmt"

func printSlice[T any](item []T) {
	for _, v := range item {
		fmt.Println(v)
	}
}

func main() {
	items := []int{1, 2, 3, 4, 5}
	names := []string{"Alice", "Bob", "Charlie"}
	printSlice(items)
	printSlice(names)
}
