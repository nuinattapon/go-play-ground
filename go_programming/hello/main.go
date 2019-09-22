package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now() // start time
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	fmt.Fprintln(w, "Hello, 世界, สวัสดี from Go")
	hostName, _ := os.Hostname()
	fmt.Fprintf(w, "Server Name: %s\n", hostName)
	addrs, _ := net.LookupIP(hostName)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			fmt.Fprintf(w, "Server Addr: %s:8888\n", ipv4)
		}
	}
	fmt.Fprintln(w, "Remote Addr:", r.RemoteAddr)
	fmt.Fprintln(w, "Method:", r.Method)
	fmt.Fprintln(w, "URI:", r.URL.RequestURI())
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		startTime.Year(), startTime.Month(), startTime.Day(),
		startTime.Hour(), startTime.Minute(), startTime.Second())
	fmt.Fprintln(w, "Date:", formatted)

	elapsed := float64(time.Since(startTime).Microseconds()) / 1000.0
	fmt.Printf("%s %s %v %.2f ms\n", r.Method, r.URL.RequestURI(), 200, elapsed)
}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now() // start time

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	user := User{
		Id:    1,
		Name:  "Nattapon Sub-Anake",
		Email: "nattapon.subanake@gmail.com",
		Phone: "66815517772",
	}
	json.NewEncoder(w).Encode(user)
	elapsed := float64(time.Since(startTime).Microseconds()) / 1000.0
	fmt.Printf("%s %s %v %.2f ms\n", r.Method, r.URL.RequestURI(), 200, elapsed)
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now() // start time

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	temp, err := template.ParseFiles("template.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to load template")
	}
	user := User{
		Id:    1,
		Name:  "Nattapon Sub-Anake",
		Email: "nattapon.subanake@gmail.com",
		Phone: "66815517772",
	}
	temp.Execute(w, user)

	elapsed := float64(time.Since(startTime).Microseconds()) / 1000.0
	fmt.Printf("%s %s %v %.2f ms\n", r.Method, r.URL.RequestURI(), 200, elapsed)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/template", templateHandler)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))

}
