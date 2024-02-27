package models

import (
	"fmt"
	entities "mvp-2-spms/domain-aggregate"
	"time"
)

type Student struct {
	Id                     uint   `gorm:"column:id"`
	Name                   string `gorm:"column:name"`
	Surname                string `gorm:"column:surname"`
	Middlename             string `gorm:"column:middlename"`
	EnrollmentYear         uint   `gorm:"column:enrollment_year"`
	EducationalProgrammeId uint   `gorm:"column:educational_programme_id"`
}

func (Student) TableName() string {
	return "student"
}

func (s Student) MapToEntity() entities.Student {
	return entities.Student{
		Person: entities.Person{
			Id:         fmt.Sprint(s.Id),
			Name:       s.Name,
			Surname:    s.Surname,
			Middlename: s.Middlename,
		},
		EducationalProgrammeId: fmt.Sprint(s.EducationalProgrammeId),
		Cource:                 s.GetCource(),
		//EnrollmentYear: s.EnrollmentYear,
	}
}

func (s *Student) GetCource() uint {
	currentDate := time.Now()
	if currentDate.Month() > 9 {
		return uint(currentDate.Year()) - s.EnrollmentYear + 1
	}
	return uint(currentDate.Year()) - s.EnrollmentYear
}
