package managetasks

import (
	"fmt"
	"mvp-2-spms/services/interfaces"
	"mvp-2-spms/services/manage-tasks/inputdata"
	"mvp-2-spms/services/manage-tasks/outputdata"
)

type TaskInteractor struct {
	projectRepo interfaces.IProjetRepository
	taskRepo    interfaces.ITaskRepository
}

func InitTaskInteractor(projRepo interfaces.IProjetRepository, taskRepo interfaces.ITaskRepository) *TaskInteractor {
	return &TaskInteractor{
		projectRepo: projRepo,
		taskRepo:    taskRepo,
	}
}

func (p *TaskInteractor) AddTask(input inputdata.AddTask, cloudDrive interfaces.ICloudDrive) outputdata.AddTask {
	// add to db
	task := p.taskRepo.CreateTask(input.MapToTaskEntity())
	// get project folder id
	projFolder := p.projectRepo.GetProjectCloudFolderId(fmt.Sprint(input.ProjectId))
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
	tasks := p.taskRepo.GetProjectTasksWithCloud(fmt.Sprint(input.ProjectId))
	output := outputdata.MapToGetProjectTasks(tasks)
	return output
}
