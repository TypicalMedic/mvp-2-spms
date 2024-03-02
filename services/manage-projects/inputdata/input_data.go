package inputdata

import (
	"fmt"
	entities "mvp-2-spms/domain-aggregate"
	"mvp-2-spms/services/models"
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
	ProfessorId         uint
	Theme               string
	StudentId           uint
	Year                uint
	RepositoryOwnerName string
	RepositoryName      string
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

func (as AddProject) MapToRepositoryEntity() models.Repository {
	return models.Repository{
		RepoId:    as.RepositoryName,
		OwnerName: as.RepositoryOwnerName,
	}
}
