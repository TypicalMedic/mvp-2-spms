package project

import "time"

type SupervisorReview struct {
	CreationDate time.Time
	Criterias    []Criteria
}
