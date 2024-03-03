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

func (m *MeetingInteractor) AddMeeting(input inputdata.AddMeeting) outputdata.AddMeeting {
	// adding meeting to db, returns created meeting (with id)
	meeting := m.meetingRepo.CreateMeeting(input.MapToMeetingEntity())
	// getting calendar info, should be checked for existance later
	plannerInfo := m.accountRepo.GetAccountPlannerData(fmt.Sprint(input.ProfessorId))
	// add meeting to calendar
	meeitngPlanner := m.plannerService.AddMeeting(meeting, plannerInfo)
	// add meeting id from planner
	m.meetingRepo.AssignPlannerMeeting(meeitngPlanner)
	// returning id
	output := outputdata.MapToAddMeeting(meeting)
	return output
}

func (m *MeetingInteractor) GetProfessorMeetings(input inputdata.GetProfessorMeetings) outputdata.GetProfesorMeetings {
	// get from db
	meetings := m.meetingRepo.GetProfessorMeetings(fmt.Sprint(input.ProfessorId), input.From)
	meetEntities := []outputdata.GetProfesorMeetingsEntities{}
	for _, meet := range meetings {
		student := m.studentRepo.GetStudentById(meet.ParticipantId)
		projTheme := m.projectRepo.GetStudentCurrentProjectTheme(meet.ParticipantId)
		// find meeting in cloud
		plannerId := m.meetingRepo.GetMeetingPlannerId(meet.Id)
		hasPlanner := m.plannerService.FindMeetingById(plannerId)
		meetEntities = append(meetEntities, outputdata.GetProfesorMeetingsEntities{
			Meeting:           meet,
			Student:           student,
			ProjectTheme:      projTheme,
			HasPlannerMeeting: hasPlanner,
		})
	}
	output := outputdata.MapToGetProfesorMeetings(meetEntities)
	return output
}
