package internal

import (
	mngInterfaces "mvp-2-spms/services/interfaces"
	hInterfaces "mvp-2-spms/web_server/handlers/interfaces"
)

type StudentsProjectsManagementApp struct {
	Intercators Intercators
}

type Intercators struct {
	ProjectManager hInterfaces.IProjetInteractor
	StudentManager hInterfaces.IStudentInteractor
	MeetingManager hInterfaces.IMeetingInteractor
}
type Repositories struct {
	Projects     mngInterfaces.IProjetRepository
	Students     mngInterfaces.IStudentRepository
	Universities mngInterfaces.IUniversityRepository
}
