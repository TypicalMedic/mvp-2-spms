package handlers

import (
	"encoding/json"
	"mvp-2-spms/internal"
	"net/http"
)

type GitHubHandler struct {
	repos internal.GitRepositoryHubs
}

func InitGitHubHandler(repos internal.GitRepositoryHubs) GitHubHandler {
	return GitHubHandler{
		repos: repos,
	}
}

func (h *GitHubHandler) GetLink(w http.ResponseWriter, r *http.Request) {
	redirectURI := r.URL.Query().Get("redirect")
	result := (*h.repos[internal.GitHub]).GetAuthLink(redirectURI)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *GitHubHandler) Auth(w http.ResponseWriter, r *http.Request) {
	cred := GetCredentials(r)
	(*h.repos[internal.GitHub]).Authentificate(cred.GoogleDriveToken)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
