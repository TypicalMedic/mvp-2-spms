package managestudents

import (
	"mvp-2-spms/services/interfaces"
	"mvp-2-spms/services/manage-students/inputdata"
	"mvp-2-spms/services/manage-students/outputdata"
)

type StudentInteractor struct {
	studentRepo interfaces.IStudentRepository
	projetRepo  interfaces.IProjetRepository
	uniRepo     interfaces.IUniversityRepository
}

func InitStudentInteractor(stRepo interfaces.IStudentRepository, pjRepo interfaces.IProjetRepository, uRepo interfaces.IUniversityRepository) *StudentInteractor {
	return &StudentInteractor{
		studentRepo: stRepo,
		projetRepo:  pjRepo,
		uniRepo:     uRepo,
	}
}

func (p *StudentInteractor) AddStudent(input inputdata.AddStudent) outputdata.AddStudent {
	// adding student to db, returns created student (with id)
	student, _ := p.studentRepo.CreateStudent(input.MapToStudentEntity())
	// returning id
	output := outputdata.MapToAddStudent(student)
	return output
}

func (p *StudentInteractor) GetStudents(input inputdata.GetStudents) outputdata.GetStudents {
	// get from database
	stEntities := []outputdata.GetStudentsEntities{}
	students, _ := p.studentRepo.GetStudents()
	for _, student := range students {
		project, _ := p.projetRepo.GetStudentCurrentProject(student.Id)
		edProg := p.uniRepo.GetEducationalProgrammeById(student.EducationalProgrammeId)
		stEntities = append(stEntities, outputdata.GetStudentsEntities{
			ProjectTheme:         project.Theme,
			Student:              student,
			EducationalProgramme: edProg.Name,
			PtojectId:            project.Id,
		})
	}
	output := outputdata.MapToGetStudents(stEntities)
	return output
}
