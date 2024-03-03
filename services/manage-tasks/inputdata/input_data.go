package inputdata

import (
	"fmt"
	entities "mvp-2-spms/domain-aggregate"
	"time"
)

type AddTask struct {
	ProfessorId uint
	ProjectId   uint
	Name        string
	Description string
	Deadline    time.Time
}

func (at AddTask) MapToTaskEntity() entities.Task {
	return entities.Task{
		ProjectId:   fmt.Sprint(at.ProjectId),
		Name:        at.Name,
		Description: at.Description,
		Deadline:    at.Deadline,
	}
}
