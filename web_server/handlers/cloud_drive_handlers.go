package handlers

import (
	"encoding/base64"
	"mvp-2-spms/internal"
	"mvp-2-spms/web_server/handlers/interfaces"
	"net/http"
	"strconv"
	"strings"

	"mvp-2-spms/services/manage-accounts/inputdata"
	"mvp-2-spms/services/models"
)

type CloudDriveHandler struct {
	drives            internal.CloudDrives
	accountInteractor interfaces.IAccountInteractor
}

func InitCloudDriveHandler(drives internal.CloudDrives, acc interfaces.IAccountInteractor) CloudDriveHandler {
	return CloudDriveHandler{
		drives:            drives,
		accountInteractor: acc,
	}
}

func (h *CloudDriveHandler) GetGoogleDriveLink(w http.ResponseWriter, r *http.Request) {
	user := GetSessionUser(r)
	id, _ := strconv.Atoi(user.GetProfId())
	returnURL := r.URL.Query().Get("redirect")
	redirectURI := "http://127.0.0.1:8080/auth/integration/access/googledrive"
	result, _ := h.drives[models.GoogleDrive].GetAuthLink(redirectURI, int(uint(id)), returnURL)
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
	// http.Redirect(w, r, result, http.StatusTemporaryRedirect)
}

func (h *CloudDriveHandler) OAuthCallbackGoogleDrive(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	decodedState, _ := base64.URLEncoding.DecodeString(state)

	// needs further update
	params := strings.Split(string(decodedState), ",")
	accountId, _ := strconv.Atoi(params[0])
	redirect := params[1]

	input := inputdata.SetDriveIntegration{
		AccountId: uint(accountId),
		AuthCode:  code,
		Type:      int(models.GoogleDrive),
	}
	result := h.accountInteractor.SetDriveIntegration(input, h.drives[models.GoogleDrive])
	w.Header().Add("Google-Calendar-Token", result.AccessToken)
	w.Header().Add("Google-Calendar-Token-Exp", result.Expiry.String())
	http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
}
