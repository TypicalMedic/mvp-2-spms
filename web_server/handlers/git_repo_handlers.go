package handlers

import (
	"encoding/base64"
	"mvp-2-spms/internal"
	"mvp-2-spms/services/manage-accounts/inputdata"
	"mvp-2-spms/web_server/handlers/interfaces"
	"net/http"
	"strconv"
	"strings"
)

type GitRepoHandler struct {
	repos             internal.GitRepositoryHubs
	accountInteractor interfaces.IAccountInteractor
}

func InitGitRepoHandler(repos internal.GitRepositoryHubs, acc interfaces.IAccountInteractor) GitRepoHandler {
	return GitRepoHandler{
		repos:             repos,
		accountInteractor: acc,
	}
}

func (h *GitRepoHandler) GetGitHubLink(w http.ResponseWriter, r *http.Request) {
	returnURL := r.URL.Query().Get("redirect")
	accountId, _ := strconv.Atoi(r.URL.Query().Get("account"))
	redirectURI := "http://127.0.0.1:8080/auth/integration/access/github"
	result := h.repos[internal.GitHub].GetAuthLink(redirectURI, accountId, returnURL)
	http.Redirect(w, r, result, http.StatusTemporaryRedirect)
}

func (h *GitRepoHandler) OAuthCallbackGitHub(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	decodedState, _ := base64.URLEncoding.DecodeString(state)

	// needs further update
	params := strings.Split(string(decodedState), ",")
	accountId, _ := strconv.Atoi(params[0])
	redirect := params[1]

	input := inputdata.SetRepoHubIntegration{
		AccountId: uint(accountId),
		AuthCode:  code,
		Type:      int(internal.GoogleCalendar),
	}

	result := h.accountInteractor.SetRepoHubIntegration(input, h.repos[internal.GitHub])
	w.Header().Add("Google-Calendar-Token", result.AccessToken)
	w.Header().Add("Google-Calendar-Token-Exp", result.Expiry.String())
	http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
}
