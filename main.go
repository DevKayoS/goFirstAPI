package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/DevKayoS/goFirstAPI/api"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}

	slog.Info("all system offline")
}

func run() error {
	handler := api.NewHandler()

	app := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := app.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
