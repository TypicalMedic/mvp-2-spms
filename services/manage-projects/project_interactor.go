package manageprojects

import (
	"fmt"
	"mvp-2-spms/services/interfaces"
	"mvp-2-spms/services/manage-projects/inputdata"
	"mvp-2-spms/services/manage-projects/outputdata"
)

type ProjectInteractor struct {
	projectRepo interfaces.IProjetRepository
	studentRepo interfaces.IStudentRepository
	uniRepo     interfaces.IUniversityRepository
	repoHub     interfaces.IGitRepositoryHub
	cloudDrive  interfaces.ICloudDrive
	accountRepo interfaces.IAccountRepository
}

func InitProjectInteractor(projRepo interfaces.IProjetRepository, stRepo interfaces.IStudentRepository,
	repo interfaces.IGitRepositoryHub, uniRepo interfaces.IUniversityRepository, cloudDrive interfaces.ICloudDrive, accRepo interfaces.IAccountRepository) *ProjectInteractor {
	return &ProjectInteractor{
		projectRepo: projRepo,
		studentRepo: stRepo,
		repoHub:     repo,
		uniRepo:     uniRepo,
		cloudDrive:  cloudDrive,
		accountRepo: accRepo,
	}
}

// returns all professor projects (basic information)
func (p *ProjectInteractor) GetProfessorProjects(input inputdata.GetPfofessorProjects) outputdata.GetProfessorProjects {
	// get from database
	projEntities := []outputdata.GetProfessorProjectsEntities{}
	projects := p.projectRepo.GetProfessorProjects(fmt.Sprint(input.ProfessorId))
	for _, project := range projects {
		student := p.studentRepo.GetStudentById(project.StudentId)
		projEntities = append(projEntities, outputdata.GetProfessorProjectsEntities{
			Project: project,
			Student: student,
		})
	}
	output := outputdata.MapToGetProfessorProjects(projEntities)
	return output
}

// returns all commits from all branches from specific time
func (p *ProjectInteractor) GetProjectCommits(input inputdata.GetProjectCommits) outputdata.GetProjectCommits {
	// get project repo id
	repo := p.projectRepo.GetProjectRepository(fmt.Sprint(input.ProjectId))
	// get from repo
	commits := p.repoHub.GetRepositoryCommitsFromTime(repo, input.From)
	output := outputdata.MapToGetProjectCommits(commits)
	return output
}

// returns detailed project data (with student data and ed programme)
func (p *ProjectInteractor) GetProjectById(input inputdata.GetProjectById) outputdata.GetProjectById {
	// get project by id
	project := p.projectRepo.GetProjectById(fmt.Sprint(input.ProjectId))
	// getting student info
	student := p.studentRepo.GetStudentById(project.StudentId)
	edProg := p.uniRepo.GetEducationalProgrammeById(student.EducationalProgrammeId)
	output := outputdata.MapToGetProjectsById(project, student, edProg)
	return output
}

func (p *ProjectInteractor) AddProject(input inputdata.AddProject) outputdata.AddProject {
	// add to db with repository
	proj := p.projectRepo.CreateProjectWithRepository(input.MapToProjectEntity(), input.MapToRepositoryEntity())
	// getting professor drive info, should be checked for existance later
	driveInfo := p.accountRepo.GetAccountDriveData(fmt.Sprint(input.ProfessorId))
	// add folder to cloud
	driveProject := p.cloudDrive.AddProjectFolder(proj.Project, driveInfo)
	// add folder id from drive
	p.projectRepo.AssignDriveFolder(driveProject)
	// returning id
	output := outputdata.MapToAddProject(proj.Project)
	return output
}
