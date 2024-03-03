package googlecalendar

import (
	"log"
	googleapi "mvp-2-spms/integrations/google-api"
	"time"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const DAYS_PERIOD = 7
const HOURS_IN_DAY = 24
const EVENT_DURATION_HOURS = 1

type googleCalendarApi struct {
	api *calendar.Service
}

func InitCalendarApi(googleAPI googleapi.GoogleAPI) googleCalendarApi {
	api, err := calendar.NewService(googleAPI.Context, option.WithHTTPClient(googleAPI.Client))
	c := googleCalendarApi{api}
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}
	return c
}

// startTime should be UTC+0!!!
func (c *googleCalendarApi) AddEvent(startTime time.Time, summary string, desc string, calendarId string) (*calendar.Event, error) {
	endTime := startTime.Add(EVENT_DURATION_HOURS * time.Hour).Format(time.RFC3339)
	event := &calendar.Event{
		Summary:     summary,
		Description: desc,
		Start: &calendar.EventDateTime{
			DateTime: startTime.Format(time.RFC3339),
			TimeZone: "Asia/Yekaterinburg",
		},
		End: &calendar.EventDateTime{
			DateTime: endTime,
			TimeZone: "Asia/Yekaterinburg", //////////////////////////////////////?????????????????????????????????????
		},
		Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=1"},
	}
	result, err := c.api.Events.Insert(calendarId, event).Do()
	if err == nil {
		return result, nil
	}
	return nil, err
}

func (c *googleCalendarApi) GetEventById(eventId string, calendarId string) (*calendar.Event, error) {
	event, err := c.api.Events.Get(calendarId, eventId).Do()
	if err == nil {
		return event, nil
	}
	return nil, err
}

func (c *googleCalendarApi) GetSchedule(startTime time.Time, calendarId string) (*calendar.Events, error) {
	events, err := c.api.Events.List("marusya.pletneva2012@gmail.com").ShowDeleted(false).SingleEvents(true).TimeMin(startTime.Format(time.RFC3339)).OrderBy("startTime").Do()
	if err == nil {
		return events, nil
	}
	return nil, err
}
