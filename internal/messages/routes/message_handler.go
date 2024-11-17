package routes

import (
	"net/http"
)

func HandleMessageRoutes() {
	http.HandleFunc("POST /message", createMessage)
	http.HandleFunc("PUT /message/{id}", updateMessage)
	http.HandleFunc("GET /message/{username}", listMessage)
	http.HandleFunc("DELETE /message/{id}", deleteMessage)
	http.HandleFunc("PUT /message/like/{id}", likeMessage)
}
