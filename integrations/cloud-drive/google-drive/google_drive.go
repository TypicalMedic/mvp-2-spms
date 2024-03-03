package clouddrive

import (
	"fmt"
	entities "mvp-2-spms/domain-aggregate"
	"mvp-2-spms/services/models"
)

type GoogleDrive struct {
	api googleDriveApi
}

func InitGoogleDrive(api googleDriveApi) GoogleDrive {
	return GoogleDrive{api: api}
}

func (c *GoogleDrive) AddProjectFolder(project entities.Project, driveInfo models.CloudDriveIntegration) models.DriveProject {
	folderName := fmt.Sprint("Project ", project.Id, "_", project.Theme)
	folder, _ := c.api.CreateFolder(folderName, driveInfo.BaseFolderId)
	return models.DriveProject{
		Project:         project,
		ProjectFolderId: folder.Id,
	}
}

func (c *GoogleDrive) AddTaskToDrive(task entities.Task, projectFolderId string) models.DriveTask
