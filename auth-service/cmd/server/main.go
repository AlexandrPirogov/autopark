package main

import (
	"auth-service/server"
	"context"
	"log"
)

func main() {
	s := server.New(context.Background())
	log.Println("Running server")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
