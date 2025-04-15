package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func SendJson(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("erro ao fazer marshal de json", "error", err)
		SendJson(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("erro ao enviar a resposta", "error", err)
		return
	}
}
