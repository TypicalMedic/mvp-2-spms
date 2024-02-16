package projectrepository

import (
	"mvp-2-spms/database"
	"mvp-2-spms/database/models"
	"mvp-2-spms/domain/project"
)

type ProjectRepository struct {
	dbContext database.Database
}

func InitProjectRepository(dbcxt database.Database) *ProjectRepository {
	return &ProjectRepository{
		dbContext: dbcxt,
	}
}

func (r *ProjectRepository) GetProfessorProjects(profId uint) []project.Project {
	var projects []models.Project
	r.dbContext.DB.Select("*").Where("supervisor_id = ?", profId).Find(&projects)
	result := []project.Project{}
	for _, pj := range projects {
		// вынести в маппер
		result = append(result, pj.MapToEntity())
	}
	return result
}
