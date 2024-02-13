package models

type Project struct {
	Id                 uint `gorm:"column:id"`
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

func (Project) TableName() string {
	return "project"
}
