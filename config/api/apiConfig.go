package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	msg "github.com/ruanzerah/cloppus/internal/messages/routes"
	"github.com/ruanzerah/cloppus/internal/oauth"
	usr "github.com/ruanzerah/cloppus/internal/user/routes"
)

func InitAPI() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	oauth.RegisterAuthRoutes(r)
	msg.RegisterMessageRoutes(r)
	usr.RegisterUserRoutes(r)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `<html><body><a href="/auth/login">Log in with GitHub</a></body></html>`)
	})

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Printf("Error starting server: %v", err)
		return err
	}
	log.Println("Server started successfully")
	return nil
}
