package main

import (
	"link-shortner/api"
	"log/slog"
	"net/http"
	"time"
)

func main(){
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}
	slog.Info("all systems offline")
}

func run() error {
	db := make(map[string]string)
	handler := api.NewHandler(db)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}
	
	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		slog.Error("failed to start server", "error", err)
		return err
	}
	return nil
}
