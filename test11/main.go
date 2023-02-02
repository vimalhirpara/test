package main

import "fmt"

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	var n int
	fmt.Print("Enter the number of terms: ")
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", Fibonacci(i))
	}
}
