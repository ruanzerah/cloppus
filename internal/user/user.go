package user

import "github.com/ruanzerah/cloppus/internal/messages"

type User struct {
	Username string
	Email    string
	Auth     bool
	Password string
	Messages []messages.Message
}
