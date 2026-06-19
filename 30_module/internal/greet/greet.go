package greet

import "strings"

// Exported function that returns a greeting message for the given name.
func Hello(name string) string {
	clean := normalizeName(name)
	return "Hello, " + clean + "!"
}

// normalizeName trims whitespace from the name and converts it to uppercase.
func normalizeName(name string) string {
	n := strings.TrimSpace(name)

	if n == "" {
		return "Guest"
	}

	return strings.ToUpper(n)
}
