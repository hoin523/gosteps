package main

import (
	"chat-server/server" // Import the server package
	"log"
)

func main() {
	chatServer := server.NewServer()
	port := "8085" // You can change the port here
	log.Printf("Starting chat server on port %s...", port)
	chatServer.Start(port)
}
