package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Pokemon struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Result    `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {
	data, err := ioutil.ReadFile("pokemon.json")

	if err != nil {
		fmt.Println(err)
	}

	p := Pokemon{}
	err = json.Unmarshal(data, &p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", p)
}
