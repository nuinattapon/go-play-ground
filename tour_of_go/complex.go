package main

import (
	"fmt"
)

func main() {
	var n complex128 = 1 + 2i

	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", real(n), real(n))
	fmt.Printf("%v, %T\n", imag(n), imag(n))

	a := 1 + 2i
	// b := 2 + 5.2i
	b := complex(2, 5.2)
	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n", b, b)

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
}
