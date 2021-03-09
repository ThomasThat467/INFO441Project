package handlers

import (
	"github.com/ThomasThat467/INFO441Project/tree/main/servers/models/users"
	"github.com/ThomasThat467/INFO441Project/tree/main/servers/sessions"
)

type HandlerContext struct {
	SignKey      string
	SessionStore sessions.Store
	UserStore    users.Store
}