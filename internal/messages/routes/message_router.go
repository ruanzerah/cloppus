package routes

import (
	"github.com/go-chi/chi/v5"
)

func RegisterMessageRoutes(r chi.Router) {
	r.Route("/message", func(r chi.Router) {
		r.Post("/", createMessage)
		r.Get("/{id}", listMessage)
		r.Put("/{id}", updateMessage)
		r.Delete("/{id}", deleteMessage)
		r.Put("/like/{id}", likeMessage)
	})
}
