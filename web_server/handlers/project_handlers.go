package handlers

import (
	"encoding/json"
	"mvp-2-spms/services/manage-projects/inputdata"
	"mvp-2-spms/web_server/handlers/interfaces"
	"net/http"
	"strconv"
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
	professorIdCookie, _ := r.Cookie("professor_id")
	professorId, _ := strconv.ParseUint(professorIdCookie.Value, 10, 32)
	input := inputdata.GetPfofessorProjects{
		ProfessorId: uint(professorId),
	}
	result := h.projectInteractor.GetProfessorProjects(input)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
