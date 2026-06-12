package main

import "fmt"

func main() {
	day := 3

	// Switch with an expression
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Invalid day")
	}

	// Switch without an expression
	switch {
	case day == 1:
		fmt.Println("Monday")
	case day == 2:
		fmt.Println("Tuesday")
	case day == 3:
		fmt.Println("Wednesday")
	case day == 4:
		fmt.Println("Thursday")
	case day == 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Invalid day")

	}
}
