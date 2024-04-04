package handlers

import (
	"encoding/json"
	"mvp-2-spms/internal"
	"net/http"

	"golang.org/x/oauth2"
)

type GitRepoHandler struct {
	repos internal.GitRepositoryHubs
}

func InitGitRepoHandler(repos internal.GitRepositoryHubs) GitRepoHandler {
	return GitRepoHandler{
		repos: repos,
	}
}

func (h *GitRepoHandler) GetGitHubLink(w http.ResponseWriter, r *http.Request) {
	redirectURI := r.URL.Query().Get("redirect")
	result := h.repos[internal.GitHub].GetAuthLink(redirectURI, 0, "")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *GitRepoHandler) AuthGitHub(w http.ResponseWriter, r *http.Request) {
	// cred := GetCredentials(r)
	h.repos[internal.GitHub].Authentificate(oauth2.Token{})
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
