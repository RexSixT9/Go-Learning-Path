package main

import (
	"fmt"

	"github.com/rexsixt9/go-module/internal/greet"
)

func main() {
	msg := greet.Hello("Alice")
	fmt.Println(msg)
}
