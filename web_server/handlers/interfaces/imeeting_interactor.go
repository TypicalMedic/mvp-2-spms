package interfaces

import (
	"mvp-2-spms/services/manage-meetings/inputdata"
	"mvp-2-spms/services/manage-meetings/outputdata"
)

type IMeetingInteractor interface {
	AddMeeting(input inputdata.AddMeeting) outputdata.AddMeeting
	GetProfessorMeetings(input inputdata.GetProfessorMeetings) outputdata.GetProfesorMeetings
}
