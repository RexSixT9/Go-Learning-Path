package main

import "fmt"

type age int

func main() {

	var FirstName string = "John"
	var LastName string = "Doe"
	var Age age = 30
	var IsMarried bool = true
	var City, Country string = "New York", "USA"

	fmt.Println("First Name:", FirstName)
	fmt.Println("Last Name:", LastName)
	fmt.Println("Age:", Age)
	fmt.Println("Is Married:", IsMarried)
	fmt.Println("City:", City)
	fmt.Println("Country:", Country)
}
