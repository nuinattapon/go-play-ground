package main

import (
    "fmt"
    "net/http"
)

// CREATE A HTTP SERVER

// http.ResponseWriter assembles the servers response and writes to 
// the client
// http.Request is the clients request
func handler(w http.ResponseWriter, r *http.Request) {

	// Writes to the client
    fmt.Fprintf(w, "Hello World\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Earth\n")
}

func main() {

	// Calls for function handlers output to match the directory /
    http.HandleFunc("/", handler)
    
    // Calls for function handler2 output to match directory /earth
    http.HandleFunc("/earth", handler2)package main

import "fmt"
import "time"

// GO ROUTINES

func count(id int) {
	for i := 0; i < 10; i++ { 
		fmt.Println(id, ":", i) 
		
		// Pause the function for 1 second to allow other functions to execute
		time.Sleep(time.Millisecond * 1000)
	} 
} 

func main() { 

	// A go routine is a function that runs in parallel with other functions
	// We define one by using go followed by the function name
	
	for i := 0; i < 10; i++ {
		go count(i) 
	}
	
	// Wait for the timer to make sure the go routine has time to 
	// finish otherwise the program would end before that happens
	time.Sleep(time.Millisecond * 11000)
}
    
    // Listen to port 8080 and handle requests
    http.ListenAndServe(":8080", nil)
}