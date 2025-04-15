package api

import (
	"net/http"

	"github.com/DevKayoS/goFirstAPI/api/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler() http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/api/shorten", services.HandlePost)
	r.Get("/{code}", services.HandleGet)

	return r
}
