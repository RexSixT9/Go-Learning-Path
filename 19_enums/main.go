// Enums in Go are typically implemented using constants and iota. In this example, we define an OrderStatus type and use iota to create a set of constants representing different order statuses. The changeStatus function takes an OrderStatus as an argument and prints the new status. When you run this code, it will output the status changes for each of the defined order statuses.

package main

import "fmt"

type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed OrderStatus = "confirmed"
	Prepared  OrderStatus = "prepared"
	Delivered OrderStatus = "delivered"
)

func changeStatus(status OrderStatus) {
	fmt.Println("Status changed to:", status)
}

func main() {
	changeStatus(Received)
	changeStatus(Confirmed)
	changeStatus(Prepared)
	changeStatus(Delivered)
}
