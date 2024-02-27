package handlers

import (
	"encoding/json"
	"mvp-2-spms/services/manage-students/inputdata"
	"mvp-2-spms/web_server/handlers/interfaces"
	requestbodies "mvp-2-spms/web_server/handlers/request-bodies"
	"net/http"
	"strconv"
)

type StudentHandler struct {
	studentInteractor interfaces.IStudentInteractor
}

func InitStudentHandler(studInteractor interfaces.IStudentInteractor) StudentHandler {
	return StudentHandler{
		studentInteractor: studInteractor,
	}
}

func (h *StudentHandler) AddStudent(w http.ResponseWriter, r *http.Request) {
	professorIdCookie, _ := r.Cookie("professor_id")
	professorId, _ := strconv.ParseUint(professorIdCookie.Value, 10, 32)

	headerContentTtype := r.Header.Get("Content-Type")
	// проверяем соответсвтвие типа содержимого запроса
	if headerContentTtype != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	// декодируем тело запроса
	var reqB requestbodies.AddStudent
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	input := inputdata.AddStudent{
		ProfessorId:            uint(professorId),
		Name:                   reqB.Name,
		Surname:                reqB.Surname,
		Middlename:             reqB.Middlename,
		EducationalProgrammeId: uint(reqB.EducationalProgrammeId),
		Cource:                 uint(reqB.Cource),
	}

	student_id := h.studentInteractor.AddStudent(input)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student_id)
}
