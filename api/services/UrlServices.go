package services

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/DevKayoS/goFirstAPI/utils"
	"github.com/go-chi/chi/v5"
)

type PostBody struct {
	URL string `json:"url"`
}

func HandlePost(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body PostBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			utils.SendJson(w, utils.Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
		}

		if _, err := url.Parse(body.URL); err != nil {
			utils.SendJson(w, utils.Response{Error: "invalid url passed"}, http.StatusBadRequest)
		}

		code := utils.GenCode()
		db[code] = body.URL
		utils.SendJson(w, utils.Response{Data: code}, http.StatusCreated)
	}
}

func HandleGet(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")

		url, ok := db[code]

		if !ok {
			utils.SendJson(w, utils.Response{Error: "something went wrong"}, http.StatusInternalServerError)
		}

		utils.SendJson(w, utils.Response{Data: url}, http.StatusOK)
	}
}
