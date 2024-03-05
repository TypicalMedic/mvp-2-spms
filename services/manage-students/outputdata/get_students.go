package outputdata

import (
	entities "mvp-2-spms/domain-aggregate"
	"strconv"
)

type GetStudents struct {
	Projects []getStudentsData `json:"students"`
}

func MapToGetStudents(studentEntities []GetStudentsEntities) GetStudents {
	outputStudents := []getStudentsData{}
	for _, studenttEntitiy := range studentEntities {
		id, _ := strconv.Atoi(studenttEntitiy.Student.Id)
		outputStudents = append(outputStudents,
			getStudentsData{
				Id:                   id,
				Name:                 studenttEntitiy.Student.Name,
				Surname:              studenttEntitiy.Student.Surname,
				Middlename:           studenttEntitiy.Student.Middlename,
				EducationalProgramme: studenttEntitiy.EducationalProgramme,
				Cource:               int(studenttEntitiy.Student.Cource),
				PtojectTheme:         studenttEntitiy.ProjectTheme,
			})
	}
	return GetStudents{
		Projects: outputStudents,
	}
}

type GetStudentsEntities struct {
	ProjectTheme         string
	Student              entities.Student
	EducationalProgramme string
}

type getStudentsData struct {
	Id                   int `json:"id"`
	Name                 string
	Surname              string
	Middlename           string
	EducationalProgramme string
	Cource               int
	PtojectTheme         string
}
