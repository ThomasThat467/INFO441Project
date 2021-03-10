package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ThomasThat467/INFO441Project/tree/main/servers/models/users"
	"github.com/ThomasThat467/INFO441Project/tree/main/servers/sessions"

	"golang.org/x/crypto/bcrypt"
)

//TODO: define HTTP handler functions as described in the
//assignment description. Remember to use your handler context
//struct as the receiver on these functions so that you have
//access to things like the session store and user store.

/// UsersHandler handles requests for users resource
func (context *HandlerContext) UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			http.Error(w, "The request body must be a JSON", http.StatusUnsupportedMediaType)
			return
		}
		newUser := &users.NewUser{}
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(newUser); err != nil {
			http.Error(w, "Error converting user to JSON\n"+err.Error(), http.StatusBadRequest)
			return
		}
		user, err := newUser.ToUser()
		if err != nil {
			http.Error(w, "Error creating user\n"+err.Error(), http.StatusBadRequest)
			return
		}
		_, err = context.UserStore.Insert(user)
		if err != nil {
			http.Error(w, "Error inserting user to DB\n"+err.Error(), http.StatusBadRequest)
			return
		}
		newSession := &SessionState{
			Time: time.Now(),
			User: *user,
		}
		_, err = sessions.BeginSession(context.SignKey, context.SessionStore, newSession, w)
		if err != nil {
			http.Error(w, "Error beginning session\n"+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		resBody, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "Error responding with JSON"+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(resBody)
	} else {
		http.Error(w, "Current request method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// SpecificUserHandler handles requests for a specific user
func (context *HandlerContext) SpecificUserHandler(w http.ResponseWriter, r *http.Request) {
	currSess := &SessionState{}
	_, err := sessions.GetState(r, context.SignKey, context.SessionStore, currSess)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if r.Method == "GET" {
		var userID int64
		if path.Base(r.URL.Path) == "me" {
			userID = int64(currSess.User.ID)
		} else {
			lastPath, err := strconv.Atoi(path.Base(r.URL.Path))
			if err != nil {
				http.Error(w, "This user ID does not exist", http.StatusNotFound)
				return
			}
			userID = int64(lastPath)
		}
		user, err := context.UserStore.GetByID(int64(userID))
		if err != nil {
			http.Error(w, "This user ID does not exist", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resBody, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(resBody)
	} else if r.Method == "PATCH" {
		lastPath := path.Base(r.URL.Path)
		if lastPath != "me" {
			userID, err := strconv.Atoi(lastPath)
			if err != nil || int64(userID) != currSess.User.ID {
				http.Error(w, "Cannot request this user ID", http.StatusForbidden)
				return
			}
		}
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			http.Error(w, "The request body must be a JSON", http.StatusUnsupportedMediaType)
			return
		}
		updates := &users.Updates{}
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(updates); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := context.UserStore.Update(currSess.User.ID, updates)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		currSess.User = *user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resBody, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(resBody)
	} else {
		http.Error(w, "Current request method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// SessionsHandler handles requests for sessions resource
func (context *HandlerContext) SessionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			http.Error(w, "The request body must be a JSON", http.StatusUnsupportedMediaType)
			return
		}
		creds := &users.Credentials{}
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(creds); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := context.UserStore.GetByEmail(creds.Email)
		if err != nil {
			bcrypt.GenerateFromPassword([]byte("abcde1234567"), 13)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if err := user.Authenticate(creds.Password); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ip := r.Header.Get("X-FORWARDED-FOR")
		if ip != "" {
			ip = strings.Split(ip, ", ")[0]
		} else {
			ip = r.RemoteAddr
		}
		userLog := &users.UserLog{
			ID:        user.ID,
			StartAt:   time.Now(),
			IPAddress: ip,
		}
		_, err = context.UserStore.InsertUserLog(userLog)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sessionState := &SessionState{
			Time: time.Now(),
			User: *user,
		}
		_, err = sessions.BeginSession(context.SignKey, context.SessionStore, sessionState, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		resBody, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(resBody)
	} else {
		http.Error(w, "Current request method is not allowed", http.StatusMethodNotAllowed)
		return
	}
}

//SpecificSessionHandler handles requests related to a specific session
func (context *HandlerContext) SpecificSessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		if path.Base(r.URL.Path) != "mine" {
			http.Error(w, "Cannot end a session that isn't yours", http.StatusForbidden)
			return
		}
		_, err := sessions.EndSession(r, context.SignKey, context.SessionStore)
		if err != nil {
			http.Error(w, "Error ending this session", http.StatusForbidden)
			return
		}
		w.Write([]byte("signed out"))
	} else {
		http.Error(w, "Current request method is not allowed", http.StatusMethodNotAllowed)
		return
	}
}