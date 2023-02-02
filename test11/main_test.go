package main

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
	}

	for _, test := range tests {
		result := Fibonacci(test.n)
		if result != test.expected {
			t.Errorf("For n=%d, expected %d but got %d", test.n, test.expected, result)
		} else {
			fmt.Printf("For n=%d, expected %d and got %d\n", test.n, test.expected, result)
		}
	}
}
