package main

import (
	"main/client"
	"main/handler"
	"main/server"
)

func main() {
	server := server.NewServer()
	handlers := handler.NewHandlers()

	go server.Run("8080", handlers.Configure())

	client := client.NewClient()
	client.Run()
}
