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

func (d *GoogleDrive) AddProjectFolder(project entities.Project, driveInfo models.CloudDriveIntegration) (models.DriveProject, error) {
	folderName := fmt.Sprint("Project ", project.Id, "_", project.Theme)
	folder, _ := d.api.CreateFolder(folderName, driveInfo.BaseFolderId)
	return models.DriveProject{
		Project: project,
		DriveFolder: models.DriveFolder{
			Id:   folder.Id,
			Link: folder.WebViewLink,
		},
	}, nil
}

func (d *GoogleDrive) AddProfessorBaseFolder() (models.DriveData, error) {
	folderName := "Student Project Management System Project Folder"
	existingFolders, _ := d.api.GetFoldersByName(folderName)
	count := 0
	for len(existingFolders.Files) != 0 {
		count++
		folderNameCopy := fmt.Sprint(folderName, " (", count, ")")
		existingFolders, _ = d.api.GetFoldersByName(folderNameCopy)
	}

	folder, _ := d.api.CreateFolder(fmt.Sprint(folderName, " (", count, ")"))
	return models.DriveData{
		BaseFolderId: folder.Id,
	}, nil
}

func (d *GoogleDrive) GetFolderNameById(id string) (string, error) {
	folder, _ := d.api.GetFolderById(id)
	return folder.Name, nil
}

func (d *GoogleDrive) AddTaskToDrive(task entities.Task, projectFolderId string) (models.DriveTask, error) {
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
	}, nil
}
func (d *GoogleDrive) GetAuthLink(redirectURI string, accountId int, returnURL string) (string, error) {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// encode as JSON!
	statestr := base64.URLEncoding.EncodeToString([]byte(fmt.Sprint(accountId, ",", returnURL)))
	url := d.api.GetAuthLink(redirectURI, statestr)
	return url, nil
}

func (d *GoogleDrive) GetToken(code string) (*oauth2.Token, error) {
	token := d.api.GetToken(code)
	return token, nil
}

func (d *GoogleDrive) Authentificate(token *oauth2.Token) error {
	d.api.AuthentificateService(token)
	return nil
}
