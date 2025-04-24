package services

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/DevKayoS/goFirstAPI/store"
	"github.com/DevKayoS/goFirstAPI/utils"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

type PostBody struct {
	URL string `json:"url"`
}

type getShortenedUrlResponse struct {
	FullUrl string `json:"full_url"`
}

func HandleCreateShortenUrl(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body PostBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			utils.SendJson(w, utils.Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
		}

		if _, err := url.Parse(body.URL); err != nil {
			utils.SendJson(w, utils.Response{Error: "invalid url passed"}, http.StatusBadRequest)
		}

		code, err := store.SaveShortenedUrl(r.Context(), body.URL)

		if err != nil {
			slog.Error("failed to create code", "error", err)
			utils.SendJson(w, utils.Response{Error: "something went wrong"}, http.StatusBadRequest)
		}

		utils.SendJson(w, utils.Response{Data: code}, http.StatusCreated)
	}
}

func HandleGetShortenedUrl(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")

		fullUrl, err := store.GetFullUrl(r.Context(), code)

		if err != nil {
			if errors.Is(err, redis.Nil) {
				utils.SendJson(w, utils.Response{Error: "code not found"}, http.StatusNotFound)
				return
			}
			slog.Error("failed to get code", "error", err)
			utils.SendJson(w, utils.Response{Error: "something went wrong"}, http.StatusInternalServerError)
			return
		}

		utils.SendJson(w, utils.Response{Data: getShortenedUrlResponse{FullUrl: fullUrl}}, http.StatusOK)
	}
}

// redirect
func HandleGet(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")

		url, ok := db[code]

		if !ok {
			http.Error(w, "url nao encontrada", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, url, http.StatusPermanentRedirect)
	}
}
