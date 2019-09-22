package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		fmt.Printf("%4d\n", sum)
		sum *= 2
	}
	fmt.Printf("%4d\n", sum)
}
