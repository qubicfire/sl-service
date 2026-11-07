package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/qubicfire/sl-service/internal/delivery/handlers"
	"github.com/qubicfire/sl-service/internal/delivery/handlers/link"
	"github.com/qubicfire/sl-service/internal/delivery/handlers/login"
	"github.com/qubicfire/sl-service/internal/delivery/middleware"
	"github.com/qubicfire/sl-service/internal/repository"
)

func main() {
	r := chi.NewRouter()
	repo, err := repository.NewDBRepository()
	defer repo.Close()

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// USE CORS
	r.Use()

	r.Post("/api/login", login.LoginHandler(repo))
	r.Post("/api/register", handlers.RegisterHandler(repo))
	r.Get("/{alias}", handlers.RedirectHandler)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth())

		r.Get("/api/dashboard", handlers.DashboardHandler())
		r.Post("/api/link/create", link.CreateLinkHandler(repo))
		r.Post("/api/link/delete", link.DeleteLinkHandler(repo))
	})

	slog.Info("Starting server on port: 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Failed to start server at port 8080:", err)
	}
}
