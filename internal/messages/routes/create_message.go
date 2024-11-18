package routes

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/ruanzerah/cloppus/internal/repository"
	"github.com/ruanzerah/cloppus/pkg"
)

func createMessage(w http.ResponseWriter, r *http.Request) {
	db := &repository.Queries{}

	pathID := r.PathValue("id")
	userID, err := uuid.Parse(pathID)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	user, err := db.ListUser(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to find user", http.StatusBadRequest)
		return
	}

	var messageBody repository.Message
	if err := json.NewDecoder(r.Body).Decode(&messageBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if messageBody.Subject == "" || messageBody.Content == "" {
		http.Error(w, "missing subject or content", http.StatusBadRequest)
		return
	}

	err = db.CreateMessage(r.Context(), repository.CreateMessageParams{
		Owner:   user.Username,
		Subject: messageBody.Subject,
		Content: messageBody.Content,
	})
	if err != nil {
		http.Error(w, "Failed to create message", http.StatusInternalServerError)
		return
	}
	res := pkg.DefaultResponse()
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
