package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"mvp-2-spms/internal"
	ainputdata "mvp-2-spms/services/manage-accounts/inputdata"
	"mvp-2-spms/services/models"
	"mvp-2-spms/web_server/handlers/interfaces"
	requestbodies "mvp-2-spms/web_server/handlers/request-bodies"
	"net/http"
	"os"
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

func (h *PlannerIntegrationHandler) GetProfessorPlanners(w http.ResponseWriter, r *http.Request) {
	user, err := GetSessionUser(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	id, err := strconv.Atoi(user.GetProfId())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	integInput := ainputdata.GetPlannerIntegration{
		AccountId: uint(id),
	}

	calendarInfo, err := h.accountInteractor.GetPlannerIntegration(integInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	result, err := h.accountInteractor.GetProfessorIntegrPlanners(fmt.Sprint(id), h.planners[models.PlannerName(calendarInfo.Type)])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *PlannerIntegrationHandler) SetProfessorPlanner(w http.ResponseWriter, r *http.Request) {
	user, err := GetSessionUser(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	id, err := strconv.Atoi(user.GetProfId())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	headerContentTtype := r.Header.Get("Content-Type")
	// проверяем соответсвтвие типа содержимого запроса
	if headerContentTtype != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	// декодируем тело запроса
	var reqB requestbodies.SetProfessorPlanner
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&reqB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.accountInteractor.SetProfessorPlanner(reqB.Id, fmt.Sprint(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *PlannerIntegrationHandler) GetGoogleCalendarLink(w http.ResponseWriter, r *http.Request) {
	user, err := GetSessionUser(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	id, err := strconv.Atoi(user.GetProfId())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	returnURL := r.URL.Query().Get("redirect")
	redirectURI := os.Getenv("RETURN_URL") + "/api/v1/auth/integration/access/googlecalendar"

	result, err := (h.planners[models.GoogleCalendar]).GetAuthLink(redirectURI, int(uint(id)), returnURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func (h *PlannerIntegrationHandler) OAuthCallbackGoogleCalendar(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	decodedState, err := base64.URLEncoding.DecodeString(state)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// needs further update
	params := strings.Split(string(decodedState), ",")

	accountId, err := strconv.Atoi(params[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	redirect := params[1]

	input := ainputdata.SetPlannerIntegration{
		AccountId: uint(accountId),
		AuthCode:  code,
		Type:      int(models.GoogleCalendar),
	}

	result, err := h.accountInteractor.SetPlannerIntegration(input, h.planners[models.GoogleCalendar])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Add("Google-Calendar-Token", result.AccessToken)
	w.Header().Add("Google-Calendar-Token-Exp", result.Expiry.String())
	http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
}
