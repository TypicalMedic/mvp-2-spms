package domainaggregate

import (
	"fmt"
	"time"
)

type ProjectStatus int

const (
	ProjectNotConfirmed ProjectStatus = iota
	ProjectInProgress
	ProjectFinished
	ProjectCancelled
)

type ProjectStage int

const (
	Analysis ProjectStage = iota
	Design
	Development
	Testing
	Deployment
)

type Project struct {
	Id           string
	Theme        string
	SupervisorId string
	StudentId    string
	Year         uint
	Stage        ProjectStage
	Status       ProjectStatus
	cloud        projectOnCloud
	repo         projectInRepository
}

// in DDD this should be gotten through the repository (ala GetProjectInRepo(...))
func (p *Project) Repo() projectInRepository {
	return p.repo
}

// in DDD this should be gotten through the repository (ala GetProjectOnCloud(...))
func (p *Project) Cloud() projectOnCloud {
	return p.cloud
}

type projectOnCloud struct {
	FolderId   string
	FolderName string
}

type projectInRepository struct {
	RepoId   string
	RepoName string
	Commits  []commit
	Branches []branch
}

type branch struct {
	Name string
	Head commit
}

type commit struct {
	SHA         string
	Name        string
	Description string
	Date        time.Time
	// id repo
}

func (s ProjectStatus) String() string {
	switch s {
	case ProjectStatus(ProjectNotConfirmed):
		return "NotConfirmed"
	case ProjectStatus(ProjectInProgress):
		return "InProgress"
	case ProjectStatus(ProjectFinished):
		return "Finished"
	case ProjectStatus(ProjectCancelled):
		return "Cancelled"
	default:
		return fmt.Sprintf("%d", int(s))
	}
}

func (s ProjectStage) String() string {
	switch s {
	case ProjectStage(Analysis):
		return "Analysis"
	case ProjectStage(Design):
		return "Design"
	case ProjectStage(Development):
		return "Development"
	case ProjectStage(Testing):
		return "Testing"
	case ProjectStage(Deployment):
		return "Deployment"
	default:
		return fmt.Sprintf("%d", int(s))
	}
}
