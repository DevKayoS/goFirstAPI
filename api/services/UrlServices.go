package services

import "net/http"

type PostBody struct {
	URL string `json:"url"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func HandlePost(w http.ResponseWriter, r *http.Request) {

}

func HandleGet(w http.ResponseWriter, r *http.Request) {

}
