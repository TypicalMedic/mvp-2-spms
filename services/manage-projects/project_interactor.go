package manageprojects

import (
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
	projects := p.projectRepo.GetProfessorProjects(input.ProfessorId)
	for _, project := range projects {
		project.Student = p.studentRepo.GetStudentById(project.Student.Id)
	}
	output := outputdata.MapToGetProfessorProjects(projects)
	return output
}
