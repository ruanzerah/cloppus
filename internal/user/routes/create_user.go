package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ruanzerah/cloppus/internal/oauth"
	"github.com/ruanzerah/cloppus/internal/repository"
	"github.com/ruanzerah/cloppus/pkg"
	"golang.org/x/oauth2"
)

func createUser(w http.ResponseWriter, r *http.Request, queries *repository.Queries) {
	cookie, err := r.Cookie("oauth_state")
	if err != nil {
		http.Error(w, "State validation failed", http.StatusUnauthorized)
		log.Println("Failed to retrieve state cookie:", err)
		return
	}
	token := &oauth2.Token{AccessToken: cookie.Value}
	client := oauth.OauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		http.Error(w, "Failed to fetch user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var user oauth.GitHubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		http.Error(w, "Failed to parse user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = queries.CreateUser(r.Context(), repository.CreateUserParams{
		Username:  user.Login,
		Email:     user.Email,
		CreatedAt: time.Now(),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := pkg.DefaultResponse()
	if err := pkg.WriteJSON(w, http.StatusOK, res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
