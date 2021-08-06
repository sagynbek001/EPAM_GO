package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Info struct {
	Host       string      `json:"host"`
	UserAgent  string      `json:"user-agent"`
	RequestURI string      `json:"request_uri"`
	Header     http.Header `json:"headers"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// send Not found in such case
	info := Info{
		Host:       r.Host,
		UserAgent:  r.UserAgent(),
		RequestURI: r.RequestURI,
		Header:     r.Header,
	}
	new_json, _ := json.MarshalIndent(info, "", "    ")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
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
