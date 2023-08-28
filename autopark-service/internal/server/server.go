package server

import (
	"autopark-service/internal/server/api"
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
	r.HandleFunc("/brand/register", api.RegisterBrand)
	r.HandleFunc("/brand/list", api.ReadBrands)

	r.Post("/car/register", api.CreateCar)
	r.Post("/car/list", api.ReadCars)
	r.Post("/car/remove", api.DeleteCar)
	//r.Get("/car/set/list", api.ReadSetCars)

	r.Put("/car/{uid}/set", api.SetCar)
	r.Put("/car/{uid}/unset", api.UnsetCar)

	return &server{
		http: &http.Server{
			Addr:        ":8081",
			Handler:     r,
			BaseContext: func(l net.Listener) context.Context { return ctx },
		},
	}
}
