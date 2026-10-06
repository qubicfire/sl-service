package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/qubicfire/sl-service/internal/delivery/handlers"
	"github.com/qubicfire/sl-service/internal/delivery/middleware"
)

func startDatabase() {
}

func main() {
	startDatabase()

	jwtSecret := []byte("secret-key")
	r := chi.NewRouter()

	r.Use()

	r.Post("/api/login", handlers.LoginHandler)
	r.Post("/api/register", handlers.RegisterHandler)
	r.Get("/{alias}", handlers.RedirectHandler)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth(jwtSecret))

		r.Get("/api/dashboard", handlers.DashboardHandler)
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Failed to start server at port 8080:", err)
	}
}
