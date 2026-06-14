package main

import "fmt"

// In Go, an interface is a type that defines a set of method signatures. It specifies a contract that any type must fulfill to be considered an implementation of that interface. An interface allows us to define behavior without specifying the underlying data structure, enabling polymorphism and decoupling code from specific implementations.

// In this example, we have defined an interface called paymenter that has a single method makePayment. Any type that implements this method can be considered a paymenter. We have two concrete types, razorpay and stripe, that implement the makePayment method. The payment struct uses the paymenter interface to process payments through different gateways without being tightly coupled to any specific implementation. This design allows for flexibility and extensibility, as we can easily add new payment gateways by implementing the paymenter interface without modifying the existing code.

// The commented-out code for fakePayment demonstrates that if we try to use a type that does not implement the paymenter interface, it will result in a compile-time error, ensuring type safety and adherence to the defined contract.

type paymenter interface {
	makePayment(amount float64)
	refundPayment(amount float64, account string)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float64) {
	p.gateway.makePayment(amount)
}

func (p payment) refundPayment(amount float64, account string) {
	p.gateway.refundPayment(amount, account)
}

type razorpay struct {
}

func (r razorpay) makePayment(amount float64) {
	fmt.Printf("Processing Razorpay payment of amount: %.2f\n", amount)
}

func (r razorpay) refundPayment(amount float64, account string) {
	fmt.Printf("Refunding Razorpay payment of amount: %.2f for account: %s\n", amount, account)
}

type stripe struct {
}

func (s stripe) makePayment(amount float64) {
	fmt.Printf("Processing Stripe payment of amount: %.2f\n", amount)
}

func (s stripe) refundPayment(amount float64, account string) {
	fmt.Printf("Refunding Stripe payment of amount: %.2f for account: %s\n", amount, account)
}

// type fakePayment struct{}

// func (f fakePayment) makePayment(amount float64) {
// 	fmt.Printf("Processing Fake payment of amount: %.2f\n", amount)
// }

func main() {
	razorpayGateway := razorpay{}
	stripeGateway := stripe{}
	// fakeGateway := fakePayment{}

	newPayment1 := payment{gateway: razorpayGateway}
	newPayment1.makePayment(100.0)
	newPayment1.refundPayment(50.0, "user123")

	newPayment2 := payment{gateway: stripeGateway}
	newPayment2.makePayment(200.0)
	newPayment2.refundPayment(100.0, "user456")

	// newPayment3 := payment{gateway: fakeGateway}
	// newPayment3.makePayment(300.0)
}
