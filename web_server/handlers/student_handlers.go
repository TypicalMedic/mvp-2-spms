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
	user := GetSessionUser(r)
	id, _ := strconv.Atoi(user.GetProfId())

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
		ProfessorId:            uint(id),
		Name:                   reqB.Name,
		Surname:                reqB.Surname,
		Middlename:             reqB.Middlename,
		EducationalProgrammeId: uint(reqB.EducationalProgrammeId),
		Cource:                 uint(reqB.Cource),
	}

	student_id, _ := h.studentInteractor.AddStudent(input)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student_id)
}

func (h *StudentHandler) GetStudents(w http.ResponseWriter, r *http.Request) {
	user := GetSessionUser(r)
	id, _ := strconv.Atoi(user.GetProfId())
	input := inputdata.GetStudents{
		ProfessorId: uint(id),
	}
	result, _ := h.studentInteractor.GetStudents(input)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
