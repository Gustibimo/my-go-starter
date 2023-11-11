package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	println("gRPC Server")

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()

	// Start the server
	log.Println("Starting gRPC server on port 8080...")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}