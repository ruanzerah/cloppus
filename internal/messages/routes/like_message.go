package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ruanzerah/cloppus/internal/repository"
	"github.com/ruanzerah/cloppus/pkg"
)

func likeMessage(w http.ResponseWriter, r *http.Request, queries *repository.Queries) {
	pathId := chi.URLParam(r, "id")
	messageId, err := uuid.Parse(pathId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = queries.LikeMessage(r.Context(), messageId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := pkg.WriteJSON(w, http.StatusOK, pkg.DefaultResponse()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
