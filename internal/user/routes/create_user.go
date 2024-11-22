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
	cookie, err := r.Cookie("github_token")
	if err != nil {
		http.Error(w, "Github validation failed", http.StatusUnauthorized)
		log.Println("Failed to retrieve state cookie:", err)
		return
	}
	token := &oauth2.Token{AccessToken: cookie.Value}
	client := oauth.OauthConfig.Client(context.Background(), token)
	userResp, err := client.Get("https://api.github.com/user")
	if err != nil {
		http.Error(w, "Failed to fetch user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer userResp.Body.Close()

	var user oauth.GitHubUser
	if err := json.NewDecoder(userResp.Body).Decode(&user); err != nil {
		http.Error(w, "Failed to parse user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := queries.ListUserByName(r.Context(), user.Login); err == nil {
		http.Error(w, "User already registered.", http.StatusBadRequest)
		return
	}
	var email []oauth.GithubEmail
	emailResp, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		http.Error(w, "Failed to fetch user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer emailResp.Body.Close()
	if err := json.NewDecoder(emailResp.Body).Decode(&email); err != nil {
		http.Error(w, "Failed to parse user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	primaryEmail := ""

	for _, v := range email {
		if v.Primary {
			primaryEmail = v.Email
		} else {
			continue
		}
	}
	if primaryEmail == "" {
		http.Error(w, "Failed to get primary email", http.StatusInternalServerError)
		return
	}
	err = queries.CreateUser(r.Context(), repository.CreateUserParams{
		Username:  user.Login,
		Email:     primaryEmail,
		CreatedAt: time.Now(),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := pkg.WriteJSON(w, http.StatusOK, "CREATE USER SUCCESS"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
