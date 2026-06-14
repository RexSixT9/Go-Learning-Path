package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	input := "3"
	level, err := parseLevel(input)

	if err != nil {
		return err
	}
	fmt.Println("Parsed level:", level)
	return nil

}

func parseLevel(level string) (int, error) {
	n, err := strconv.Atoi(level)
	if err != nil {
		return 0, fmt.Errorf("invalid level: %w", err)
	}
	if n < 1 || n > 10 {
		return 0, fmt.Errorf("level out of range: %d", n)
	}
	return n, nil
}
