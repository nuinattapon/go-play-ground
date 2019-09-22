package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	aDir := ""
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No directory provided so we use the current directory")
		aDir = "."
	} else {
		aDir = args[0]
	}

	files, err := ioutil.ReadDir(aDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		//print file which its size is not zero
		switch {
		case file.IsDir():
			fmt.Println(file.Name() + " is an directory")
		case file.Size() != 0:
			fmt.Println(file.Name() + " is not an empty file")
		default:
			fmt.Println(file.Name() + " is an empty file")

		}

	}
}
