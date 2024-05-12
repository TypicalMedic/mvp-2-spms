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

func (r *MeetingRepository) CreateMeeting(meeting entities.Meeting) (entities.Meeting, error) {
	dbmeeting := models.Meeting{}
	dbmeeting.MapEntityToThis(meeting)
	r.dbContext.DB.Create(&dbmeeting)
	return dbmeeting.MapToEntity(), nil
}

func (r *MeetingRepository) AssignPlannerMeeting(meeting usecasemodels.PlannerMeeting) error {
	r.dbContext.DB.Model(&models.Meeting{}).Where("id = ?", meeting.Meeting.Id).Update("planner_id", meeting.MeetingPlannerId)
	return nil
}

func (r *MeetingRepository) GetProfessorMeetings(profId string, from time.Time, to time.Time) ([]entities.Meeting, error) {
	var meetings []models.Meeting
	if to.IsZero() {
		r.dbContext.DB.Select("*").Where("professor_id = ? AND meeting_time > ?", profId, from).Order("meeting_time asc").Find(&meetings)
	} else {
		r.dbContext.DB.Select("*").Where("professor_id = ? AND meeting_time > ? AND meeting_time < ?", profId, from, to).Order("meeting_time asc").Find(&meetings)
	}
	result := []entities.Meeting{}
	for _, m := range meetings {
		result = append(result, m.MapToEntity())
	}
	return result, nil
}

func (r *MeetingRepository) GetMeetingPlannerId(meetId string) (string, error) {
	meeting := models.Meeting{}
	r.dbContext.DB.Select("planner_id").Where("id = ?", meetId).Find(&meeting)
	return fmt.Sprint(meeting.PlannerId), nil
}
