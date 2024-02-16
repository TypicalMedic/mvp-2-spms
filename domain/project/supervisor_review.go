package project

import "time"

type SupervisorReview struct {
	Id           uint
	CreationDate time.Time
	Criterias    []Criteria
}
