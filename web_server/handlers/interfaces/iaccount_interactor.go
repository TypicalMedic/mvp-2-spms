package interfaces

import (
	"mvp-2-spms/services/interfaces"
	"mvp-2-spms/services/manage-accounts/inputdata"
	"mvp-2-spms/services/manage-accounts/outputdata"
)

type IAccountInteractor interface {
	GetPlannerIntegration(input inputdata.GetPlannerIntegration) outputdata.GetPlannerIntegration
	GetDriveIntegration(input inputdata.GetDriveIntegration) outputdata.GetDriveIntegration
	GetRepoHubIntegration(input inputdata.GetRepoHubIntegration) outputdata.GetRepoHubIntegration
	SetPlannerIntegration(input inputdata.SetPlannerIntegration, planner interfaces.IPlannerService) outputdata.SetPlannerIntegration
}
