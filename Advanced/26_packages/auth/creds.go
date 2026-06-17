package auth

import "github.com/fatih/color"

func Login(username, password string) {
	color.Red("Logging in with username: %s and password: %s", username, password)
}
