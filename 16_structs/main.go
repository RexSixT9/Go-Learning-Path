// In this example, we define a struct type called 'User' with three fields: 'ID', 'Name', and 'Age'. We then create instances of the 'User' struct, modify one of the fields, and print the results to demonstrate how structs work in Go. Structs are a powerful way to group related data together and can be used to create complex data structures in Go. We also define another struct called 'Order' to show how we can have multiple struct types in the same program, each with its own set of fields.

package main

import (
	"fmt"
	"time"
)

// type User struct {
// 	ID   int
// 	Name string
// 	Age  int
// }

// type Order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time
// }

type customer struct {
	name        string
	phoneNumber string
}

// The Order struct includes an embedded struct called 'customer', which allows us to represent the details of the customer associated with the order. This demonstrates how we can use composition in Go to create more complex data structures by embedding one struct within another.

type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // Nanosecond precision timestamp
	updatedAt time.Time // Track when the order was last modified
	customer  customer  // Embedded struct to represent the customer details associated with the order
}

// func (o *Order) changeStatus(newStatus string) {
// 	o.status = newStatus
// }

// func (o Order) getAmount() float32 {
// 	return o.amount
// }

// func NewOrder(id string, amount float32, status string) *Order {
// 	myOrder := Order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &myOrder
// }

func main() {
	// user1 := User{
	// 	ID:   1,
	// 	Name: "Alice",
	// 	Age:  30,
	// }
	// fmt.Println("User:", user1)

	// user1.Age = 31
	// fmt.Println("Updated User:", user1)

	// user2 := User{ID: 2, Name: "Bob"}
	// fmt.Println("User 2:", user2)

	// order1 := Order{
	// 	id:     "ORD123",
	// 	amount: 99.99,
	// 	status: "pending",
	// }
	// order1.createdAt = time.Now()
	// order1.changeStatus("completed")
	// fmt.Println("Order Amount:", order1.getAmount())

	// fmt.Println("Order:", order1)

	// myOrder := NewOrder("ORD456", 149.99, "pending")
	// fmt.Println("My Order:", myOrder)

	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{
	// 	name:   "Go",
	// 	isGood: true,
	// }
	// fmt.Println("LanguageName:", language.name, "\nIsGood:", language.isGood)

	newCustomer := customer{
		name:        "Alice",
		phoneNumber: "987-654-3210",
	}
	myOrder := Order{
		id:        "ORD789",
		amount:    199.99,
		status:    "pending",
		customer:  newCustomer,
		createdAt: time.Now(),
		updatedAt: time.Now(), // Add an updatedAt field to track when the order was last modified
	}

	fmt.Printf("Order ID: %s\nAmount: %.2f\nStatus: %s\nCustomer Name: %s\nCustomer Phone: %s\nCreated At: %s\nUpdated At: %s\n",
		myOrder.id, myOrder.amount, myOrder.status, myOrder.customer.name, myOrder.customer.phoneNumber, myOrder.createdAt, myOrder.updatedAt)

	myOrder.status = "completed"
	myOrder.customer.phoneNumber = "123-456-7890"
	myOrder.updatedAt = time.Now() // Update the updatedAt field when the order is modified

	fmt.Printf("Order ID: %s\nAmount: %.2f\nStatus: %s\nCustomer Name: %s\nCustomer Phone: %s\nCreated At: %s\nUpdated At: %s\n",
		myOrder.id, myOrder.amount, myOrder.status, myOrder.customer.name, myOrder.customer.phoneNumber, myOrder.createdAt, myOrder.updatedAt)

}
