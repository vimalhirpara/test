package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {

	// Read environment file
	env, _ := godotenv.Read(".env")

	// Print environment variable value
	fmt.Printf("Username: %s\n", env["uname"])
	fmt.Printf("Password: %s\n", env["pass"])
}
