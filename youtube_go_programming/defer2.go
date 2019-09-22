package main

import "fmt"

func main() {
	a := "Start"
	defer fmt.Println(a)

	a = "Stop"
}
