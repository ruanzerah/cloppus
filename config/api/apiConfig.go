package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"

	msg "github.com/ruanzerah/cloppus/internal/messages/routes"
	"github.com/ruanzerah/cloppus/internal/oauth"
	"github.com/ruanzerah/cloppus/internal/repository"
	usr "github.com/ruanzerah/cloppus/internal/user/routes"
)

func InitAPI(dbConn *pgx.Conn) error {
	queries := repository.New(dbConn)
	defer dbConn.Close(context.Background())
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	oauth.RegisterAuthRoutes(r)
	msg.RegisterMessageRoutes(r, queries)
	usr.RegisterUserRoutes(r, queries)

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
