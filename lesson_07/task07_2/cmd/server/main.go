package main

import (
	"log"
	"net/http"

	"github.com/wshaman/course-jsonrpc/handlers"
)

func main() {
	addr := "localhost:8081"
	endpoint1 := "/rpc/register"
	endpoint2 := "/rpc/list"
	http.HandleFunc(endpoint1, handlers.Handle)
	http.HandleFunc(endpoint2, handlers.Handle)
	log.Printf("Started JSON RPC server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
