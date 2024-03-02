package interfaces

import (
	entities "mvp-2-spms/domain-aggregate"
	"mvp-2-spms/services/models"
)

// transfers data in domain entities
type IProjetRepository interface {
	GetProfessorProjects(profId string) []entities.Project
	// возвращать вообще все здесь??? а что делать если там нет чего-то в дб? как понять?
	// писать что будет возвращено в структуре
	// но тогда будет неявное раскрытие деталей реализации
	// ====> будем переделывать domain походу
	// потому что возвращать всю инфу (которой может быть дофига) очень затратно
	// т.е. сущность проекта не будет содержать список тасок
	// таски проекта будут получаться через обращение к бдшке
	// наверно так изначально предполагается
	GetProjectRepository(projId string) models.Repository
	GetProjectById(projId string) entities.Project
	CreateProject(entities.Project) entities.Project
	CreateProjectWithRepository(entities.Project, models.Repository) models.ProjectInRepository
	AssignDriveFolder(models.DriveProject)
}

// transfers data in domain entities
type IStudentRepository interface {
	GetStudentById(studId string) entities.Student
	CreateStudent(entities.Student) entities.Student
}

type IUniversityRepository interface {
	GetEducationalProgrammeById(epId string) entities.EducationalProgramme
}

// transfers data in domain entities
type IMeetingRepository interface {
	CreateMeeting(entities.Meeting) entities.Meeting
	AssignPlannerMeeting(models.PlannerMeeting)
}

type IAccountRepository interface {
	GetAccountPlannerData(id string) models.PlannerIntegration  // returns planner integration for later usage of api key???
	GetAccountDriveData(id string) models.CloudDriveIntegration // returns drive integration for later usage of api key???
}
