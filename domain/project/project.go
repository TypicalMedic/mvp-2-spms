package project

import (
	"fmt"
	"mvp-2-spms/domain/people"
	"mvp-2-spms/domain/repositoryhub"
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
	Id               uint
	Theme            string
	Supervisor       people.Professor
	Student          people.Student
	Year             uint
	Grade            float32
	Tasks            []Task
	SupervisorReview SupervisorReview
	Repository       repositoryhub.Repository
	Stage            ProjectStage
	Status           ProjectStatus
}

func (p *Project) CalculateGrade() {
	var grade float32 = 0
	for _, gr := range p.SupervisorReview.Criterias {
		grade += gr.Grade * gr.Weight
	}
	p.Grade = grade
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
