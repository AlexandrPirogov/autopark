package server

import (
	"auth-service/server/api"
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

	r.Post("/login/admin", api.LoginAdmin)
	r.Post("/login/manager", api.LoginManager)

	r.Get("/verify", api.VerifyJWT)
	r.Get("/logout/admin", api.LogoutAdmin)
	r.Get("/logount/manager", api.LogoutManager)
	r.Post("/register/manager", api.RegisterManager)
	return &server{
		http: &http.Server{
			Addr:        ":8080",
			Handler:     r,
			BaseContext: func(l net.Listener) context.Context { return ctx },
		},
	}
}
