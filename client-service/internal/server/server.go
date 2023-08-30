// package server holds functionality for creating and starting server
package server

import (
	"client-service/internal/server/api/autopark"
	"client-service/internal/server/api/booking"
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

// ListenAndServe just a wrapper around http.ListenAndServe
func (s *server) ListenAndServe() error {
	return s.http.ListenAndServe()
}

// ListenAndServe just a wrapper around http.Shutdown
func (s *server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}

// New creates new instance of server
//
// Pre-cond: given context
//
// Post-cond: returned pointer to the new instance server
func New(ctx context.Context) *server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/brand/list", autopark.BrandList)
	r.Post("/car/list", autopark.CarList)

	r.Post("/booking/approve", booking.Approve)
	r.Post("/booking/cancel", booking.Cancel)
	r.Post("/booking/choose", booking.Choose)
	r.Post("/booking/create", booking.Create)
	r.Post("/booking/finish", booking.Finish)

	return &server{
		http: &http.Server{
			Addr:        ":8080",
			Handler:     r,
			BaseContext: func(l net.Listener) context.Context { return ctx },
		},
	}
}
