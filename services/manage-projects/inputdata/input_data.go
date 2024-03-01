package inputdata

import (
	"fmt"
	entities "mvp-2-spms/domain-aggregate"
	"time"
)

type GetPfofessorProjects struct {
	ProfessorId uint
}

type GetProjectCommits struct {
	ProfessorId uint
	ProjectId   uint
	From        time.Time
}

type GetProjectById struct {
	ProfessorId uint
	ProjectId   uint
}

type AddProject struct {
	ProfessorId uint
	Theme       string
	StudentId   uint
	Year        uint
}

func (as AddProject) MapToProjectEntity() entities.Project {
	return entities.Project{
		Theme:        as.Theme,
		SupervisorId: fmt.Sprint(as.ProfessorId),
		StudentId:    fmt.Sprint(as.StudentId),
		Year:         as.Year,
		Stage:        entities.ProjectStage(entities.Analysis),
		Status:       entities.ProjectStatus(entities.ProjectInProgress),
	}
}
