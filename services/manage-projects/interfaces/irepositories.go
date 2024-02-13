package interfaces

import (
	"mvp-2-spms/domain/project"
)

// transfers data in domain entities
type IProjetRepository interface {
	GetProfessorProjects(profId uint) []project.Project // возвращать вообще все здесь??? а что делать если там нет чего-то в дб? как понять?
}
