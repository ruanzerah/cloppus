package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ruanzerah/cloppus/internal/repository"
	"github.com/ruanzerah/cloppus/pkg"
)

func listMessage(w http.ResponseWriter, r *http.Request) {
	db := &repository.Queries{}

	pathID := chi.URLParam(r, "id")
	userID, err := uuid.Parse(pathID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	messages, err := db.ListMessages(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := pkg.WriteJSON(w, http.StatusOK, messages); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
