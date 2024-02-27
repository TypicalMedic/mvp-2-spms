package universityrepository

import (
	"mvp-2-spms/database"
	"mvp-2-spms/database/models"
	entities "mvp-2-spms/domain-aggregate"
)

type UniversityRepository struct {
	dbContext database.Database
}

func InitUniversityRepository(dbcxt database.Database) *UniversityRepository {
	return &UniversityRepository{
		dbContext: dbcxt,
	}
}

func (u *UniversityRepository) GetEducationalProgrammeById(epId string) entities.EducationalProgramme {
	var edProg models.EducationalProgramme
	u.dbContext.DB.Select("*").Where("id = ?", epId).Find(&edProg)
	return edProg.MapToEntity()
}
