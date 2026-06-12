package main

import "fmt"

func main() {
	var FirstName string = "John"      // Standard variable declaration with type and initialization
	LastName := "Doe"                  // Short variable declaration
	var Age = 30                       // Type inference, the type is inferred from the assigned value
	City, Country := "New York", "USA" // Short variable declaration for multiple variables

	fmt.Println("First Name:", FirstName)
	fmt.Println("Last Name:", LastName)
	fmt.Println("Age:", Age)
	fmt.Println("City:", City)
	fmt.Println("Country:", Country)

	// Short variable declaration for string concatenation
	petName := "Buddy"
	ownerName := "Alice"
	petsAndOwners := petName + " is owned by " + ownerName

	fmt.Println(petsAndOwners)

	// Short variable declaration for a boolean variable
	isMarried := true
	fmt.Println("Is Married:", isMarried)

	// Short variable declaration for boolean variables
	isSubscribed := false
	isAdmin := true
	isLoggedIn := true

	// Short variable declaration for a boolean expression
	isDashboardOpen := isLoggedIn && isSubscribed
	canDelete := isAdmin || (isLoggedIn && isSubscribed)

	fmt.Println("Is Dashboard Open:", isDashboardOpen)
	fmt.Println("Can Delete:", canDelete)

	// variable declaration for constants.
	// Constants are declared using the const keyword and cannot be changed after they are set.
	const maxUsers = 100
	const Pi float64 = 3.14

	fmt.Println("Max Users:", maxUsers)
	fmt.Println("Pi:", Pi)

}
