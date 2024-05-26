package managetasks

import (
	"errors"
	"fmt"
	"mvp-2-spms/services/interfaces"
	"mvp-2-spms/services/manage-tasks/inputdata"
	"mvp-2-spms/services/manage-tasks/outputdata"
	"mvp-2-spms/services/models"

	"golang.org/x/oauth2"
)

type TaskInteractor struct {
	projectRepo interfaces.IProjetRepository
	taskRepo    interfaces.ITaskRepository
	accountRepo interfaces.IAccountRepository
}

func InitTaskInteractor(projRepo interfaces.IProjetRepository, taskRepo interfaces.ITaskRepository, accRepo interfaces.IAccountRepository) *TaskInteractor {
	return &TaskInteractor{
		projectRepo: projRepo,
		taskRepo:    taskRepo,
		accountRepo: accRepo,
	}
}

func (p *TaskInteractor) AddTask(input inputdata.AddTask, cloudDrive interfaces.ICloudDrive) (outputdata.AddTask, error) {
	// add to db
	task, err := p.taskRepo.CreateTask(input.MapToTaskEntity())
	if err != nil {
		return outputdata.AddTask{}, err
	}

	// get project folder id
	folderFound := true
	projFolder, err := p.projectRepo.GetProjectCloudFolderId(fmt.Sprint(input.ProjectId))
	if err != nil {
		if !errors.Is(err, models.ErrProjectCloudFolderNotFound) {
			return outputdata.AddTask{}, err
		}
		folderFound = false
	}

	if folderFound {
		// getting professor drive info, should be checked for existance later
		driveInfo, err := p.accountRepo.GetAccountDriveData(fmt.Sprint(input.ProfessorId))
		if err != nil {
			return outputdata.AddTask{}, err
		}

		//////////////////////////////////////////////////////////////////////////////////////////////////////
		// check for access token first????????????????????????????????????????????
		token := &oauth2.Token{
			RefreshToken: driveInfo.ApiKey,
		}

		err = cloudDrive.Authentificate(token)
		if err != nil {
			return outputdata.AddTask{}, err
		}

		// add folder to cloud (create folder and task file)
		driveTask, err := cloudDrive.AddTaskToDrive(task, projFolder)
		if err != nil {
			return outputdata.AddTask{}, err
		}

		// add folder id and file id from drive
		p.taskRepo.AssignDriveTask(driveTask)
	}

	// returning id
	output := outputdata.MapToAddTask(task)
	return output, nil
}

func (p *TaskInteractor) GetProjectTasks(input inputdata.GetProjectTasks) (outputdata.GetProjectTasks, error) {
	// get tasks from db
	tasks, err := p.taskRepo.GetProjectTasksWithCloud(fmt.Sprint(input.ProjectId))
	if err != nil {
		return outputdata.GetProjectTasks{}, err
	}

	output := outputdata.MapToGetProjectTasks(tasks)
	return output, nil
}
