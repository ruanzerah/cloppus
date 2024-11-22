package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ruanzerah/cloppus/internal/oauth"
	"github.com/ruanzerah/cloppus/internal/repository"
)

func RegisterUserRoutes(r chi.Router, queries *repository.Queries) {
	r.Route("/user", func(r chi.Router) {
		r.Use(oauth.AuthMiddleware)
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			listUser(w, r, queries)
		})
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			createUser(w, r, queries)
		})
		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			deleteUser(w, r, queries)
		})
	})
}
