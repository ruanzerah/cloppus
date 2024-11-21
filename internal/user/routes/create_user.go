package routes

import (
	"net/http"

	"github.com/ruanzerah/cloppus/internal/oauth"
	"github.com/ruanzerah/cloppus/internal/repository"
	"github.com/ruanzerah/cloppus/pkg"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	db := &repository.Queries{}

	githubData := r.Context().Value("user").(oauth.GitHubUser)

	err := db.CreateUser(r.Context(), repository.CreateUserParams{
		Username: githubData.Login,
		Email:    githubData.Email,
	})
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusBadRequest)
		return
	}
	res := pkg.DefaultResponse()
	if err := pkg.WriteJSON(w, http.StatusOK, res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
