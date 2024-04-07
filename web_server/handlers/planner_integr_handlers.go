package handlers

import (
	"encoding/base64"
	"mvp-2-spms/internal"
	"mvp-2-spms/services/manage-accounts/inputdata"
	"mvp-2-spms/services/models"
	"mvp-2-spms/web_server/handlers/interfaces"
	"net/http"
	"strconv"
	"strings"
)

type PlannerIntegrationHandler struct {
	planners          internal.Planners
	accountInteractor interfaces.IAccountInteractor
}

func InitPlannerIntegrationHandler(planners internal.Planners, acc interfaces.IAccountInteractor) PlannerIntegrationHandler {
	return PlannerIntegrationHandler{
		planners:          planners,
		accountInteractor: acc,
	}
}

func (h *PlannerIntegrationHandler) GetGoogleCalendarLink(w http.ResponseWriter, r *http.Request) {
	cred := GetCredentials(r)
	returnURL := r.URL.Query().Get("redirect")
	redirectURI := "http://127.0.0.1:8080/auth/integration/access/googlecalendar"
	result := (h.planners[models.GoogleCalendar]).GetAuthLink(redirectURI, int(cred.ProfessorId), returnURL)
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func (h *PlannerIntegrationHandler) OAuthCallbackGoogleCalendar(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	decodedState, _ := base64.URLEncoding.DecodeString(state)

	// needs further update
	params := strings.Split(string(decodedState), ",")
	accountId, _ := strconv.Atoi(params[0])
	redirect := params[1]

	input := inputdata.SetPlannerIntegration{
		AccountId: uint(accountId),
		AuthCode:  code,
		Type:      int(models.GoogleCalendar),
	}
	result := h.accountInteractor.SetPlannerIntegration(input, h.planners[models.GoogleCalendar])
	w.Header().Add("Google-Calendar-Token", result.AccessToken)
	w.Header().Add("Google-Calendar-Token-Exp", result.Expiry.String())
	http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
}
