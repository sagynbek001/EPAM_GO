package main

import (
	"main/handler"
	"main/server"
)

func main() {
	server := server.NewServer()
	handlers := handler.NewHandlers()

	server.Run("8080", handlers.Configure())
}
