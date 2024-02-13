package projectrepository

import (
	"mvp-2-spms/database"
	"mvp-2-spms/database/models"
	"mvp-2-spms/domain/people"
	"mvp-2-spms/domain/project"
	"mvp-2-spms/domain/repositoryhub"
)

type ProjectRepository struct {
	DBContext database.Database
}

func (r *ProjectRepository) GetProfessorProjects(profId uint) []project.Project {
	var projects []models.Project
	r.DBContext.DB.Select("*").Where("supervisor_id = ?", profId).Find(&projects)
	result := []project.Project{}
	for _, pj := range projects {
		// вынести в маппер
		result = append(result,
			project.Project{
				Id:    pj.ID,
				Theme: pj.Theme,
				Supervisor: people.Professor{
					Person: people.Person{
						Id: profId,
					},
				},
				Student: people.Student{
					Person: people.Person{
						Id: pj.StudentId,
					},
				},
				Year:  pj.Year,
				Grade: pj.Grade,
				SupervisorReview: project.SupervisorReview{
					Id: pj.SupervisorReviewId,
				},
				Repository: repositoryhub.Repository{
					Id: pj.RepoId,
				},
				Stage:  project.ProjectStage(pj.StageId),
				Status: project.ProjectStatus(pj.StatusId),
			})
	}
	return result
}
