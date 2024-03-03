package interfaces

import (
	entities "mvp-2-spms/domain-aggregate"
	"mvp-2-spms/services/models"
)

type ICloudDrive interface {
	AddProjectFolder(project entities.Project, driveInfo models.CloudDriveIntegration) models.DriveProject
	AddTaskToDrive(task entities.Task, projectFolderId string) models.DriveTask
}
