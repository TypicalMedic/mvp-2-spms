package handlers

import (
	"encoding/json"
	"mvp-2-spms/internal"
	"net/http"
)

type GoogleCalendarHandler struct {
	planners internal.Planners
}

func InitGoogleCalendarHandler(planners internal.Planners) GoogleCalendarHandler {
	return GoogleCalendarHandler{
		planners: planners,
	}
}

func (h *GoogleCalendarHandler) GetLink(w http.ResponseWriter, r *http.Request) {
	redirectURI := r.URL.Query().Get("redirect")
	result := (*h.planners[internal.GoogleCalendar]).GetAuthLink(redirectURI)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *GoogleCalendarHandler) Auth(w http.ResponseWriter, r *http.Request) {
	cred := GetCredentials(r)
	(*h.planners[internal.GoogleCalendar]).Authentificate(cred.GoogleCalendarToken)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
