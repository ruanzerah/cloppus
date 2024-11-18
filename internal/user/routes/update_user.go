package routes

import (
	"net/http"
	"strings"
)

func updateUser(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if strings.Contains(path, "/usermame") {
		// TODO: Logic
	} else if strings.Contains(path, "/password") {
		// TODO: Logic
	} else {
		// None
	}
}
