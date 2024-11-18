package routes

import "net/http"

func HandleUserRoutes() {
	http.HandleFunc("POST /user", createUser)
	http.HandleFunc("GET /user/{id}", listUser)
	http.HandleFunc("PUT /user/{id}/username", updateUser)
	http.HandleFunc("PUT /user/{id}/password", updateUser)
	http.HandleFunc("DELETE /user/{id}", deleteUser)
}
