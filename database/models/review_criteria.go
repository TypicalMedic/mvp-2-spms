package models

import (
	entities "mvp-2-spms/domain-aggregate"
)

type ReviewCriteria struct {
	Id                 uint    `gorm:"column:id"`
	Description        string  `gorm:"column:description"`
	Grade              float32 `gorm:"column:grade"`
	GradeWeight        float32 `gorm:"column:grade_weight"`
	SupervieorReviewId uint    `gorm:"column:supervisor_review_id"`
}

func (ReviewCriteria) TableName() string {
	return "review_criteria"
}

func (r ReviewCriteria) MapToEntity() entities.Criteria {
	return entities.Criteria{
		Description: r.Description,
		Grade:       r.Grade,
		Weight:      r.GradeWeight,
	}
}
