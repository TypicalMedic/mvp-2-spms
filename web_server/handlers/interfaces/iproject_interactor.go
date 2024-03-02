package interfaces

import (
	"mvp-2-spms/services/manage-projects/inputdata"
	"mvp-2-spms/services/manage-projects/outputdata"
)

type IProjetInteractor interface {
	GetProfessorProjects(input inputdata.GetPfofessorProjects) outputdata.GetProfessorProjects
	GetProjectCommits(input inputdata.GetProjectCommits) outputdata.GetProjectCommits
	GetProjectById(input inputdata.GetProjectById) outputdata.GetProjectById
	AddProject(input inputdata.AddProject) outputdata.AddProject
}
