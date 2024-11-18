package routes

import "net/http"

func HandleMessageRoutes() {
	http.HandleFunc("POST /user/{id}/message", createMessage)
	http.HandleFunc("PUT /user/{username}/message/{id}", updateMessage)
	http.HandleFunc("GET /user/message/{id}", listMessage)
	http.HandleFunc("DELETE /user/message/{id}", deleteMessage)
	http.HandleFunc("PUT /user/message/like/{id}", likeMessage)
}
