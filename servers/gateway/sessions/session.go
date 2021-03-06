package sessions

import (
	"errors"
	"net/http"
	"strings"
)

const headerAuthorization = "Authorization"
const paramAuthorization = "auth"
const schemeBearer = "Bearer "

//ErrNoSessionID is used when no session ID was found in the Authorization header
var ErrNoSessionID = errors.New("no session ID found in " + headerAuthorization + " header")

//ErrInvalidScheme is used when the authorization scheme is not supported
var ErrInvalidScheme = errors.New("authorization scheme not supported")

//BeginSession creates a new SessionID, saves the `sessionState` to the store, adds an
//Authorization header to the response with the SessionID, and returns the new SessionID
func BeginSession(signingKey string, store Store, sessionState interface{}, w http.ResponseWriter) (SessionID, error) {
	// Fail fast error handling
	// create a new SessionID
	newSessionID, sessionErr := NewSessionID(signingKey)
	if sessionErr != nil {
		return InvalidSessionID, sessionErr
	}
	// save the sessionState to the store
	if saveErr := store.Save(newSessionID, sessionState); saveErr != nil {
		return InvalidSessionID, saveErr
	}

	// add a header to the ResponseWriter that looks like this:
	//    "Authorization: Bearer <sessionID>"
	w.Header().Add(headerAuthorization, schemeBearer+string(newSessionID))
	return newSessionID, nil
}

//GetSessionID extracts and validates the SessionID from the request headers
func GetSessionID(r *http.Request, signingKey string) (SessionID, error) {
	sessionID := r.Header.Get(headerAuthorization)
	if sessionID == "" {
		if r.URL.Query().Get("auth") == "" {
			return "", ErrNoSessionID
		}
		sessionID = r.URL.Query().Get("auth")
	}

	if !strings.HasPrefix(sessionID, schemeBearer) {
		return InvalidSessionID, ErrInvalidScheme
	}

	sessionID = strings.Replace(sessionID, schemeBearer, "", 1)
	validatedID, err := ValidateID(sessionID, signingKey)
	if err != nil {
		return InvalidSessionID, err
	}

	return validatedID, nil
}

//GetState extracts the SessionID from the request,
//gets the associated state from the provided store into
//the `sessionState` parameter, and returns the SessionID
func GetState(r *http.Request, signingKey string, store Store, sessionState interface{}) (SessionID, error) {
	sessionID, sessionErr := GetSessionID(r, signingKey)
	if sessionErr != nil {
		return InvalidSessionID, sessionErr
	}

	if getErr := store.Get(sessionID, sessionState); getErr != nil {
		return InvalidSessionID, getErr
	}

	return sessionID, nil
}

//EndSession extracts the SessionID from the request,
//and deletes the associated data in the provided store, returning
//the extracted SessionID.
func EndSession(r *http.Request, signingKey string, store Store) (SessionID, error) {
	sessionID, sessionErr := GetSessionID(r, signingKey)
	if sessionErr != nil {
		return InvalidSessionID, sessionErr
	}

	if deleteErr := store.Delete(sessionID); deleteErr != nil {
		return InvalidSessionID, deleteErr
	}

	return sessionID, nil
}
