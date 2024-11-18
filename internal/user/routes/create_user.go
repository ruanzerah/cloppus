package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ruanzerah/cloppus/internal/repository"
	"github.com/ruanzerah/cloppus/pkg"
	"golang.org/x/crypto/bcrypt"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	db := &repository.Queries{}

	var userBody repository.User

	err := json.NewDecoder(r.Body).Decode(&userBody)
	if err != nil {
		http.Error(w, "Failed to decode request", http.StatusBadRequest)
		return
	}
	if userBody.Username == "" || userBody.Email == "" || userBody.Hash == "" {
		http.Error(w, "missing credentials", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userBody.Hash), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	err = db.CreateUser(r.Context(), repository.CreateUserParams{
		Username: userBody.Username,
		Email:    userBody.Email,
		Auth:     false,
		Hash:     string(hash),
	})
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	res := pkg.DefaultResponse()
	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
