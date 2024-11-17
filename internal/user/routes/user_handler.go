package routes

import "net/http"

func HandleUserRoutes() {
	http.HandleFunc("POST /user", createUser)
	http.HandleFunc("PUT /user", updateUser)
	http.HandleFunc("DELETE /user", deleteUser)
}
