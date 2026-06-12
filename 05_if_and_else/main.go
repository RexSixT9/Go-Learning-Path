package main

import "fmt"

func main() {

	score := 85

	// Using if-else statements to evaluate the score and print feedback
	if score >= 90 {
		fmt.Println("Excellent!")
	} else if score >= 80 {
		fmt.Println("Good job!")
	} else if score >= 70 {
		fmt.Println("Not bad!")
	} else {
		fmt.Println("Needs improvement.")
	}

	item := 5
	pricePerItem := 20

	// Using an if statement with a short variable declaration to calculate total cost
	if total := item * pricePerItem; total >= 100 {
		fmt.Println("Total cost is", total, ", which is expensive.")
	} else {
		fmt.Println("Total cost is", total, ", which is affordable.")
	}

	// Using an if statement with a short variable declaration to check if a number is even or odd
	if num := 10; num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	age := 20
	citizen := true

	if age >= 18 && citizen {
		fmt.Println("Eligible")
	} else {
		fmt.Println("Not eligible")
	}
}
