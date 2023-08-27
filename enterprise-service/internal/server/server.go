// package server wrapps an http.Server and collects api
// for enterprise-service
package server

import (
	"context"
	"enterprise-service/internal/server/api"
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

// ListenAndServer just a wrapper around http.ListenAndServe
func (s *server) ListenAndServe() error {
	return s.http.ListenAndServe()
}

// Shutdown just a wrapper around http.Shutdown
func (s *server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}

// New returns new instance of http.Server
//
// Pre-cond: given context
//
// Post-cond: pointer to the new instance of server for returned
func New(ctx context.Context) *server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/register", api.RegisterEnterprise)
	r.Get("/{id}/list", api.ReadEnerprise)
	r.Post("/{id}/register/manager", api.RegisterManager)
	r.Post("/list", api.ReadEnerprises)
	return &server{
		http: &http.Server{
			Addr:        ":8080",
			Handler:     r,
			BaseContext: func(l net.Listener) context.Context { return ctx },
		},
	}
}
