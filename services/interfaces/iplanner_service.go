package interfaces

import (
	entities "mvp-2-spms/domain-aggregate"
	"mvp-2-spms/services/manage-meetings/models"
)

type IPlannerService interface {
	AddMeeting(meeting entities.Meeting, plannerInfo models.PlannerIntegration) models.PlannerMeeting
}
