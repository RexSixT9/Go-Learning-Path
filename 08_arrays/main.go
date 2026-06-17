package main

import "fmt"

func main() {
	var array [5]int

	// Assigning values to the array
	array[0] = 10
	array[1] = 20
	array[2] = 30
	array[3] = 40
	array[4] = 50

	// Accessing values from the array
	fmt.Println("Array values:")
	fmt.Println(array)

	// Initializing an array with values
	res := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array values:")
	fmt.Println(res)
}
