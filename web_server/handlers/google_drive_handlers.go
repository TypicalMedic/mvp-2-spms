package handlers

import (
	"encoding/json"
	"mvp-2-spms/internal"
	"net/http"
)

type GoogleDriveHandler struct {
	drives internal.CloudDrives
}

func InitGoogleDriveHandler(drives internal.CloudDrives) GoogleDriveHandler {
	return GoogleDriveHandler{
		drives: drives,
	}
}

func (h *GoogleDriveHandler) GetLink(w http.ResponseWriter, r *http.Request) {
	redirectURI := r.URL.Query().Get("redirect")
	result := (*h.drives[internal.GoogleDrive]).GetAuthLink(redirectURI)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *GoogleDriveHandler) Auth(w http.ResponseWriter, r *http.Request) {
	cred := GetCredentials(r)
	(*h.drives[internal.GoogleDrive]).Authentificate(cred.GoogleDriveToken)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
