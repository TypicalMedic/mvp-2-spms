package interfaces

import (
	entities "mvp-2-spms/domain-aggregate"
	"time"
)

type IGitRepositoryHub interface {
	GetRepositoryCommitsFromTime(repo entities.ProjectInRepository, fromTime time.Time) []entities.Commit
}
