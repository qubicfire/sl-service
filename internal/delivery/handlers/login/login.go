package login

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/qubicfire/sl-service/internal/config"
	"github.com/qubicfire/sl-service/internal/repository"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func SendError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(ErrorResponse{
		Message: message,
	})
}

func LoginHandler(repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginRequest LoginRequest

		slog.Info("Got login request")

		if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		user, err := repo.GetUserByEmail(loginRequest.Email)
		if err != nil {
			SendError(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		if user.Password != loginRequest.Password {
			SendError(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		token, err := config.GenerateToken(user.ID)
		if err != nil {
			SendError(w, "Failed to process login", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})
	}
}
