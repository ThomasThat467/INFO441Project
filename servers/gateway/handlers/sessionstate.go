package handlers

import (
	"time"

	"github.com/ThomasThat467/INFO441Project/tree/main/servers/models/users"
)

//TODO: define a session state struct for this web server
//see the assignment description for the fields you should include
//remember that other packages can only see exported fields!

/* type SessionUserTime struct {
	// references from https://drstearns.github.io/tutorials/gojson/#secfieldsmustbeexportedtobeencoded
	SessionTime time.Time  `json:"time,omitempty"`
	SessionUser users.User `json:"user,omitempty"`
} */

type SessionState struct {
	// references from https://drstearns.github.io/tutorials/gojson/#secfieldsmustbeexportedtobeencoded
	SigningKey  string     `json:"signingKey"`
	SessionTime time.Time  `json:"time"`
	SessionUser users.User `json:"user"`
}
