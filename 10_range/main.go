package main

import "fmt"

func main() {
	views := []int{10, 20, 30, 40, 50}

	// Using range to iterate over the slice
	for i, v := range views {
		fmt.Println(i, v)
	}
}
