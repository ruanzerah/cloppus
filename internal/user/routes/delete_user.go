package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ruanzerah/cloppus/internal/repository"
	"github.com/ruanzerah/cloppus/pkg"
)

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db := &repository.Queries{}
	pathID := chi.URLParam(r, "id")

	userID, err := uuid.Parse(pathID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = db.DeleteUser(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusBadRequest)
		return
	}
	res := pkg.DefaultResponse()
	if err := pkg.WriteJSON(w, http.StatusOK, res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
