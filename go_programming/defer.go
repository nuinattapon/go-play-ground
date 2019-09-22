package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("HTTP Response Code %d - %s\n\n", res.StatusCode, http.StatusText(res.StatusCode))
	defer res.Body.Close() // make sure that the resource will be closed()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
