package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ruanzerah/cloppus/internal/repository"
	"github.com/ruanzerah/cloppus/pkg"
)

func deleteUser(w http.ResponseWriter, r *http.Request, queries *repository.Queries) {
	pathID := chi.URLParam(r, "id")

	userID, err := uuid.Parse(pathID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = queries.DeleteUser(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusBadRequest)
		return
	}

	if err := pkg.WriteJSON(w, http.StatusOK, "DELETE USER SUCCESS"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
