package main

import "fmt"

func main() {

	// Creating a slice using a slice literal
	res := []int{1, 2, 3}
	fmt.Println(res)

	// Using make to create a slice
	score := make([]int, 0, 5)
	score = append(score, 10)
	score = append(score, 20)
	score = append(score, 30)

	fmt.Println(score, len(score), cap(score))
}
