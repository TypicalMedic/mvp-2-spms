package project

import (
	"mvp-2-spms/domain/people"
	"mvp-2-spms/domain/repositoryhub"
)

type Project struct {
	Theme            string
	Supervisor       people.Professor
	Student          people.Student
	Year             uint
	Grade            float32
	Tasks            []Task
	SupervisorReview SupervisorReview
	Repsitory        repositoryhub.Repository
}

func (p *Project) CalculateGrade() {
	var grade float32 = 0
	for _, gr := range p.SupervisorReview.Criterias {
		grade += gr.Grade * gr.Weight
	}
	p.Grade = grade
}
