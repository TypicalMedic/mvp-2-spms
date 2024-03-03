package interfaces

import (
	"mvp-2-spms/services/manage-tasks/inputdata"
	"mvp-2-spms/services/manage-tasks/outputdata"
)

type ITaskInteractor interface {
	AddTask(input inputdata.AddTask) outputdata.AddTask
}
