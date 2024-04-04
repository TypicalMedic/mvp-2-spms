package handlers

import (
	"encoding/json"
	"mvp-2-spms/internal"
	ainputdata "mvp-2-spms/services/manage-accounts/inputdata"
	"mvp-2-spms/services/manage-tasks/inputdata"
	"mvp-2-spms/web_server/handlers/interfaces"
	requestbodies "mvp-2-spms/web_server/handlers/request-bodies"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type TaskHandler struct {
	taskInteractor    interfaces.ITaskInteractor
	accountInteractor interfaces.IAccountInteractor
	cloudDrives       internal.CloudDrives
}

func InitTaskHandler(taskInteractor interfaces.ITaskInteractor) TaskHandler {
	return TaskHandler{
		taskInteractor: taskInteractor,
	}
}

func (h *TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	cred := GetCredentials(r)
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

	integInput := ainputdata.GetDriveIntegration{
		AccountId: cred.ProfessorId,
	}
	driveInfo := h.accountInteractor.GetDriveIntegration(integInput)
	input := inputdata.AddTask{
		ProfessorId: cred.ProfessorId,
		Name:        reqB.Name,
		Description: reqB.Description,
		Deadline:    reqB.Deadline,
		ProjectId:   uint(projectId),
	}

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO: pass api key/clone with new key///////////////////////////////////////////////////////////////////////////////
	task_id := h.taskInteractor.AddTask(input, h.cloudDrives[internal.CloudDriveName(driveInfo.Type)])
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task_id)
}

func (h *TaskHandler) GetAllProjectTasks(w http.ResponseWriter, r *http.Request) {
	cred := GetCredentials(r)
	projectId, _ := strconv.ParseUint(chi.URLParam(r, "projectID"), 10, 32)
	input := inputdata.GetProjectTasks{
		ProfessorId: cred.ProfessorId,
		ProjectId:   uint(projectId),
	}
	result := h.taskInteractor.GetProjectTasks(input)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
