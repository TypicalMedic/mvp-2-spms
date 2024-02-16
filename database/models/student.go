package models

import (
	"mvp-2-spms/domain/people"
)

type Student struct {
	Id             uint   `gorm:"column:id"`
	Name           string `gorm:"column:name"`
	Surname        string `gorm:"column:surname"`
	Middlename     string `gorm:"column:middlename"`
	EnrollmentYear uint   `gorm:"column:enrollment_year"`
}

func (Student) TableName() string {
	return "student"
}
func (s Student) MapToEntity() people.Student {
	return people.Student{
		Person: people.Person{
			Id:         s.Id,
			Name:       s.Name,
			Surname:    s.Surname,
			Middlename: s.Middlename,
		},
		EnrollmentYear: s.EnrollmentYear,
	}
}
