package server

import (
	"auth-service/server/api"
	amiddle "auth-service/server/middleware"
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

func (s *server) ShutDown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}

func New(ctx context.Context) *server {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.HandleFunc("/login/admin", api.LoginAdmin)

	r.Group(func(r chi.Router) {
		r.Use(amiddle.AuthJWT)
		r.Get("/logout/admin", api.LogoutAdmin)
	})
	return &server{
		http: &http.Server{
			Addr:        ":8085",
			Handler:     r,
			BaseContext: func(l net.Listener) context.Context { return ctx },
		},
	}
}
