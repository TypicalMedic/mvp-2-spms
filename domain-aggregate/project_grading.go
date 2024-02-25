package domainaggregate

import "time"

const (
	defenceGradeWeight    = 0.6
	supervisorGradeWeight = 1 - defenceGradeWeight
)

type ProjectGrading struct {
	ProjectId        string
	DefenceGrade     float32
	supervisorReview supervisorReview
}

// what if there're no grades yet?
func (pg ProjectGrading) CalculateGrade() float32 {
	return pg.supervisorReview.GetGrade()*supervisorGradeWeight + pg.DefenceGrade*defenceGradeWeight
}

type supervisorReview struct {
	Id           uint
	CreationDate time.Time
	criterias    []Criteria
}

// what if there're no grades yet?
func (s *supervisorReview) GetGrade() float32 {
	var supGrade float32 = 0
	for _, gr := range s.criterias {
		supGrade += gr.Grade * gr.Weight
	}
	return supGrade
}

type Criteria struct {
	Description string
	Grade       float32
	Weight      float32 // from 0 to 1, sum of criterias = 1
}
