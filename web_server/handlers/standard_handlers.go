package handlers

import (
	"net/http"
	"strconv"
)

func Ping(w http.ResponseWriter, r *http.Request) {

}

// move to adapters later
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
