package main

import (
	"log"

	grpcserver "github.com/acidonper/jump-app-golang-grpc/internal/server"
)

func main() {
	log.Printf("Starting Server Process...")
	grpcserver.Start()
}