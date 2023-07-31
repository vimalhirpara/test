package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a ticker that ticks every 10 seconds
	ticker := time.NewTicker(1 * time.Second)

	// Run the loop to print "Hello, World!" on each tick
	for range ticker.C {
		fmt.Println("Hello, World!")
	}
}
