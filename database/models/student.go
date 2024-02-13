package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name           string `gorm:"column:name"`
	Surname        string `gorm:"column:surname"`
	Middlename     string `gorm:"column:middlename"`
	EnrollmentYear uint   `gorm:"column:enrollment_year"`
}
