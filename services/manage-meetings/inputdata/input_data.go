package inputdata

import (
	"fmt"
	entities "mvp-2-spms/domain-aggregate"
	"time"
)

type AddMeeting struct {
	ProfessorId uint
	Name        string
	Description string
	MeetingTime time.Time
	StudentId   int
	IsOnline    bool
}

func (am *AddMeeting) MapToMeetingEntity() entities.Meeting {
	return entities.Meeting{
		OrganizerId:   fmt.Sprint(am.ProfessorId),
		Name:          am.Name,
		Description:   am.Description,
		ParticipantId: fmt.Sprint(am.StudentId),
		Time:          am.MeetingTime,
		IsOnline:      am.IsOnline,
		Status:        entities.MeetingStatus(entities.MeetingPlanned),
	}
}
