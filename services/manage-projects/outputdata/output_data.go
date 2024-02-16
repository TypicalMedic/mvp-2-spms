package outputdata

import "mvp-2-spms/domain/project"

type GetProfessorProjects struct {
	Projects []getProfProjProjectData `json:"projects"`
}

func MapToGetProfessorProjects(projectEntities []project.Project) GetProfessorProjects {
	outputProjects := []getProfProjProjectData{}
	for _, project := range projectEntities {
		outputProjects = append(outputProjects,
			getProfProjProjectData{
				Id:          project.Id,
				Theme:       project.Theme,
				Status:      project.Status.String(),
				Stage:       project.Stage.String(),
				Year:        project.Year,
				StudentName: project.Student.FullNameToString(),
				Cource:      project.Student.GetCource(),
			})
	}
	return GetProfessorProjects{
		Projects: outputProjects,
	}
}

type getProfProjProjectData struct {
	Id          uint   `json:"id"`
	Theme       string `json:"theme"`
	StudentName string `json:"student_name"`
	Cource      uint   `json:"cource"`
	Status      string `json:"status"`
	Stage       string `json:"stage"`
	Year        uint   `json:"year"`
}
