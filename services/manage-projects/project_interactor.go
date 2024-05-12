package manageprojects

import (
	"fmt"
	"mvp-2-spms/services/interfaces"
	"mvp-2-spms/services/manage-projects/inputdata"
	"mvp-2-spms/services/manage-projects/outputdata"
	"mvp-2-spms/services/models"

	"golang.org/x/oauth2"
)

type ProjectInteractor struct {
	projectRepo interfaces.IProjetRepository
	studentRepo interfaces.IStudentRepository
	uniRepo     interfaces.IUniversityRepository
	accountRepo interfaces.IAccountRepository
}

func InitProjectInteractor(projRepo interfaces.IProjetRepository, stRepo interfaces.IStudentRepository,
	uniRepo interfaces.IUniversityRepository, accRepo interfaces.IAccountRepository) *ProjectInteractor {
	return &ProjectInteractor{
		projectRepo: projRepo,
		studentRepo: stRepo,
		uniRepo:     uniRepo,
		accountRepo: accRepo,
	}
}

// returns all professor projects (basic information)
func (p *ProjectInteractor) GetProfessorProjects(input inputdata.GetProfessorProjects) (outputdata.GetProfessorProjects, error) {
	// get from database
	projEntities := []outputdata.GetProfessorProjectsEntities{}
	projects, _ := p.projectRepo.GetProfessorProjects(fmt.Sprint(input.ProfessorId))
	for _, project := range projects {
		student, _ := p.studentRepo.GetStudentById(project.StudentId)
		projEntities = append(projEntities, outputdata.GetProfessorProjectsEntities{
			Project: project,
			Student: student,
		})
	}
	output := outputdata.MapToGetProfessorProjects(projEntities)
	return output, nil
}

// returns all commits from all branches from specific time
func (p *ProjectInteractor) GetProjectCommits(input inputdata.GetProjectCommits, gitRepositoryHub interfaces.IGitRepositoryHub) (outputdata.GetProjectCommits, error) {
	// get project repo id
	repo, _ := p.projectRepo.GetProjectRepository(fmt.Sprint(input.ProjectId))

	// getting professor repo hub info, should be checked for existance later
	repohubInfo, _ := p.accountRepo.GetAccountRepoHubData(fmt.Sprint(input.ProfessorId))

	//////////////////////////////////////////////////////////////////////////////////////////////////////
	// check for access token first????????????????????????????????????????????
	token := &oauth2.Token{
		RefreshToken: repohubInfo.ApiKey,
	}
	gitRepositoryHub.Authentificate(token)
	// update token
	repohubInfo.ApiKey = token.RefreshToken
	// save new token
	p.accountRepo.UpdateAccountRepoHubIntegration(repohubInfo)
	// get from repo
	commits, _ := gitRepositoryHub.GetRepositoryCommitsFromTime(repo, input.From)
	output := outputdata.MapToGetProjectCommits(commits)
	return output, nil
}

// returns detailed project data (with student data and ed programme)
func (p *ProjectInteractor) GetProjectById(input inputdata.GetProjectById) (outputdata.GetProjectById, error) {
	// get project by id
	project, _ := p.projectRepo.GetProjectById(fmt.Sprint(input.ProjectId))
	cloudFolder, _ := p.projectRepo.GetProjectFolderLink(fmt.Sprint(input.ProjectId))
	// getting student info
	student, _ := p.studentRepo.GetStudentById(project.StudentId)
	edProg, _ := p.uniRepo.GetEducationalProgrammeById(student.EducationalProgrammeId)
	output := outputdata.MapToGetProjectsById(project, student, edProg, cloudFolder)
	return output, nil
}

// returns project statistics
func (p *ProjectInteractor) GetProjectStatsById(input inputdata.GetProjectStatsById) (outputdata.GetProjectStatsById, error) {
	stats := models.ProjectStats{}
	projId := fmt.Sprint(input.ProjectId)

	stats.ProjectGrading, _ = p.projectRepo.GetProjectGradingById(projId)
	stats.MeetingInfo, _ = p.projectRepo.GetProjectMeetingInfoById(projId)
	stats.TasksInfo, _ = p.projectRepo.GetProjectTaskInfoById(projId)
	output := outputdata.MapToGetProjectStatsById(stats)
	return output, nil
}

func (p *ProjectInteractor) AddProject(input inputdata.AddProject, cloudDrive interfaces.ICloudDrive) (outputdata.AddProject, error) {
	// add to db with repository
	proj, _ := p.projectRepo.CreateProjectWithRepository(input.MapToProjectEntity(), input.MapToRepositoryEntity())
	// getting professor drive info, should be checked for existance later
	driveInfo, _ := p.accountRepo.GetAccountDriveData(fmt.Sprint(input.ProfessorId))

	//////////////////////////////////////////////////////////////////////////////////////////////////////
	// check for access token first????????????????????????????????????????????
	token := &oauth2.Token{
		RefreshToken: driveInfo.ApiKey,
	}
	cloudDrive.Authentificate(token)

	// add folder to cloud
	driveProject, _ := cloudDrive.AddProjectFolder(proj.Project, driveInfo)
	// add folder id from drive
	p.projectRepo.AssignDriveFolder(driveProject)
	// returning id
	output := outputdata.MapToAddProject(proj.Project)
	return output, nil
}
