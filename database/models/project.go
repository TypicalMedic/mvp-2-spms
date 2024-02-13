package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Theme              string
	Year               uint
	SupervisorId       uint
	StudentId          uint
	Grade              float32
	SupervisorReviewId uint
	RepoId             uint
	StageId            uint
	StatusId           uint
}
