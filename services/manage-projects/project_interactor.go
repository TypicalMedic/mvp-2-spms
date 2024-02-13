package manageprojects

import (
	"mvp-2-spms/services/manage-projects/inputdata"
	"mvp-2-spms/services/manage-projects/interfaces"
	"mvp-2-spms/services/manage-projects/outputdata"
)

type ProjectInteractor struct {
	projectRepo interfaces.IProjetRepository
}

// returns all professor projects (basic information)
func (p *ProjectInteractor) GetPfofessorProjects(input inputdata.GetPfofessorProjects) outputdata.GetPfofessorProjects {
	// get from database
	projects := p.projectRepo.GetProfessorProjects(input.ProfessorId)
	outputProjects := []outputdata.ProjectData{}
	for _, project := range projects {
		outputProjects = append(outputProjects,
			outputdata.ProjectData{
				Id:          project.Id,
				Theme:       project.Theme,
				Status:      project.Status.String(),
				Stage:       project.Stage.String(),
				Year:        project.Year,
				StudentName: project.Student.FullNameToString(),
				Cource:      project.Student.GetCource(),
			})
	}
	output := outputdata.GetPfofessorProjects{
		Projects: outputProjects,
	}
	return output
}
