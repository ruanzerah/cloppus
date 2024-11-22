package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ruanzerah/cloppus/internal/oauth"
	"github.com/ruanzerah/cloppus/internal/repository"
)

func RegisterMessageRoutes(r chi.Router, queries *repository.Queries) {
	r.Route("/message", func(r chi.Router) {
		r.With(oauth.AuthMiddleware).Post("/", func(w http.ResponseWriter, r *http.Request) {
			createMessage(w, r, queries)
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			listMessage(w, r, queries)
		})
		r.With(oauth.AuthMiddleware).Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
			updateMessage(w, r, queries)
		})
		r.With(oauth.AuthMiddleware).Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			deleteMessage(w, r, queries)
		})
		r.With(oauth.AuthMiddleware).Put("/like/{id}", func(w http.ResponseWriter, r *http.Request) {
			likeMessage(w, r, queries)
		})
	})
}
