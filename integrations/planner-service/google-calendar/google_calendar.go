package googlecalendar

import (
	"encoding/base64"
	"fmt"
	entities "mvp-2-spms/domain-aggregate"
	"mvp-2-spms/services/models"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

type GoogleCalendar struct {
	api googleCalendarApi
}

func InitGoogleCalendar(api googleCalendarApi) *GoogleCalendar {
	return &GoogleCalendar{api: api}
}

func (c *GoogleCalendar) AddMeeting(meeting entities.Meeting, plannerInfo models.PlannerIntegration) (models.PlannerMeeting, error) {
	event, _ := c.api.AddEvent(meeting.Time, meeting.Name, meeting.Description, plannerInfo.PlannerData.Id)
	return models.PlannerMeeting{
		Meeting:          meeting,
		MeetingPlannerId: event.Id,
	}, nil
}

func (c *GoogleCalendar) GetAllPlanners() ([]models.PlannerData, error) {
	planners, _ := c.api.GetAllCalendars()
	result := []models.PlannerData{}
	for _, pl := range planners.Items {
		result = append(result, models.PlannerData{
			Id:   pl.Id,
			Name: pl.Summary,
		})
	}
	return result, nil
}

func (c *GoogleCalendar) GetScheduleMeetinIds(from time.Time, plannerInfo models.PlannerIntegration) ([]string, error) {
	events, _ := c.api.GetSchedule(from, plannerInfo.PlannerData.Id)
	result := []string{}
	for _, event := range events.Items {
		result = append(result, strings.Split(event.Id, "_")[0])
	}
	return result, nil
}

func (c *GoogleCalendar) FindMeetingById(meetId string, plannerInfo models.PlannerIntegration) (bool, error) {
	event, _ := c.api.GetEventById(meetId, plannerInfo.PlannerData.Id)
	return event.Id != "", nil
}

func (c *GoogleCalendar) GetAuthLink(redirectURI string, accountId int, returnURL string) (string, error) {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// encode as JSON!
	statestr := base64.URLEncoding.EncodeToString([]byte(fmt.Sprint(accountId, ",", returnURL)))
	url, _ := c.api.GetAuthLink(redirectURI, statestr)
	return url, nil
}

func (c *GoogleCalendar) GetToken(code string) (*oauth2.Token, error) {
	token, _ := c.api.GetToken(code)
	return token, nil
}

func (c *GoogleCalendar) Authentificate(token *oauth2.Token) error {
	c.api.AuthentificateService(token)
	return nil
}
