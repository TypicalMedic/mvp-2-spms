package interfaces

import (
	"mvp-2-spms/domain/people"
	"mvp-2-spms/domain/project"
)

// transfers data in domain entities
type IProjetRepository interface {
	GetProfessorProjects(profId uint) []project.Project
	// возвращать вообще все здесь??? а что делать если там нет чего-то в дб? как понять?
	// писать что будет возвращено в структуре
}

// transfers data in domain entities
type IStudentRepository interface {
	GetStudentById(studId uint) people.Student
}
