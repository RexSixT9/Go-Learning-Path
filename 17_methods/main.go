package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "John", Age: 20}
	u.Greet()
	u.IsAdult()
	u.Greet()
}

// The Greet method is defined with a value receiver (User), which means it operates on a copy of the User struct. When we call u.Greet(), it prints the name and age of the user without modifying the original struct.
func (u User) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", u.Name, u.Age)
}

// The IsAdult method is defined with a pointer receiver (*User), which allows it to modify the Age field of the User struct. When we call u.IsAdult(), it increments the Age by 1, demonstrating how methods can change the state of a struct when using pointer receivers.
func (u *User) IsAdult() {
	u.Age++
}
