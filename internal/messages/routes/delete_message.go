package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ruanzerah/cloppus/internal/repository"
	"github.com/ruanzerah/cloppus/pkg"
)

func deleteMessage(w http.ResponseWriter, r *http.Request, queries *repository.Queries) {
	pathID := chi.URLParam(r, "id")

	messageID, err := uuid.Parse(pathID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = queries.DeleteMessage(r.Context(), messageID)
	if err != nil {
		http.Error(w, "Failed to delete message", http.StatusBadRequest)
		return
	}
	if err := pkg.WriteJSON(w, http.StatusOK, pkg.DefaultResponse()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
