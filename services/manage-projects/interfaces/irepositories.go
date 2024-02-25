package interfaces

import (
	entites "mvp-2-spms/domain-aggregate"
)

// transfers data in domain entities
type IProjetRepository interface {
	GetProfessorProjects(profId string) []entites.Project
	// возвращать вообще все здесь??? а что делать если там нет чего-то в дб? как понять?
	// писать что будет возвращено в структуре
	// но тогда будет неявное раскрытие деталей реализации
	// ====> будем переделывать domain походу
	// потому что возвращать всю инфу (которой может быть дофига) очень затратно
	// т.е. сущность проекта не будет содержать список тасок
	// таски проекта будут получаться через обращение к бдшке
	// наверно так изначально предполагается
}

// transfers data in domain entities
type IStudentRepository interface {
	GetStudentById(studId string) entites.Student
}
