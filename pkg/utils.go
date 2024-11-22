package pkg

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ruanzerah/cloppus/internal/repository"
)

type Response struct {
	Message string `json:"message"`
}

func ValidateMessage(messageBody *repository.Message) error {
	if messageBody.Subject == "" || messageBody.Content == "" {
		return errors.New("missing subject or content")
	}
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
