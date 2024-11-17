package api

import (
	"fmt"
	"net/http"

	"github.com/ruanzerah/cloppus/internal/messages"
	"github.com/ruanzerah/cloppus/internal/user"
)

func InitAPI(addr string) error {
	mux := http.NewServeMux()

	user.HandleUserRoutes()
	messages.HandleMessageRoutes()

	if addr == "" {
		return fmt.Errorf("API address is not configured")
	}

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		return err
	}
	return nil
}
