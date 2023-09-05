package server

import (
	"context"
	"enterprise-front/internal/server/handler"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// New creates new instance of http.Server
//
// Pre-cond: given context
//
// Post-cond: new instance of http.Server was returned
func New(ctx context.Context) *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("./public/"))
	r.Handle("/public/*", http.StripPrefix("/public/", fs))

	r.Post("/api/enterprises/register", handler.EnterpriseRegisterAPI)
	r.Post("/api/enterprises/{title}/register/manager", handler.RegisterManagerAPI)

	r.Get("/enterprises/register", handler.EnterprisesRegisterPage)
	r.Get("/enterprises/{title}", handler.EnterpriseByTitlePage)
	r.Get("/enterprises", handler.EnterprisesPage)
	r.Get("/{title}/managers/register", handler.ManagersRegister)
	r.Get("/managers", handler.Managers)
	r.Get("/login", handler.Login)
	r.Get("/", handler.Index)
	return &http.Server{
		Addr:        ":8080",
		Handler:     r,
		BaseContext: func(l net.Listener) context.Context { return ctx },
	}
}
