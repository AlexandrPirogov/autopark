package server

import (
	"context"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type EnterpriseServer interface {
	ListenAndServe() error
	ShutDown(context.Context) error
}

type server struct {
	http *http.Server
}

func New(ctx context.Context) *server {
	r := chi.NewRouter()

	return &server{
		http: &http.Server{
			Addr:        ":8080",
			Handler:     r,
			BaseContext: func(l net.Listener) context.Context { return ctx },
		},
	}
}
