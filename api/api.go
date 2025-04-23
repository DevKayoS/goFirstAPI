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

	r.Route("/api", func(r chi.Router) {
		r.Route("/url", func(r chi.Router) {
			r.Post("/api/shorten", services.HandleCreateShortenUrl(db))
			r.Get("/{code}", services.HandleGet(db))
		})
	})

	return r
}
