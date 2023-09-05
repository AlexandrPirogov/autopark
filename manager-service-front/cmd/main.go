package main

import (
	"context"
	"manager-service-front/internal/server"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	s := server.New(ctx)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal().Msgf("%v", err)
		}
	}()

	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	log.Warn().Msg("Running service-manager-front")
	sig := <-cancelChan
	log.Printf("Got signal %v\n", sig)

	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelShutdown()

	if err := s.Shutdown(ctxShutdown); err != nil {
		log.Printf("Shutdown error %v\n", err)
		defer os.Exit(1)
		return
	}

	log.Printf("Server shutdowned\n")

	cancel()
	defer os.Exit(0)
	close(cancelChan)
}
