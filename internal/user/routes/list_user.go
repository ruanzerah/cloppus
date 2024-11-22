package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ruanzerah/cloppus/internal/repository"
	"github.com/ruanzerah/cloppus/pkg"
)

func listUser(w http.ResponseWriter, r *http.Request, queries *repository.Queries) {
	pathID := chi.URLParam(r, "id")

	userID, err := uuid.Parse(pathID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	user, err := queries.ListUserById(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusBadRequest)
		return
	}

	if err := pkg.WriteJSON(w, http.StatusOK, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
