package handlers

import (
	"time"

	"github.com/ThomasThat467/INFO441Project/tree/main/servers/models/users"
)

type SessionState struct {
	Time time.Time
	User users.User
}