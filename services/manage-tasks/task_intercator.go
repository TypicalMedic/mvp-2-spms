package managetasks

import (
	"fmt"
	"mvp-2-spms/services/interfaces"
	"mvp-2-spms/services/manage-tasks/inputdata"
	"mvp-2-spms/services/manage-tasks/outputdata"

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

func (p *TaskInteractor) AddTask(input inputdata.AddTask, cloudDrive interfaces.ICloudDrive) outputdata.AddTask {
	// add to db
	task, _ := p.taskRepo.CreateTask(input.MapToTaskEntity())
	// get project folder id
	projFolder, _ := p.projectRepo.GetProjectCloudFolderId(fmt.Sprint(input.ProjectId))

	// getting professor drive info, should be checked for existance later
	driveInfo, _ := p.accountRepo.GetAccountDriveData(fmt.Sprint(input.ProfessorId))

	//////////////////////////////////////////////////////////////////////////////////////////////////////
	// check for access token first????????????????????????????????????????????
	token := &oauth2.Token{
		RefreshToken: driveInfo.ApiKey,
	}
	cloudDrive.Authentificate(token)

	// add folder to cloud (create folder and task file)
	driveTask := cloudDrive.AddTaskToDrive(task, projFolder)
	// add folder id and file id from drive
	p.taskRepo.AssignDriveTask(driveTask)
	// returning id
	output := outputdata.MapToAddTask(task)
	return output
}

func (p *TaskInteractor) GetProjectTasks(input inputdata.GetProjectTasks) outputdata.GetProjectTasks {
	// get tasks from db
	tasks, _ := p.taskRepo.GetProjectTasksWithCloud(fmt.Sprint(input.ProjectId))
	output := outputdata.MapToGetProjectTasks(tasks)
	return output
}
