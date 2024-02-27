package managemeetings

import (
	"mvp-2-spms/services/interfaces"
	"mvp-2-spms/services/manage-meetings/inputdata"
	"mvp-2-spms/services/manage-meetings/outputdata"
)

type MeetingInteractor struct {
	meetingRepo    interfaces.IMeetingRepository
	plannerService interfaces.IPlennerService
}

func InitMeetingInteractor(mtRepo interfaces.IMeetingRepository, planner interfaces.IPlennerService) *MeetingInteractor {
	return &MeetingInteractor{
		meetingRepo:    mtRepo,
		plannerService: planner,
	}
}

func (p *MeetingInteractor) AddMeeting(input inputdata.AddMeeting) outputdata.AddMeeting {
	// adding meeting to db, returns created meeting (with id)
	meeting := p.meetingRepo.CreateMeeting(input.MapToMeetingEntity())
	// add meeting to calendar
	p.plannerService.AddMeeting(meeting)
	// returning id
	output := outputdata.MapToAddMeeting(meeting)
	return output
}
