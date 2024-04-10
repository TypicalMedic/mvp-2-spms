package models

import (
	"fmt"
	entities "mvp-2-spms/domain-aggregate"
	"strconv"
)

type Professor struct {
	Id            uint   `gorm:"column:id"`
	Name          string `gorm:"column:name"`
	Surname       string `gorm:"column:surname"`
	Middlename    string `gorm:"column:middlename"`
	ScienceDegree string `gorm:"column:science_degree"`
	UniversityId  uint   `gorm:"column:university_id"`
}

func (*Professor) TableName() string {
	return "professor"
}

func (s *Professor) MapToEntity() entities.Professor {
	return entities.Professor{
		Person: entities.Person{
			Id:         fmt.Sprint(s.Id),
			Name:       s.Name,
			Surname:    s.Surname,
			Middlename: s.Middlename,
		},
		ScienceDegree: s.ScienceDegree,
		UniversityId:  fmt.Sprint(s.UniversityId),
	}
}

func (s *Professor) MapEntityToThis(entity entities.Professor) {
	pId, _ := strconv.Atoi(entity.Id)
	uId, _ := strconv.Atoi(entity.UniversityId)
	s.Id = uint(pId)
	s.Name = entity.Name
	s.Surname = entity.Surname
	s.Middlename = entity.Middlename
	s.ScienceDegree = entity.ScienceDegree
	s.UniversityId = uint(uId)
}
