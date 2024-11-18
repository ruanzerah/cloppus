package routes

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/ruanzerah/cloppus/internal/repository"
	"github.com/ruanzerah/cloppus/pkg"
)

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db := &repository.Queries{}
	pathID := r.PathValue("id")

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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
