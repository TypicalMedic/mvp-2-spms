package handlers

import (
	"encoding/json"
	"mvp-2-spms/internal"
	"net/http"

	"golang.org/x/oauth2"
)

type CloudDriveHandler struct {
	drives internal.CloudDrives
}

func InitCloudDriveHandler(drives internal.CloudDrives) CloudDriveHandler {
	return CloudDriveHandler{
		drives: drives,
	}
}

func (h *CloudDriveHandler) GetGoogleDriveLink(w http.ResponseWriter, r *http.Request) {
	redirectURI := r.URL.Query().Get("redirect")
	result := h.drives[internal.GoogleDrive].GetAuthLink(redirectURI, 0, "")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *CloudDriveHandler) AuthGoogleDrive(w http.ResponseWriter, r *http.Request) {
	// cred := GetCredentials(r)
	h.drives[internal.GoogleDrive].Authentificate(oauth2.Token{})
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
