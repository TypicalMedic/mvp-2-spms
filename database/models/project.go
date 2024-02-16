package models

import (
	"mvp-2-spms/domain/people"
	projectEntites "mvp-2-spms/domain/project"
	"mvp-2-spms/domain/repositoryhub"
)

type Project struct {
	Id                 uint    `gorm:"column:id"`
	Theme              string  `gorm:"column:theme"`
	Year               uint    `gorm:"column:year"`
	SupervisorId       uint    `gorm:"column:supervisor_id"`
	StudentId          uint    `gorm:"column:student_id"`
	Grade              float32 `gorm:"column:grade"`
	SupervisorReviewId uint    `gorm:"column:supervisor_review_id"`
	RepoId             uint    `gorm:"column:repo_id"`
	StageId            uint    `gorm:"column:stage_id"`
	StatusId           uint    `gorm:"column:status_id"`
}

func (Project) TableName() string {
	return "project"
}

func (pj Project) MapToEntity() projectEntites.Project {
	return projectEntites.Project{
		Id:    pj.Id,
		Theme: pj.Theme,
		Supervisor: people.Professor{
			Person: people.Person{
				Id: pj.SupervisorId,
			},
		},
		Student: people.Student{
			Person: people.Person{
				Id: pj.StudentId,
			},
		},
		Year:  pj.Year,
		Grade: pj.Grade,
		SupervisorReview: projectEntites.SupervisorReview{
			Id: pj.SupervisorReviewId,
		},
		Repository: repositoryhub.Repository{
			Id: pj.RepoId,
		},
		Stage:  projectEntites.ProjectStage(pj.StageId),
		Status: projectEntites.ProjectStatus(pj.StatusId),
	}
}
