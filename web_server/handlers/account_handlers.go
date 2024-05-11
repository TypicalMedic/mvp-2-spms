package handlers

import (
	"encoding/json"
	"fmt"
	"mvp-2-spms/internal"
	"mvp-2-spms/services/manage-accounts/inputdata"
	"mvp-2-spms/services/models"
	"mvp-2-spms/web_server/handlers/interfaces"
	"net/http"
	"strconv"
)

type AccountHandler struct {
	accountInteractor interfaces.IAccountInteractor
	cloudDrives       internal.CloudDrives
}

func InitAccountHandler(accountInteractor interfaces.IAccountInteractor, cd internal.CloudDrives) AccountHandler {
	return AccountHandler{
		accountInteractor: accountInteractor,
		cloudDrives:       cd,
	}
}

func (h *AccountHandler) GetAccountIntegrations(w http.ResponseWriter, r *http.Request) {
	user := GetSessionUser(r)
	id, _ := strconv.Atoi(user.GetProfId())
	input := inputdata.GetAccountIntegrations{
		AccountId: uint(id),
	}
	result := h.accountInteractor.GetAccountIntegrations(input)

	if result.CloudDrive.BaseFolderId != "" {
		result.CloudDrive.BaseFolderName = h.accountInteractor.GetDriveBaseFolderName(
			result.CloudDrive.BaseFolderId, fmt.Sprint(id), h.cloudDrives[models.CloudDriveName(result.CloudDrive.Type.Id)])
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *AccountHandler) GetAccountInfo(w http.ResponseWriter, r *http.Request) {
	user := GetSessionUser(r)
	id, _ := strconv.Atoi(user.GetProfId())
	input := inputdata.GetProfessorInfo{
		AccountId: uint(id),
	}
	result := h.accountInteractor.GetProfessorInfo(input)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
