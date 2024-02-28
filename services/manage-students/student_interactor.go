package managestudents

import (
	"mvp-2-spms/services/interfaces"
	"mvp-2-spms/services/manage-students/inputdata"
	"mvp-2-spms/services/manage-students/outputdata"
)

type StudentInteractor struct {
	studentRepo interfaces.IStudentRepository
}

func InitStudentInteractor(stRepo interfaces.IStudentRepository) *StudentInteractor {
	return &StudentInteractor{
		studentRepo: stRepo,
	}
}

func (p *StudentInteractor) AddStudent(input inputdata.AddStudent) outputdata.AddStudent {
	// adding student to db, returns created student (with id)
	student := p.studentRepo.CreateStudent(input.MapToStudentEntity())
	// returning id
	output := outputdata.MapToAddStudent(student)
	return output
}
