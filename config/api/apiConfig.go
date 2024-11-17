package api

import (
	"fmt"
	"net/http"

	msg "github.com/ruanzerah/cloppus/internal/messages/routes"
	usr "github.com/ruanzerah/cloppus/internal/user/routes"
)

func InitAPI(addr string) error {
	mux := http.NewServeMux()

	usr.HandleUserRoutes()
	msg.HandleMessageRoutes()

	if addr == "" {
		return fmt.Errorf("API address is not configured")
	}

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		return err
	}
	return nil
}
