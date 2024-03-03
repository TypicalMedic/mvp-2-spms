package googlecalendar

import (
	entities "mvp-2-spms/domain-aggregate"
	"mvp-2-spms/services/models"
)

type GoogleCalendar struct {
	api googleCalendarApi
}

func InitGoogleCalendar(api googleCalendarApi) GoogleCalendar {
	return GoogleCalendar{api: api}
}

func (c *GoogleCalendar) AddMeeting(meeting entities.Meeting, plannerInfo models.PlannerIntegration) models.PlannerMeeting {
	event, _ := c.api.AddEvent(meeting.Time, meeting.Name, meeting.Description, plannerInfo.PlannerData.Id)
	return models.PlannerMeeting{
		Meeting:   meeting,
		PlannerId: event.Id,
	}
}

func (c *GoogleCalendar) FindMeetingById(meetId string, plannerInfo models.PlannerIntegration) bool {
	_, err := c.api.GetEventById(meetId, plannerInfo.PlannerData.Id)
	return err == nil
}
