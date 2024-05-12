package models

import "errors"

var (
	ErrAccountNotFound            = errors.New("account not found")
	ErrProfessorNotFound          = errors.New("professor not found")
	ErrAccountPlannerDataNotFound = errors.New("account planner integration not found")
	ErrAccountDriveDataNotFound   = errors.New("account drive integration not found")
	ErrAccountRepoHubDataNotFound = errors.New("account repo hub integration not found")

	ErrMeetingNotFound          = errors.New("meeting not found")
	ErrMeetingPlannerIdNotFound = errors.New("meeting planner id not found")

	ErrProjectNotFound                = errors.New("project not found")
	ErrProjectCloudFolderNotFound     = errors.New("project cloud folder not found")
	ErrProjectCloudFolderLinkNotFound = errors.New("project cloud folder link not found")
	ErrProjectRepoNotFound            = errors.New("project repository not found")
	ErrStudentHasNoCurrentProject     = errors.New("student has no current object")

	ErrStudentNotFound = errors.New("student not found")

	ErrTaskNotFound = errors.New("task not found")

	ErrEdProgrammmeNotFound = errors.New("educational programme not found")

	ErrUniNoFound = errors.New("university not found")

	Err = errors.New("")
)
