package taskrepository

import (
	"mvp-2-spms/database"
	"mvp-2-spms/database/models"
	entities "mvp-2-spms/domain-aggregate"
	usecaseModels "mvp-2-spms/services/models"
)

type TaskRepository struct {
	dbContext database.Database
}

func InitTaskRepository(dbcxt database.Database) *TaskRepository {
	return &TaskRepository{
		dbContext: dbcxt,
	}
}

func (r *TaskRepository) CreateTask(task entities.Task) entities.Task {
	dbtask := models.Task{}
	dbtask.MapEntityToThis(task)
	r.dbContext.DB.Create(&dbtask)
	return dbtask.MapToEntity()
}

func (r *TaskRepository) AssignDriveTask(task usecaseModels.DriveTask) {
	dbCloudFolder := models.CloudFolder{}
	dbCloudFolder.MapUseCaseModelToThis(task.DriveFolder)
	r.dbContext.DB.Create(&dbCloudFolder)
	r.dbContext.DB.Model(&models.Task{}).Select("folder_id", "task_file_id").Where("id = ?", task.Task.Id).Updates(
		models.Task{
			FolderId:   task.DriveFolder.Id,
			TaskFileId: task.TaskFileId,
		})
}

func (r *TaskRepository) GetProjectTasks(projId string) []entities.Task {
	var tasks []models.Task
	r.dbContext.DB.Select("*").Where("project_id = ?", projId).Find(&tasks)
	result := []entities.Task{}
	for _, t := range tasks {
		result = append(result, t.MapToEntity())
	}
	return result
}
