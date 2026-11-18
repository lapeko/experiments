package api

import (
	"fmt"
	"net/http"
	"profiling/internal/api/cpu"
	"profiling/internal/api/healthz"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type api struct {
	r chi.Router
}

func NewApi() *api {
	return &api{
		r: chi.NewRouter(),
	}
}

func (a *api) Start() error {
	r := a.r
	r.Use(middleware.Logger)

	r.Route("/cpu", cpu.Router)
	r.Route("/healthz", healthz.Router)

	r.Mount("/debug", middleware.Profiler())

	fmt.Println("Server running on port 3000")
	return http.ListenAndServe(":3000", r)
}
