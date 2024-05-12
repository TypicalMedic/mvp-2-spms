package managemeetings

import (
	"fmt"
	"mvp-2-spms/services/interfaces"
	"mvp-2-spms/services/manage-meetings/inputdata"
	"mvp-2-spms/services/manage-meetings/outputdata"
	"slices"

	"golang.org/x/oauth2"
)

type MeetingInteractor struct {
	meetingRepo interfaces.IMeetingRepository
	accountRepo interfaces.IAccountRepository
	projectRepo interfaces.IProjetRepository
	studentRepo interfaces.IStudentRepository
}

func InitMeetingInteractor(mtRepo interfaces.IMeetingRepository, accRepo interfaces.IAccountRepository,
	sRepo interfaces.IStudentRepository, projRepo interfaces.IProjetRepository) *MeetingInteractor {
	return &MeetingInteractor{
		meetingRepo: mtRepo,
		accountRepo: accRepo,
		studentRepo: sRepo,
		projectRepo: projRepo,
	}
}

func (m *MeetingInteractor) AddMeeting(input inputdata.AddMeeting, planner interfaces.IPlannerService) (outputdata.AddMeeting, error) {
	// adding meeting to db, returns created meeting (with id)
	meeting, _ := m.meetingRepo.CreateMeeting(input.MapToMeetingEntity())
	// getting calendar info, should be checked for existance later
	plannerInfo, _ := m.accountRepo.GetAccountPlannerData(fmt.Sprint(input.ProfessorId))
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	// check for access token first????????????????????????????????????????????
	token := &oauth2.Token{
		RefreshToken: plannerInfo.ApiKey,
	}
	planner.Authentificate(token)
	// add meeting to calendar
	meeitngPlanner, _ := planner.AddMeeting(meeting, plannerInfo)
	// add meeting id from planner
	m.meetingRepo.AssignPlannerMeeting(meeitngPlanner)
	// returning id
	output := outputdata.MapToAddMeeting(meeting)
	return output, nil
}

func (m *MeetingInteractor) GetProfessorMeetings(input inputdata.GetProfessorMeetings, planner interfaces.IPlannerService) (outputdata.GetProfesorMeetings, error) {
	// get from db
	meetings, _ := m.meetingRepo.GetProfessorMeetings(fmt.Sprint(input.ProfessorId), input.From, input.To)
	meetEntities := []outputdata.GetProfesorMeetingsEntities{}
	// getting calendar info, should be checked for existance later
	plannerInfo, _ := m.accountRepo.GetAccountPlannerData(fmt.Sprint(input.ProfessorId))
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	// check for access token first????????????????????????????????????????????
	token := &oauth2.Token{
		RefreshToken: plannerInfo.ApiKey,
	}
	planner.Authentificate(token)
	plannerMetingsIds, _ := planner.GetScheduleMeetinIds(input.From, plannerInfo)
	for _, meet := range meetings {
		student, _ := m.studentRepo.GetStudentById(meet.ParticipantId)
		proj, _ := m.projectRepo.GetStudentCurrentProject(meet.ParticipantId)
		// getting planner meeting id
		plannerId, _ := m.meetingRepo.GetMeetingPlannerId(meet.Id)
		// check if meeting exists in planner
		hasPlanner := slices.Contains(plannerMetingsIds, plannerId)
		meetEntities = append(meetEntities, outputdata.GetProfesorMeetingsEntities{
			Meeting:           meet,
			Student:           student,
			Project:           proj,
			HasPlannerMeeting: hasPlanner,
		})
	}
	output := outputdata.MapToGetProfesorMeetings(meetEntities)
	return output, nil
}
