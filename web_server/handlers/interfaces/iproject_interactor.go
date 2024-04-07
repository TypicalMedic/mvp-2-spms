package interfaces

import (
	"mvp-2-spms/services/interfaces"
	"mvp-2-spms/services/manage-projects/inputdata"
	"mvp-2-spms/services/manage-projects/outputdata"
)

type IProjetInteractor interface {
	GetProfessorProjects(input inputdata.GetProfessorProjects) outputdata.GetProfessorProjects
	GetProjectCommits(input inputdata.GetProjectCommits, gitRepositoryHub interfaces.IGitRepositoryHub) outputdata.GetProjectCommits
	GetProjectById(input inputdata.GetProjectById) outputdata.GetProjectById
	AddProject(input inputdata.AddProject, cloudDrive interfaces.ICloudDrive) outputdata.AddProject
	GetProjectStatsById(input inputdata.GetProjectStatsById) outputdata.GetProjectStatsById
}
