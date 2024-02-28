package managemeetings

import (
	"fmt"
	"mvp-2-spms/services/interfaces"
	"mvp-2-spms/services/manage-meetings/inputdata"
	"mvp-2-spms/services/manage-meetings/outputdata"
)

type MeetingInteractor struct {
	meetingRepo    interfaces.IMeetingRepository
	plannerService interfaces.IPlannerService
	accountRepo    interfaces.IAccountRepository
}

func InitMeetingInteractor(mtRepo interfaces.IMeetingRepository, planner interfaces.IPlannerService, accRepo interfaces.IAccountRepository) *MeetingInteractor {
	return &MeetingInteractor{
		meetingRepo:    mtRepo,
		plannerService: planner,
		accountRepo:    accRepo,
	}
}

func (p *MeetingInteractor) AddMeeting(input inputdata.AddMeeting) outputdata.AddMeeting {
	// adding meeting to db, returns created meeting (with id)
	meeting := p.meetingRepo.CreateMeeting(input.MapToMeetingEntity())
	// getting calendar info, should be checked for existance later
	plannerInfo := p.accountRepo.GetAccountPlannerData(fmt.Sprint(input.ProfessorId))
	// add meeting to calendar
	meeitngPlanner := p.plannerService.AddMeeting(meeting, plannerInfo)
	// add meeting id from planner
	p.meetingRepo.AssignPlannerMeeting(meeitngPlanner)
	// returning id
	output := outputdata.MapToAddMeeting(meeting)
	return output
}
