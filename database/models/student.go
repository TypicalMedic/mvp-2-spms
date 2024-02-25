package models

import (
	"fmt"
	entites "mvp-2-spms/domain-aggregate"
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
func (s Student) MapToEntity() entites.Student {
	return entites.Student{
		Person: entites.Person{
			Id:         fmt.Sprint(s.Id),
			Name:       s.Name,
			Surname:    s.Surname,
			Middlename: s.Middlename,
		},
		//EnrollmentYear: s.EnrollmentYear,
	}
}
