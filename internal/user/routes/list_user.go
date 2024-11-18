package routes

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/ruanzerah/cloppus/internal/repository"
)

func listUser(w http.ResponseWriter, r *http.Request) {
	db := &repository.Queries{}
	pathID := r.PathValue("id")

	userID, err := uuid.Parse(pathID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	user, err := db.ListUser(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
