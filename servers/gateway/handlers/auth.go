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

// UsersHandler ...
func (ctx *HandlerContext) UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		contentType := r.Header.Get("Content-Type")
		//println(contentType)
		if contentType != "application/json" {
			http.Error(w, "Request body must be json but got: %d", http.StatusUnsupportedMediaType)
			return
		} else {
			//println("in else statement")
			responseBody, _ := ioutil.ReadAll(r.Body)
			newUser := users.NewUser{}
			//json.Unmarshal([]byte(responseBody), &newUser)
			err := json.Unmarshal([]byte(responseBody), &newUser)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			createdNewUser, err := newUser.ToUser()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			newInserted, err := ctx.UserStore.Insert(createdNewUser)
			if err != nil {
				//println("bad request")
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			sessionState := SessionState{SigningKey: ctx.SigningKey, SessionTime: time.Now(), SessionUser: *newInserted}
			sessionID, err := sessions.BeginSession(ctx.SigningKey, ctx.SessionStore, sessionState, w)
			if err != nil {
				fmt.Println("Failed to create session")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if err != nil {
				//println("bad request")
				// nil pointer
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Authorization", "Bearer "+sessionID.String())
			addedUser, _ := json.Marshal(newInserted)
			w.Write(addedUser)
			return
		}
	} else if r.Method == http.MethodOptions {
		//testing to get around preflight cors
		return
	} else {
		//println("big else statement")
		http.Error(w, "Method not allowed %d", http.StatusMethodNotAllowed)
		return
	}
}

// SpecificUserHandler ...
func (ctx *HandlerContext) SpecificUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := strings.TrimPrefix(r.URL.Path, "/v1/users/")
	currentUser := &users.User{}
	sessionID, err := sessions.GetSessionID(r, ctx.SigningKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	w.Header().Set("Authorization", "Bearer "+sessionID.String())

	_, err = sessions.GetState(r, ctx.SigningKey, ctx.SessionStore, &SessionState{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if r.Method == http.MethodGet || r.Method == http.MethodPatch {
		if userID == "me" {
			sessionState := &SessionState{}
			_, err := sessions.GetState(r, ctx.SigningKey, ctx.SessionStore, sessionState)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			currentUser = &sessionState.SessionUser
			userID = fmt.Sprint(currentUser.ID)
		} else {
			numID, _ := strconv.ParseInt(userID, 10, 64)
			currentUser, err = ctx.UserStore.GetByID(numID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
		}
		userID = strconv.FormatInt(currentUser.ID, 10)
		intID := currentUser.ID
		if r.Method == http.MethodGet {
			user, err := ctx.UserStore.GetByID(intID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			json, _ := json.Marshal(user)
			w.Header().Set("Content-Type", "application/json")
			w.Write(json)
			w.WriteHeader(http.StatusOK)
			return
		}
	} else if r.Method == http.MethodPatch {
		if strconv.FormatInt(currentUser.ID, 10) != userID {
			fmt.Printf("Status not found for that id. Code: %d", http.StatusNotFound)
			return
		}
		if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
			fmt.Printf("Unnaccepted content type. Response body must be in JSON. Code: %d", http.StatusUnsupportedMediaType)
			return
		}
		marshaled, err := ioutil.ReadAll(r.Body)
		var updates users.Updates
		if err == nil {
			json.Unmarshal([]byte(marshaled), &updates)
		}
		intID := currentUser.ID
		updated, err := ctx.UserStore.Update(intID, &updates)
		w.Header().Set("Content-Type", "application/json")
		marshalUser, err := json.Marshal(updated)
		if err == nil {
			w.Write(marshalUser)
		}
		w.WriteHeader(http.StatusOK)
		return

	} else if r.Method == http.MethodOptions {
		//testing to get around preflight cors
		return
	}
}

// SessionsHandler ...
func (ctx *HandlerContext) SessionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		contentType := r.Header.Get("Content-Type")
		//println(r.Method)
		if contentType != "application/json" {
			//println("ctype" + contentType)
			fmt.Printf("Request body must be json. Got code: %d \n", http.StatusUnsupportedMediaType)
			return
		}

		credentials := users.Credentials{}
		marshaled, err := ioutil.ReadAll(r.Body)

		if err == nil {
			json.Unmarshal([]byte(marshaled), &credentials)
			user, err := ctx.UserStore.GetByEmail(credentials.Email)
			if user == nil {
				bcrypt.GenerateFromPassword([]byte("abcde1234567"), 13)
				http.Error(w, "Invalid creds", http.StatusUnauthorized)
				return
			}
			if err != nil || user.Authenticate(credentials.Password) != nil {
				http.Error(w, "Invalid creds", http.StatusUnauthorized)
				return
			} else {
				sessionState := &SessionState{
					SessionTime: time.Now(),
					SessionUser: *user,
				}
				sessions.BeginSession(ctx.SigningKey, ctx.SessionStore, sessionState, w)
				headerIP := r.Header.Get("X-Forwarded-For")
				currentIP := r.RemoteAddr
				if len(headerIP) != 0 {
					currentIP = headerIP
				}

				signInUser := users.SignIn{
					ID:         int64(0),
					UserID:     user.ID,
					SignInTime: time.Now().String(),
					IP:         currentIP,
				}
				ctx.UserStore.InsertSignedIn(&signInUser)
				//w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				w.Write(marshaled)
			}
		}
	} else if r.Method == http.MethodOptions {
		//testing to get around preflight cors
		return
	} else {
		//println("in big else")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Printf("SessionsHandler - Status method not allowed. Code: %d \n", http.StatusMethodNotAllowed)
		return
	}
}

// SpecificSessionHandler ...
func (ctx *HandlerContext) SpecificSessionHandler(w http.ResponseWriter, r *http.Request) {
	mine := strings.TrimPrefix(r.URL.Path, "/v1/sessions/")
	if r.Method == http.MethodDelete {
		if mine != "mine" {
			fmt.Printf("Error status forbidden. Code: %d \n", http.StatusForbidden)
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Access forbidden"))
		} else {
			_, err := sessions.EndSession(r, ctx.SigningKey, ctx.SessionStore)
			if err == nil {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Signed out"))
			}
		}
	} else if r.Method == http.MethodOptions {
		//testing to get around preflight cors
		return
	} else {
		fmt.Printf("SpecificSessionHandler - Status method not allowed. Code: %d \n", http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
