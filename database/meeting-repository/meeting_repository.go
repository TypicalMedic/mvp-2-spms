package meetingrepository

import (
	"fmt"
	"mvp-2-spms/database"
	"mvp-2-spms/database/models"
	entities "mvp-2-spms/domain-aggregate"
	usecasemodels "mvp-2-spms/services/models"
	"time"
)

type MeetingRepository struct {
	dbContext database.Database
}

func InitMeetingRepository(dbcxt database.Database) *MeetingRepository {
	return &MeetingRepository{
		dbContext: dbcxt,
	}
}

func (r *MeetingRepository) CreateMeeting(meeting entities.Meeting) entities.Meeting {
	dbmeeting := models.Meeting{}
	dbmeeting.MapEntityToThis(meeting)
	r.dbContext.DB.Create(&dbmeeting)
	return dbmeeting.MapToEntity()
}

func (r *MeetingRepository) AssignPlannerMeeting(meeting usecasemodels.PlannerMeeting) {
	r.dbContext.DB.Model(&models.Meeting{}).Where("id = ?", meeting.Meeting.Id).Update("planner_id", meeting.PlannerId)
}

func (r *MeetingRepository) GetProfessorMeetings(profId string, from time.Time) []entities.Meeting {
	var meetings []models.Meeting
	r.dbContext.DB.Select("*").Where("professor_id = ? AND meeting_time >= ?", profId, from).Order("meeting_time asc").Find(&meetings)
	result := []entities.Meeting{}
	for _, m := range meetings {
		result = append(result, m.MapToEntity())
	}
	return result
}

func (r *MeetingRepository) GetMeetingPlannerId(meetId string) string {
	meeting := models.Meeting{}
	r.dbContext.DB.Select("planner_id").Where("id = ?", meetId).Find(&meeting)
	return fmt.Sprint(meeting.PlannerId)
}
