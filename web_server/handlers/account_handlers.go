package handlers

import (
	"encoding/json"
	"mvp-2-spms/services/manage-accounts/inputdata"
	"mvp-2-spms/web_server/handlers/interfaces"
	"net/http"
)

type AccountHandler struct {
	accountInteractor interfaces.IAccountInteractor
}

func InitAccountHandler(accountInteractor interfaces.IAccountInteractor) AccountHandler {
	return AccountHandler{
		accountInteractor: accountInteractor,
	}
}

func (h *AccountHandler) GetAccountIntegrations(w http.ResponseWriter, r *http.Request) {
	cred := GetCredentials(r)
	input := inputdata.GetAccountIntegrations{
		AccountId: cred.ProfessorId,
	}
	result := h.accountInteractor.GetAccountIntegrations(input)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *AccountHandler) GetAccountInfo(w http.ResponseWriter, r *http.Request) {
	cred := GetCredentials(r)
	input := inputdata.GetAccountInfo{
		AccountId: cred.ProfessorId,
	}
	result := h.accountInteractor.GetAccountInfo(input)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
