package main

import "fmt"

func main() {

	// Creating a map using a map literal
	ages := map[string]int{
		"john": 20,
		"jane": 30,
		"bob":  35,
	}

	fmt.Println(ages)
	fmt.Println(ages["john"], len(ages))

	var scores map[string]int     // nil map
	scores = make(map[string]int) // initialize the map
	scores["math"] = 90           // add a key-value pair to the map
	scores["english"] = 85
	fmt.Println(scores)

	// check if the key "john" exists in the map
	age, ok := ages["john"]
	if ok {
		fmt.Println("John's age is", age, ok)
	} else {
		fmt.Println("John's age not found", ok)
	}
}
