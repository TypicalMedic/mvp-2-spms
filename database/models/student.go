package models

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
