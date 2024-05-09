package handlers

import (
	"mvp-2-spms/web_server/session"
	"net/http"
	"strconv"
)

func Authentificator(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			creds := GetCredentials(r)
			userSession, exists := session.Sessions[creds.Session]
			if !exists {
				// If the session token is not present in session map, return an unauthorized error
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if userSession.IsExpired() {
				delete(session.Sessions, creds.Session)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		},
	)
}

func GetSessionUser(r *http.Request) session.UserInfo {
	// session is already validated by middleware
	creds := GetCredentials(r)
	userSession := session.Sessions[creds.Session]
	return userSession.GetUser()
}

// ///////////////////////////////////////////////////////////////////////////??
func GetCredentials(r *http.Request) Credntials {
	professorId, _ := strconv.ParseUint(r.Header.Get("Professor-Id"), 10, 32)
	session := r.Header.Get("Session-Id")
	gcTok := r.Header.Get("Google-Calendar-Token")
	gdTok := r.Header.Get("Google-Drive-Token")
	ghTok := r.Header.Get("GitHub-Token")
	return Credntials{
		ProfessorId:         uint(professorId),
		Session:             session,
		GoogleCalendarToken: gcTok,
		GoogleDriveToken:    gdTok,
		GitHubToken:         ghTok,
	}
}

type Credntials struct {
	ProfessorId         uint
	Session             string
	GoogleCalendarToken string
	GoogleDriveToken    string
	GitHubToken         string
}
