package main

import (
	"fmt"
)

func main() {
	message := "Hello, World!"
	var w Write = ConsoleWriter{}
	n, err := w.Write([]byte(message))
	fmt.Printf("%d, %v\n", n, err)
	fmt.Printf("%d\n", len(message))
}

type Write interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
