package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Info struct {
	Host       string `json:"Host"`
	UserAgent  string `json:"User-Agent"`
	RequestURI string `json:"RequestURI"`
	Header     Header `json:"Header"`
}

type Header struct {
	Accept    []string `json:"Accept"`
	UserAgent []string `json:"User-Agent"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// send Not found in such case
	info := Info{
		Host:       r.Host,
		UserAgent:  r.UserAgent(),
		RequestURI: r.RequestURI,
		Header: Header{
			Accept:    r.Header["Accept"],
			UserAgent: r.Header["User-Agent"],
		},
	}
	new_json, _ := json.MarshalIndent(info, "", "    ")
	fmt.Fprintf(w, "%s\n", new_json)
}

// server port number
const port = 8080

func main() {
	fmt.Printf("Launching server on port: %d \n\n", port)

	// set handler for route '/'
	http.HandleFunc("/", handler)
	// start server without ending
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
