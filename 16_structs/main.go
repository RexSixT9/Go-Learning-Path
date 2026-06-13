package main

import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

// In this example, we define a struct type called 'User' with three fields: 'ID', 'Name', and 'Age'. We then create instances of the 'User' struct, modify one of the fields, and print the results to demonstrate how structs work in Go. Structs are a powerful way to group related data together and can be used to create complex data structures in Go.

func main() {
	user1 := User{
		ID:   1,
		Name: "Alice",
		Age:  30,
	}
	fmt.Println("User:", user1)

	user1.Age = 31
	fmt.Println("Updated User:", user1)

	user2 := User{ID: 2, Name: "Bob"}
	fmt.Println("User 2:", user2)
}
