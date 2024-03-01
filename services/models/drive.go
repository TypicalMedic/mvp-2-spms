package models

import (
	entities "mvp-2-spms/domain-aggregate"
)

type DriveData struct {
	BaseFolderId string
}

type DriveProject struct {
	entities.Project
	ProjectFolderId string
}
