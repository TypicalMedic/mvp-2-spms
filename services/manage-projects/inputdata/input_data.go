package inputdata

import "time"

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
