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

func (u *UniversityRepository) GetUniversityById(uId string) entities.University {
	var uni models.University
	u.dbContext.DB.Select("*").Where("id = ?", uId).Find(&uni)
	return uni.MapToEntity()
}

func (u *UniversityRepository) GetUniversityEducationalProgrammes(uniId string) []entities.EducationalProgramme {
	var edProgs []models.EducationalProgramme

	u.dbContext.DB.Raw(
		`SELECT educational_programme.* 
		FROM (SELECT * 
		FROM department
		WHERE uni_id = ?) as depts
		LEFT JOIN faculty ON faculty.dept_id = depts.id
		LEFT JOIN educational_programme ON educational_programme.faculty_id = faculty.id;`,
		uniId).Scan(&edProgs)

	result := []entities.EducationalProgramme{}
	for _, p := range edProgs {
		result = append(result, p.MapToEntity())
	}
	return result
}
