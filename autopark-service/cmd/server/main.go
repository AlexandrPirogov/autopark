package main

import (
	"autopark-service/internal/server"
	"context"
	"log"
)

func main() {
	s := server.New(context.Background())
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
