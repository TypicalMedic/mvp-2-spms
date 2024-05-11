package interfaces

import (
	entities "mvp-2-spms/domain-aggregate"
	"mvp-2-spms/services/models"
)

type ICloudDrive interface {
	IIntegration
	AddProjectFolder(project entities.Project, driveInfo models.CloudDriveIntegration) models.DriveProject
	GetFolderNameById(id string) string
	AddProfessorBaseFolder() models.DriveData
	AddTaskToDrive(task entities.Task, projectFolderId string) models.DriveTask
}
