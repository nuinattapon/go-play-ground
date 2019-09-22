package main

import "fmt"

func main() {
	a := make([]int, 3, 100)

	fmt.Printf("a = %v, %T\n", a, a)
	a = append(a, 1, 2, 3, 4)
	fmt.Printf("a = %v, %T\n", a, a)
	fmt.Printf("len(a) = %v\n", len(a))
	fmt.Printf("cap(a) = %v\n", cap(a))

}
