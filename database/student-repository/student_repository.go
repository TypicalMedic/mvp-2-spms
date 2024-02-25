package studentrepository

import (
	"mvp-2-spms/database"
	"mvp-2-spms/database/models"
	entites "mvp-2-spms/domain-aggregate"
)

type StudentRepository struct {
	dbContext database.Database
}

func InitStudentRepository(dbcxt database.Database) *StudentRepository {
	return &StudentRepository{
		dbContext: dbcxt,
	}
}

func (r *StudentRepository) GetStudentById(studId string) entites.Student {
	var student models.Student
	r.dbContext.DB.Select("*").Where("id = ?", studId).Find(&student)
	result := student.MapToEntity()
	return result
}
