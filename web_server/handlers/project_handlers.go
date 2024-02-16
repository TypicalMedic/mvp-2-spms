package handlers

import (
	"encoding/json"
	"mvp-2-spms/services/manage-projects/inputdata"
	"mvp-2-spms/web_server/handlers/interfaces"
	"net/http"
)

type ProjectHandler struct {
	projectInteractor interfaces.IProjetInteractor
}

func InitProjectHandler(projInteractor interfaces.IProjetInteractor) ProjectHandler {
	return ProjectHandler{
		projectInteractor: projInteractor,
	}
}

func (h *ProjectHandler) GetAllProfProjects(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var input inputdata.GetPfofessorProjects
	decoder.Decode(&input)
	result := h.projectInteractor.GetProfessorProjects(input)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
