package handlers

import (
	"mvp-2-spms/web_server/session"
	"net/http"
	"strconv"
)

func Authentificator(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// We can obtain the session token from the requests cookies, which come with every request
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

			userSession, exists := session.Sessions[sessionToken]
			if !exists {
				// If the session token is not present in session map, return an unauthorized error
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if userSession.IsExpired() {
				delete(session.Sessions, sessionToken)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		},
	)
}

// ///////////////////////////////////////////////////////////////////////////??
func GetCredentials(r *http.Request) Credntials {
	professorId, _ := strconv.ParseUint(r.Header.Get("Professor-Id"), 10, 32)
	authTok := r.Header.Get("Auth-Token")
	gcTok := r.Header.Get("Google-Calendar-Token")
	gdTok := r.Header.Get("Google-Drive-Token")
	ghTok := r.Header.Get("GitHub-Token")
	return Credntials{
		ProfessorId:         uint(professorId),
		AuthToken:           authTok,
		GoogleCalendarToken: gcTok,
		GoogleDriveToken:    gdTok,
		GitHubToken:         ghTok,
	}
}

type Credntials struct {
	ProfessorId         uint
	AuthToken           string
	GoogleCalendarToken string
	GoogleDriveToken    string
	GitHubToken         string
}
