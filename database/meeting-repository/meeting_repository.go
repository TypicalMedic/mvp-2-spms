package meetingrepository

import (
	"mvp-2-spms/database"
	"mvp-2-spms/database/models"
	entities "mvp-2-spms/domain-aggregate"
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
