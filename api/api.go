package api

import (
	"net/http"

	"github.com/DevKayoS/goFirstAPI/api/services"
	"github.com/DevKayoS/goFirstAPI/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(store store.Store) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Route("/url", func(r chi.Router) {
			r.Post("/shorten", services.HandleCreateShortenUrl(store))
			r.Get("/{code}", services.HandleGetShortenedUrl(store))
		})
	})

	return r
}
