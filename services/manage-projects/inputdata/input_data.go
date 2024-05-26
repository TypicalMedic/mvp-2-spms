package inputdata

import (
	"fmt"
	entities "mvp-2-spms/domain-aggregate"
	"mvp-2-spms/services/models"
	"time"
)

type GetProfessorProjects struct {
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

type GetProjectStatsById struct {
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

type UpdateProject struct {
	Id                  int
	ProfessorId         *int
	Theme               *string
	StudentId           *int
	Year                *int
	RepositoryOwnerName *string
	RepositoryName      *string
	Status              *int
	Stage               *int
}

func (as UpdateProject) UpdateProjectEntity(p *entities.Project) error {
	if as.ProfessorId != nil {
		p.SupervisorId = fmt.Sprint(*as.ProfessorId)
	}
	if as.Stage != nil {
		p.Stage = entities.ProjectStage(*as.Stage)
	}
	if as.Status != nil {
		p.Status = entities.ProjectStatus(*as.Status)
	}
	if as.StudentId != nil {
		p.StudentId = fmt.Sprint(*as.StudentId)
	}
	if as.Theme != nil {
		p.Theme = *as.Theme
	}
	if as.Year != nil {
		p.Year = uint(*as.Year)
	}
	return nil
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
