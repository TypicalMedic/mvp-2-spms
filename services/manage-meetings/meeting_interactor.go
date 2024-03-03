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
	projectRepo    interfaces.IProjetRepository
	studentRepo    interfaces.IStudentRepository
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

func (p *MeetingInteractor) GetProfessorMeetings(input inputdata.GetProfessorMeetings) outputdata.GetProfesorMeetings {
	// get from db
	meetings := p.meetingRepo.GetProfessorMeetings(fmt.Sprint(input.ProfessorId), input.From)
	meetEntities := []outputdata.GetProfesorMeetingsEntities{}
	for _, meet := range meetings {
		student := p.studentRepo.GetStudentById(meet.ParticipantId)
		projTheme := p.projectRepo.GetStudentCurrentProjectTheme(meet.ParticipantId)
		meetEntities = append(meetEntities, outputdata.GetProfesorMeetingsEntities{
			Meeting:      meet,
			Student:      student,
			ProjectTheme: projTheme,
		})
	}
	output := outputdata.MapToGetProfesorMeetings(meetEntities)
	return output
}
