package api

import (
	"net/http"

	"github.com/DevKayoS/goFirstAPI/api/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(db map[string]string) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/api/shorten", services.HandlePost(db))
	r.Get("/{code}", services.HandleGet(db))

	return r
}
