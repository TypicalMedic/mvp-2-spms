package handlers

import (
	"encoding/json"
	"mvp-2-spms/services/manage-universities/inputdata"
	"mvp-2-spms/web_server/handlers/interfaces"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UniversityHandler struct {
	uniInteractor interfaces.IUniversityInteractor
}

func InitUniversityHandler(uInteractor interfaces.IUniversityInteractor) UniversityHandler {
	return UniversityHandler{
		uniInteractor: uInteractor,
	}
}

func (h *UniversityHandler) GetAllUniEdProgrammes(w http.ResponseWriter, r *http.Request) {
	cred := GetCredentials(r)
	uniId, _ := strconv.ParseUint(chi.URLParam(r, "uniID"), 10, 32)
	input := inputdata.GetUniEducationalProgrammes{
		ProfessorId:  cred.ProfessorId,
		UniversityId: uint(uniId),
	}
	result := h.uniInteractor.GetUniEdProgrammes(input)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
