package pkg

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/ruanzerah/cloppus/internal/repository"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func DefaultResponse() []byte {
	res := &Response{
		Message: "Operation Successful",
		Status:  http.StatusOK,
	}

	result, err := json.Marshal(&res)
	if err != nil {
		log.Panic(err)
	}
	return result
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
