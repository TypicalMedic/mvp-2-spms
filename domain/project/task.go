package project

import (
	"mvp-2-spms/domain/drive"
	"time"
)

type Task struct {
	Name        string
	Description string
	Deadline    time.Time
	TaskFolder  drive.Folder
	TaskFile    drive.TextFile
}
