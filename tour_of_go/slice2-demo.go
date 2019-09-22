package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:6]

	fmt.Printf("a = %v, %T\n", a, a)
	fmt.Printf("b = %v, %T\n", b, b)
	fmt.Printf("c = %v, %T\n", c, c)
	fmt.Printf("d = %v, %T\n", d, d)
	fmt.Printf("e = %v, %T\n\n", e, e)

	a[5] = 42
	fmt.Printf("a[5] = %v, %T\n", a[5], a[5])

	fmt.Printf("a = %v, %T\n", a, a)
	fmt.Printf("b = %v, %T\n", b, b)
	fmt.Printf("c = %v, %T\n", c, c)
	fmt.Printf("d = %v, %T\n", d, d)
	fmt.Printf("e = %v, %T\n", e, e)
}
