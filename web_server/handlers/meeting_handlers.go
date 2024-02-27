package handlers

import (
	"encoding/json"
	"mvp-2-spms/services/manage-meetings/inputdata"
	"mvp-2-spms/web_server/handlers/interfaces"
	requestbodies "mvp-2-spms/web_server/handlers/request-bodies"
	"net/http"
	"strconv"
)

type MeetingHandler struct {
	meetingInteractor interfaces.IMeetingInteractor
}

func InitMeetingHandler(meetInteractor interfaces.IMeetingInteractor) MeetingHandler {
	return MeetingHandler{
		meetingInteractor: meetInteractor,
	}
}

func (h *MeetingHandler) AddMeeting(w http.ResponseWriter, r *http.Request) {
	professorIdCookie, _ := r.Cookie("professor_id")
	professorId, _ := strconv.ParseUint(professorIdCookie.Value, 10, 32)

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

	input := inputdata.AddMeeting{
		ProfessorId: uint(professorId),
		Name:        reqB.Name,
		Description: reqB.Description,
		MeetingTime: reqB.MeetingTime,
		StudentId:   reqB.StudentId,
		IsOnline:    reqB.IsOnline,
	}

	meeting_id := h.meetingInteractor.AddMeeting(input)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meeting_id)
}
