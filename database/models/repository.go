package models

import (
	entities "mvp-2-spms/domain-aggregate"
)

type Repository struct {
	Id        uint   `gorm:"column:id"`
	Name      string `gorm:"column:name"`
	OwnerName string `gorm:"column:owner_name"`
	IsPublic  bool   `gorm:"column:is_public"`
}

func (Repository) TableName() string {
	return "repository"
}

func (r Repository) MapToEntity() entities.ProjectInRepository {
	return entities.ProjectInRepository{
		RepoId:    r.Name,
		OwnerName: r.OwnerName,
	}
}
