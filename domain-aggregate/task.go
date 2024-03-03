package domainaggregate

import (
	"time"
)

type Task struct {
	Id          string
	ProjectId   string
	Name        string
	Description string
	Deadline    time.Time
}
