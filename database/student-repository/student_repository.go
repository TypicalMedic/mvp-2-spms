package studentrepository

import (
	"mvp-2-spms/database"
	"mvp-2-spms/database/models"
	"mvp-2-spms/domain/people"
)

type StudentRepository struct {
	dbContext database.Database
}

func InitStudentRepository(dbcxt database.Database) *StudentRepository {
	return &StudentRepository{
		dbContext: dbcxt,
	}
}

func (r *StudentRepository) GetStudentById(studId uint) people.Student {
	var student models.Student
	r.dbContext.DB.Select("*").Where("id = ?", studId).Find(&student)
	result := people.Student{
		Person: people.Person{
			Id:         studId,
			Name:       student.Name,
			Surname:    student.Surname,
			Middlename: student.Middlename,
		},
		EnrollmentYear: student.EnrollmentYear,
	}
	return result
}