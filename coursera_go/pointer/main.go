package main

import "fmt"

func main() {
	var i int = 10
	var p *int

	fmt.Printf("i = %d, %T\n", i, i)
	p = &i
	fmt.Printf("p = %p, %T\n", p, p)
	*p = 15
	fmt.Printf("i = %d, %T\n", i, i)
	fmt.Printf("p = %p, %T\n", p, p)

}
