package studentrepository

import (
	"mvp-2-spms/database"
	"mvp-2-spms/database/models"
	entities "mvp-2-spms/domain-aggregate"
)

type StudentRepository struct {
	dbContext database.Database
}

func InitStudentRepository(dbcxt database.Database) *StudentRepository {
	return &StudentRepository{
		dbContext: dbcxt,
	}
}

func (r *StudentRepository) GetStudentById(studId string) (entities.Student, error) {
	var student models.Student
	r.dbContext.DB.Select("*").Where("id = ?", studId).Find(&student)
	result := student.MapToEntity()
	return result, nil
}

func (r *StudentRepository) CreateStudent(student entities.Student) (entities.Student, error) {
	dbstudent := models.Student{}
	dbstudent.MapEntityToThis(student)
	r.dbContext.DB.Create(&dbstudent)
	return dbstudent.MapToEntity(), nil
}

func (r *StudentRepository) GetStudents() ([]entities.Student, error) {
	var students []models.Student
	r.dbContext.DB.Select("*").Find(&students)
	result := []entities.Student{}
	for _, s := range students {
		result = append(result, s.MapToEntity())
	}
	return result, nil
}
