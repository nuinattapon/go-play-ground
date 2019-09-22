package main

import "fmt"

func main() {
	// Using var
	var isCool = true
	isCool = false

	// Shorthand
	name := "Brad"
	size := 1.3
	age := 47
	fmt.Println(name, age, isCool, size)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", size)
}
