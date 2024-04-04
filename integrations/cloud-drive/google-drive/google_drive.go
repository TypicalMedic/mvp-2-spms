package clouddrive

import (
	"encoding/base64"
	"fmt"
	entities "mvp-2-spms/domain-aggregate"
	"mvp-2-spms/services/models"

	"golang.org/x/oauth2"
)

type GoogleDrive struct {
	api googleDriveApi
}

func InitGoogleDrive(api googleDriveApi) *GoogleDrive {
	return &GoogleDrive{api: api}
}

func (d *GoogleDrive) AddProjectFolder(project entities.Project, driveInfo models.CloudDriveIntegration) models.DriveProject {
	folderName := fmt.Sprint("Project ", project.Id, "_", project.Theme)
	folder, _ := d.api.CreateFolder(folderName, driveInfo.BaseFolderId)
	return models.DriveProject{
		Project: project,
		DriveFolder: models.DriveFolder{
			Id:   folder.Id,
			Link: folder.WebViewLink,
		},
	}
}

func (d *GoogleDrive) AddTaskToDrive(task entities.Task, projectFolderId string) models.DriveTask {
	// add task folder
	folderName := fmt.Sprint("Task ", task.Id, "_", task.Name, " until: ", task.Deadline.Format("02.01.2006"))
	folder, _ := d.api.CreateFolder(folderName, projectFolderId)
	// add task file
	fileName := fmt.Sprint("Task '", task.Name, "' desctiprion")
	text := fmt.Sprint(task.Name, "\n\n", task.Description)
	file, _ := d.api.AddTextFileToFolder(fileName, text, folder.Id)
	return models.DriveTask{
		Task: task,
		DriveFolder: models.DriveFolder{
			Id:   folder.Id,
			Link: folder.WebViewLink,
		},
		TaskFileId: file.Id,
	}
}
func (d *GoogleDrive) GetAuthLink(redirectURI string, accountId int, returnURL string) string {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// encode as JSON!
	statestr := base64.URLEncoding.EncodeToString([]byte(fmt.Sprint(accountId, ",", returnURL)))
	url := d.api.GetAuthLink(redirectURI, statestr)
	return url
}

func (d *GoogleDrive) GetToken(code string) oauth2.Token {
	token := d.api.GetToken(code)
	return token
}

func (d *GoogleDrive) Authentificate(token oauth2.Token) {
	d.api.AuthentificateService(token)
}
