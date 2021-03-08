package handlers

import (
	"github.com/ThomasThat467/INFO441Project/tree/main/servers/models/users"
	"github.com/ThomasThat467/INFO441Project/tree/main/servers/sessions"
)

//TODO: define a handler context struct that
//will be a receiver on any of your HTTP
//handler functions that need access to
//globals, such as the key used for signing
//and verifying SessionIDs, the session store
//and the user store

type HandlerContext struct {
	SigningKey   string
	SessionStore sessions.Store
	UserStore    users.Store
}

func NewHandlerContext(signingKey string, sessionStore sessions.Store, userStore users.Store) *HandlerContext {

	return &HandlerContext{
		SigningKey:   signingKey,
		SessionStore: sessionStore,
		UserStore:    userStore,
	}
}
