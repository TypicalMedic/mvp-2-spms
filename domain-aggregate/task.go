package domainaggregate

import (
	"time"
)

type Task struct {
	Id          string
	ProjectId   string
	Name        string
	Description string
	Deadline    time.Time
	cloud       taskOnCloud
}

// in DDD this should be gotten through the repository (ala GetTaskOnCloud(...))
func (t *Task) Cloud() taskOnCloud {
	return t.cloud
}

type taskOnCloud struct {
	FolderId     string
	FolderName   string
	TaskFileId   string
	TaskFileName string
}
