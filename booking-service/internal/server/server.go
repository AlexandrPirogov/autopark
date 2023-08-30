package server

import (
	"booking-service/internal/server/api"
	"context"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type EnterpriseServer interface {
	ListenAndServe() error
	ShutDown(context.Context) error
}

type server struct {
	http *http.Server
}

func (s *server) ListenAndServe() error {
	return s.http.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}

func New(ctx context.Context) *server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/approve", api.ApproveBooking)
	r.Post("/cancel", api.CancelBooking)
	r.Post("/create", api.CreateBooking)
	r.Post("/choose", api.ChooseBooking)

	return &server{
		http: &http.Server{
			Addr:        ":8080",
			Handler:     r,
			BaseContext: func(l net.Listener) context.Context { return ctx },
		},
	}
}
