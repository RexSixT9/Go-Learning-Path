package main

import "fmt"

func main() {

	score := 100
	fmt.Println("Initial score:", score)

	addScore(&score)
	fmt.Println("Final score:", score)
}

// Pointers are variables that store the memory address of another variable. In Go, you can use pointers to pass references to values and to allow functions to modify the original variable. The `&` operator is used to get the memory address of a variable, and the `*` operator is used to dereference a pointer, which means accessing the value stored at that memory address.

func addScore(score *int) {
	*score = *score + 10
}
