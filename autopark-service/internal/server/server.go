package server

import (
	"autopark-service/internal/server/api"
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

func (s *server) ListenAndServe() error {
	return s.http.ListenAndServe()
}

func (s *server) ShutDown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}

func New(ctx context.Context) *server {
	r := chi.NewRouter()

	r.HandleFunc("/brand/register", api.RegisterBrand)
	r.HandleFunc("/brand/list", api.ReadBrands)

	r.HandleFunc("/car/register", api.CreateCar)
	r.HandleFunc("/car/list", api.ReadCars)
	r.HandleFunc("/car/remove", api.DeleteCar)

	return &server{
		http: &http.Server{
			Addr:        ":8081",
			Handler:     r,
			BaseContext: func(l net.Listener) context.Context { return ctx },
		},
	}
}
