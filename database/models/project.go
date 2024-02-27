package models

import (
	"fmt"
	entities "mvp-2-spms/domain-aggregate"
)

type Project struct {
	Id                 uint    `gorm:"column:id"`
	Theme              string  `gorm:"column:theme"`
	Year               uint    `gorm:"column:year"`
	SupervisorId       uint    `gorm:"column:supervisor_id"`
	StudentId          uint    `gorm:"column:student_id"`
	DefenceGrade       float32 `gorm:"column:defence_grade"`
	FinalGrade         float32 `gorm:"column:final_grade"`
	SupervisorReviewId uint    `gorm:"column:supervisor_review_id"`
	RepoId             uint    `gorm:"column:repo_id"`
	CloudId            uint    `gorm:"column:cloud_id"`
	StageId            uint    `gorm:"column:stage_id"`
	StatusId           uint    `gorm:"column:status_id"`
}

func (Project) TableName() string {
	return "project"
}

func (pj Project) MapToEntity() entities.Project {
	return entities.Project{
		Id:           fmt.Sprint(pj.Id),
		Theme:        pj.Theme,
		SupervisorId: fmt.Sprint(pj.SupervisorId),
		StudentId:    fmt.Sprint(pj.StudentId),
		Year:         pj.Year,
		Stage:        entities.ProjectStage(pj.StageId),
		Status:       entities.ProjectStatus(pj.StatusId),
	}
}
