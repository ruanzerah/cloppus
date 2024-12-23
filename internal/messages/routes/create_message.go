package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ruanzerah/cloppus/internal/repository"
	"github.com/ruanzerah/cloppus/pkg"
)

func createMessage(w http.ResponseWriter, r *http.Request, queries *repository.Queries) {
	pathID := chi.URLParam(r, "id")
	userID, err := uuid.Parse(pathID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	user, err := queries.ListUserById(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to find user", http.StatusBadRequest)
		return
	}

	var messageBody repository.Message
	if err := json.NewDecoder(r.Body).Decode(&messageBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := pkg.ValidateMessage(&messageBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = queries.CreateMessage(r.Context(), repository.CreateMessageParams{
		Owner:     user.Username,
		Subject:   messageBody.Subject,
		Content:   messageBody.Content,
		CreatedAt: time.Now(),
	})
	if err != nil {
		http.Error(w, "Failed to create message", http.StatusInternalServerError)
		return
	}
	if err := pkg.WriteJSON(w, http.StatusOK, messageBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
