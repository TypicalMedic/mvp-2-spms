package handlers

import (
	"encoding/json"
	"mvp-2-spms/internal"
	ainputdata "mvp-2-spms/services/manage-accounts/inputdata"
	minputdata "mvp-2-spms/services/manage-meetings/inputdata"
	"mvp-2-spms/services/models"
	"mvp-2-spms/web_server/handlers/interfaces"
	requestbodies "mvp-2-spms/web_server/handlers/request-bodies"
	"net/http"
	"time"
)

type MeetingHandler struct {
	meetingInteractor interfaces.IMeetingInteractor
	accountInteractor interfaces.IAccountInteractor
	planners          internal.Planners
}

func InitMeetingHandler(meetInteractor interfaces.IMeetingInteractor, acc interfaces.IAccountInteractor, pl internal.Planners) MeetingHandler {
	return MeetingHandler{
		meetingInteractor: meetInteractor,
		accountInteractor: acc,
		planners:          pl,
	}
}

func (h *MeetingHandler) AddMeeting(w http.ResponseWriter, r *http.Request) {
	cred := GetCredentials(r)

	headerContentTtype := r.Header.Get("Content-Type")
	// проверяем соответсвтвие типа содержимого запроса
	if headerContentTtype != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	// декодируем тело запроса
	var reqB requestbodies.AddMeeting
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	integInput := ainputdata.GetPlannerIntegration{
		AccountId: cred.ProfessorId,
	}
	calendarInfo := h.accountInteractor.GetPlannerIntegration(integInput)
	meetingInput := minputdata.AddMeeting{
		ProfessorId: cred.ProfessorId,
		Name:        reqB.Name,
		Description: reqB.Description,
		MeetingTime: reqB.MeetingTime,
		StudentId:   reqB.StudentId,
		IsOnline:    reqB.IsOnline,
		ProjectId:   uint(reqB.ProjectId),
	}

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO: pass api key/clone with new key///////////////////////////////////////////////////////////////////////////////
	meeting_id := h.meetingInteractor.AddMeeting(meetingInput, h.planners[models.PlannerName(calendarInfo.Type)])
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meeting_id)
}

func (h *MeetingHandler) GetProfessorMeetings(w http.ResponseWriter, r *http.Request) {
	cred := GetCredentials(r)
	from, _ := time.Parse("2006-01-02T15:04:05.000Z", r.URL.Query().Get("from"))
	input := minputdata.GetProfessorMeetings{
		ProfessorId: cred.ProfessorId,
		From:        from,
	}

	integInput := ainputdata.GetPlannerIntegration{
		AccountId: cred.ProfessorId,
	}
	calendarInfo := h.accountInteractor.GetPlannerIntegration(integInput)
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO: pass api key/clone with new key///////////////////////////////////////////////////////////////////////////////
	result := h.meetingInteractor.GetProfessorMeetings(input, h.planners[models.PlannerName(calendarInfo.Type)])
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
