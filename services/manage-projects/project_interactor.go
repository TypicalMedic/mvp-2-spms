package manageprojects

import (
	"fmt"
	"mvp-2-spms/services/manage-projects/inputdata"
	"mvp-2-spms/services/manage-projects/interfaces"
	"mvp-2-spms/services/manage-projects/outputdata"
)

type ProjectInteractor struct {
	projectRepo interfaces.IProjetRepository
	studentRepo interfaces.IStudentRepository
}

func InitProjectInteractor(projRepo interfaces.IProjetRepository, stRepo interfaces.IStudentRepository) *ProjectInteractor {
	return &ProjectInteractor{
		projectRepo: projRepo,
		studentRepo: stRepo,
	}
}

// returns all professor projects (basic information)
func (p *ProjectInteractor) GetProfessorProjects(input inputdata.GetPfofessorProjects) outputdata.GetProfessorProjects {
	// get from database
	projEntities := []outputdata.GetProfessorProjectsEntities{}
	projects := p.projectRepo.GetProfessorProjects(fmt.Sprint(input.ProfessorId))
	for _, project := range projects {
		student := p.studentRepo.GetStudentById(project.StudentId)
		projEntities = append(projEntities, outputdata.GetProfessorProjectsEntities{
			Project: project,
			Student: student,
		})
	}
	output := outputdata.MapToGetProfessorProjects(projEntities)
	return output
}
