package internal

import (
	mngInterfaces "mvp-2-spms/services/interfaces"
	hInterfaces "mvp-2-spms/web_server/handlers/interfaces"
)

type GetRepoHubName int

const (
	GitHub GetRepoHubName = iota
	GitLab
)

type CloudDriveName int

const (
	GoogleDrive CloudDriveName = iota
	YandexDisk
	OneDrive
)

type PlannerName int

const (
	GoogleCalendar PlannerName = iota
	YandexCalendar
)

type GitRepositoryHubs map[GetRepoHubName](mngInterfaces.IGitRepositoryHub)
type CloudDrives map[CloudDriveName](mngInterfaces.ICloudDrive)
type Planners map[PlannerName](mngInterfaces.IPlannerService)

type StudentsProjectsManagementApp struct {
	Intercators  Intercators
	Authorizer   hInterfaces.IAuthorizer
	Integrations Integrations
}

type Intercators struct {
	AccountManager   hInterfaces.IAccountInteractor
	ProjectManager   hInterfaces.IProjetInteractor
	StudentManager   hInterfaces.IStudentInteractor
	MeetingManager   hInterfaces.IMeetingInteractor
	TaskManager      hInterfaces.ITaskInteractor
	UnversityManager hInterfaces.IUniversityInteractor
}

type Integrations struct {
	GitRepositoryHubs
	CloudDrives
	Planners
}

type Repositories struct {
	Projects     mngInterfaces.IProjetRepository
	Students     mngInterfaces.IStudentRepository
	Universities mngInterfaces.IUniversityRepository
	Meetings     mngInterfaces.IMeetingRepository
	Accounts     mngInterfaces.IAccountRepository
	Tasks        mngInterfaces.ITaskRepository
}
