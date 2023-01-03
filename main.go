package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 2; i++ {
		drawPattern(i + 1)
	}
}

func drawPattern(c int) {

	var _c int
	var _cm int = 5

	if c == 2 {
		_c = 2
	} else {
		_c = _cm
	}

	for i := 0; i <= _c; i++ {
		for j := 0; j < _cm-i; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Printf("*")
		}
		for j := 0; j <= i; j++ {
			fmt.Printf("*")
		}
		for j := 0; j < _cm-i; j++ {
			fmt.Printf(" ")
		}
		fmt.Printf("\n")
	}
}
