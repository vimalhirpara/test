package main

import "fmt"

type student struct {
	fname string
	lname string
	age   int
}

func main() {
	fmt.Println("Test 10")
	std := student{fname: "vimal", lname: "hirapara", age: 25}
	fmt.Printf("%v", std)
}
