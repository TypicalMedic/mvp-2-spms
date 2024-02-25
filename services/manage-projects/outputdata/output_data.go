package outputdata

import (
	entites "mvp-2-spms/domain-aggregate"
	"strconv"
)

type GetProfessorProjects struct {
	Projects []getProfProjProjectData `json:"projects"`
}

func MapToGetProfessorProjects(projectEntities map[*entites.Project]entites.Student) GetProfessorProjects {
	outputProjects := []getProfProjProjectData{}
	for project, student := range projectEntities {
		id, _ := strconv.Atoi(project.Id)
		outputProjects = append(outputProjects,
			getProfProjProjectData{
				Id:          id,
				Theme:       project.Theme,
				Status:      project.Status.String(),
				Stage:       project.Stage.String(),
				Year:        int(project.Year),
				StudentName: student.FullNameToString(),
				Cource:      int(student.GetCource()),
			})
	}
	return GetProfessorProjects{
		Projects: outputProjects,
	}
}

type getProfProjProjectData struct {
	Id          int    `json:"id"`
	Theme       string `json:"theme"`
	StudentName string `json:"student_name"`
	Cource      int    `json:"cource"`
	Status      string `json:"status"`
	Stage       string `json:"stage"`
	Year        int    `json:"year"`
}
