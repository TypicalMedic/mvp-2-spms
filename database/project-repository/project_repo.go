package projectrepository

import (
	"mvp-2-spms/database"
	"mvp-2-spms/database/models"
	entities "mvp-2-spms/domain-aggregate"
	usecaseModels "mvp-2-spms/services/models"
)

type ProjectRepository struct {
	dbContext database.Database
}

func InitProjectRepository(dbcxt database.Database) *ProjectRepository {
	return &ProjectRepository{
		dbContext: dbcxt,
	}
}

func (r *ProjectRepository) GetProfessorProjects(profId string) []entities.Project {
	var projects []models.Project
	r.dbContext.DB.Select("*").Where("supervisor_id = ?", profId).Find(&projects)
	result := []entities.Project{}
	for _, pj := range projects {
		// вынести в маппер
		result = append(result, pj.MapToEntity())
	}
	return result
}

func (r *ProjectRepository) GetProjectRepository(projId string) entities.ProjectInRepository {
	var project models.Project
	r.dbContext.DB.Select("repo_id").Where("id = ?", projId).Find(&project)
	var repo models.Repository
	r.dbContext.DB.Select("*").Where("id = ?", project.RepoId).Find(&repo)
	return repo.MapToEntity()
}

func (r *ProjectRepository) GetProjectById(projId string) entities.Project {
	var project models.Project
	r.dbContext.DB.Select("*").Where("id = ?", projId).Find(&project)
	return project.MapToEntity()
}

func (r *ProjectRepository) CreateProject(entities.Project) entities.Project

func (r *ProjectRepository) AssignDriveFolder(usecaseModels.DriveProject) {}
