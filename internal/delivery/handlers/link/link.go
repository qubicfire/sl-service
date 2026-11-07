package link

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/qubicfire/sl-service/internal/delivery/middleware"
	"github.com/qubicfire/sl-service/internal/repository"
)

type CreateLinkRequest struct {
	Url string `json:"url"`
}

type DeleteLinkRequest struct {
	Url string `json:"url"`
}

func CreateLinkHandler(repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := int64(r.Context().Value(middleware.UserIDKey).(float64))
		var request CreateLinkRequest

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		_, err := repo.CreateLink(userID, request.Url, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		slog.Info("Created link %s (full: %s)", "", request.Url)
	}
}

func DeleteLinkHandler(repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Context().Value(middleware.UserIDKey).(float64)
		userID := int64(key)
		var request DeleteLinkRequest

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		err := repo.DeleteLinkByID(userID, request.Url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
