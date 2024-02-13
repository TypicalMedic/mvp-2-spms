package internal

import (
	mngInterfaces "mvp-2-spms/services/manage-projects/interfaces"
	hInterfaces "mvp-2-spms/web_server/handlers/interfaces"
)

type App struct {
	Intercators Intercators
}

type Intercators struct {
	Project hInterfaces.IProjetInteractor
}
type Repositories struct {
	Project mngInterfaces.IProjetRepository
	Student mngInterfaces.IStudentRepository
}
