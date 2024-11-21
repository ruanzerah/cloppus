package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/ruanzerah/cloppus/internal/oauth"
)

func RegisterUserRoutes(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Get("/{id}", listUser)
		r.With(oauth.AuthMiddleware).Post("/", createUser)
		r.Delete("/{id}", deleteUser)
	})
}
