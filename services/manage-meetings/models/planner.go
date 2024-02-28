package models

import (
	entities "mvp-2-spms/domain-aggregate"
)

type PlannerData struct {
	Id string
	//Owner string
}

type PlannerMeeting struct {
	entities.Meeting
	PlannerId string
}
