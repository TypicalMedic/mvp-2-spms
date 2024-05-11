package interfaces

import (
	"mvp-2-spms/services/interfaces"
	"mvp-2-spms/services/manage-accounts/inputdata"
	"mvp-2-spms/services/manage-accounts/outputdata"
)

type IAccountInteractor interface {
	GetDriveBaseFolderName(folderId, profId string, cloudDrive interfaces.ICloudDrive) string
	GetPlannerIntegration(input inputdata.GetPlannerIntegration) outputdata.GetPlannerIntegration
	GetDriveIntegration(input inputdata.GetDriveIntegration) outputdata.GetDriveIntegration
	GetRepoHubIntegration(input inputdata.GetRepoHubIntegration) outputdata.GetRepoHubIntegration
	SetPlannerIntegration(input inputdata.SetPlannerIntegration, planner interfaces.IPlannerService) outputdata.SetPlannerIntegration
	SetDriveIntegration(input inputdata.SetDriveIntegration, planner interfaces.ICloudDrive) outputdata.SetDriveIntegration
	SetRepoHubIntegration(input inputdata.SetRepoHubIntegration, planner interfaces.IGitRepositoryHub) outputdata.SetRepoHubIntegration
	GetAccountIntegrations(input inputdata.GetAccountIntegrations) outputdata.GetAccountIntegrations
	GetProfessorInfo(input inputdata.GetProfessorInfo) outputdata.GetProfessorInfo
	CheckCredsValidity(input inputdata.CheckCredsValidity) bool
	CheckUsernameExists(input inputdata.CheckUsernameExists) bool
	SignUp(input inputdata.SignUp) outputdata.SignUp
	GetAccountProfessorId(login string) string
	SetProfessorPlanner(plannerId, profId string)
	GetProfessorIntegrPlanners(profId string, planner interfaces.IPlannerService) outputdata.GetProfessorIntegrPlanners
}
