package main

import (
	"context"
	"enterprise-service/internal/server"
	"log"
)

func main() {
	s := server.New(context.Background())
	log.Println("Running server")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
