package oauth

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var (
	clientID     = os.Getenv("ID")
	clientSecret = os.Getenv("SECRET")
)

var OauthConfig = &oauth2.Config{
	ClientID:     clientID,
	ClientSecret: clientSecret,
	Endpoint:     github.Endpoint,
	RedirectURL:  "http://localhost:8080/auth/callback",
	Scopes:       []string{"user"},
}

type GitHubUser struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("github_token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Redirect(w, r, "/auth/login", http.StatusTemporaryRedirect)
				return
			}
			http.Error(w, "Failed to read token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		token := &oauth2.Token{AccessToken: tokenCookie.Value}
		if !token.Valid() {
			tokenSource := OauthConfig.TokenSource(context.Background(), token)
			newToken, err := tokenSource.Token()
			if err != nil {
				http.Error(w, "Token refresh failed: "+err.Error(), http.StatusUnauthorized)
				return
			}
			token = newToken
			http.SetCookie(w, &http.Cookie{
				Name:     "github_token",
				Value:    newToken.AccessToken,
				Path:     "/",
				HttpOnly: true,
				Secure:   true,
			})
		}

		next.ServeHTTP(w, r)
	})
}

func generateStateToken() (string, error) {
	return uuid.New().String(), nil
}

func RegisterAuthRoutes(r chi.Router) {
	r.Get("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		state, err := generateStateToken()
		if err != nil {
			http.Error(w, "Failed to generate state token", http.StatusInternalServerError)
			log.Println("Error generating state token:", err)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "oauth_state",
			Value:    state,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		})
		url := OauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	})

	r.Get("/auth/callback", func(w http.ResponseWriter, r *http.Request) {
		state := r.URL.Query().Get("state")
		if state == "" {
			http.Error(w, "State parameter is missing", http.StatusBadRequest)
			return
		}

		cookie, err := r.Cookie("oauth_state")
		if err != nil {
			http.Error(w, "State validation failed", http.StatusUnauthorized)
			log.Println("Failed to retrieve state cookie:", err)
			return
		}

		if state != cookie.Value {
			http.Error(w, "Invalid state parameter", http.StatusUnauthorized)
			return
		}
		code := r.URL.Query().Get("code")
		token, err := OauthConfig.Exchange(context.Background(), code)
		if err != nil {
			http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "github_token",
			Value:    token.AccessToken,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
		})
		http.Redirect(w, r, "/", http.StatusFound)
	})
}
