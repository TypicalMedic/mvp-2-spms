package interfaces

import (
	"mvp-2-spms/services/manage-accounts/inputdata"
	"mvp-2-spms/services/manage-accounts/outputdata"
)

type IAccountInteractor interface {
	GetPlannerIntegration(input inputdata.GetPlannerIntegration) outputdata.GetPlannerIntegration
	GetDriveIntegration(input inputdata.GetDriveIntegration) outputdata.GetDriveIntegration
	GetrepoHubIntegration(input inputdata.GetRepoHubIntegration) outputdata.GetRepoHubIntegration
}
