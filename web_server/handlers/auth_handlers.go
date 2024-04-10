package handlers

import (
	"encoding/json"
	"mvp-2-spms/services/manage-accounts/inputdata"
	"mvp-2-spms/web_server/handlers/interfaces"
	requestbodies "mvp-2-spms/web_server/handlers/request-bodies"
	"mvp-2-spms/web_server/session"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type AuthHandler struct {
	accountInteractor interfaces.IAccountInteractor
}

func InitAuthHandler(acc interfaces.IAccountInteractor) AuthHandler {
	return AuthHandler{
		accountInteractor: acc,
	}
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	// проверяем соответсвтвие типа содержимого запроса
	if headerContentTtype != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	// декодируем тело запроса
	var creds requestbodies.Credentials
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	input := inputdata.CheckCredsValidity{
		Login:    creds.Username,
		Password: creds.Password,
	}
	valid := h.accountInteractor.CheckCredsValidity(input)

	if !valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create a new random session token
	sessionToken := uuid.NewString() + "/" + creds.Username
	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	session.Sessions[sessionToken] = session.InitSession(creds.Username, expiresAt)

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
		Path:    "/",
	})

	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	// проверяем соответсвтвие типа содержимого запроса
	if headerContentTtype != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	// декодируем тело запроса
	var creds requestbodies.Credentials
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	input := inputdata.CheckUsernameExists{
		Login: creds.Username,
	}
	usernameExists := h.accountInteractor.CheckUsernameExists(input)
	if usernameExists {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode("username already exists")
		return
	}

	signupInput := inputdata.SignUp{
		Login:    creds.Username,
		Password: creds.Password,
	}

	h.accountInteractor.SignUp(signupInput)

	// Create a new random session token
	sessionToken := uuid.NewString() + "/" + creds.Username
	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	session.Sessions[sessionToken] = session.InitSession(creds.Username, expiresAt)

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
		Path:    "/",
	})

	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	// remove the users session from the session map
	delete(session.Sessions, sessionToken)

	// We need to let the client know that the cookie is expired
	// In the response, we set the session token to an empty
	// value and set its expiry as the current time
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) RefreshSession(w http.ResponseWriter, r *http.Request) {
	// session is already validated by middleware
	c, _ := r.Cookie("session_token")
	sessionToken := c.Value
	userSession := session.Sessions[sessionToken]

	newSessionToken := uuid.NewString() + "/" + userSession.GetUsername()
	expiresAt := time.Now().Add(120 * time.Second)

	// Set the token in the session map, along with the user whom it represents
	session.Sessions[newSessionToken] = session.InitSession(
		userSession.GetUsername(), expiresAt,
	)

	// Delete the older session token
	delete(session.Sessions, sessionToken)

	// Set the new token as the users `session_token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: time.Now().Add(120 * time.Second),
		Path:    "/",
	})
	w.WriteHeader(http.StatusOK)
}
