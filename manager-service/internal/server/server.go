package server

import (
	"context"
	"manager-service/internal/server/api/auth"
	"manager-service/internal/server/api/autopark"
	"manager-service/internal/server/api/enterprise"
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

	r.HandleFunc("/login", auth.Login)
	r.HandleFunc("/logout", auth.Logout)

	r.HandleFunc("/enteprises/show", enterprise.EnterprisesList)

	r.HandleFunc("/autopark/brand/list", autopark.BrandList)
	r.HandleFunc("/autopark/brand/register", autopark.BrandList)

	r.HandleFunc("/autopark/car/list", autopark.CarList)
	r.HandleFunc("/autopark/car/register", autopark.CarRegister)
	r.HandleFunc("/autopark/car/delete", autopark.CarDelete)
	return &server{
		http: &http.Server{
			Addr:        ":8090",
			Handler:     r,
			BaseContext: func(l net.Listener) context.Context { return ctx },
		},
	}
}
