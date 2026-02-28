package main

import (
	"log"

	"github.com/murilorrs/paycore/apps/server/internal/server"
)

func main() {
	server := server.NewServer()

	log.Println("Server running on port :8000")

	if err := server.Start(":8000"); err != nil {
		log.Fatal(err)
	}
}
