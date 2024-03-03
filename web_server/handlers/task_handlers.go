package handlers

import (
	"encoding/json"
	"mvp-2-spms/services/manage-tasks/inputdata"
	"mvp-2-spms/web_server/handlers/interfaces"
	requestbodies "mvp-2-spms/web_server/handlers/request-bodies"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type TaskHandler struct {
	taskInteractor interfaces.ITaskInteractor
}

func InitTaskHandler(taskInteractor interfaces.ITaskInteractor) TaskHandler {
	return TaskHandler{
		taskInteractor: taskInteractor,
	}
}

func (h *TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	professorIdCookie, _ := r.Cookie("professor_id")
	professorId, _ := strconv.ParseUint(professorIdCookie.Value, 10, 32)
	projectId, _ := strconv.ParseUint(chi.URLParam(r, "projectID"), 10, 32)

	headerContentTtype := r.Header.Get("Content-Type")
	// проверяем соответсвтвие типа содержимого запроса
	if headerContentTtype != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	// декодируем тело запроса
	var reqB requestbodies.AddTask
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	input := inputdata.AddTask{
		ProfessorId: uint(professorId),
		Name:        reqB.Name,
		Description: reqB.Description,
		Deadline:    reqB.Deadline,
		ProjectId:   uint(projectId),
	}

	task_id := h.taskInteractor.AddTask(input)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task_id)
}
