package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ruanzerah/cloppus/internal/repository"
	"github.com/ruanzerah/cloppus/pkg"
)

func updateMessage(w http.ResponseWriter, r *http.Request) {
	db := &repository.Queries{}
	pathId := chi.URLParam(r, "id")
	messageId, err := uuid.Parse(pathId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newMessage repository.Message
	if err := json.NewDecoder(r.Body).Decode(&newMessage); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	message, err := db.UpdateMessage(r.Context(), repository.UpdateMessageParams{
		ID:      messageId,
		Owner:   newMessage.Owner,
		Subject: newMessage.Subject,
		Content: newMessage.Content,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := pkg.WriteJSON(w, http.StatusOK, message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
