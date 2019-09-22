package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Printf("%2d\n", i)
		sum += i
	}
	fmt.Println("--")
	fmt.Printf("%2d\n", sum)
}
