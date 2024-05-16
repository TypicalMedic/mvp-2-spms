package manageprojects

import (
	"errors"
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

	projects, err := p.projectRepo.GetProfessorProjects(fmt.Sprint(input.ProfessorId))
	if err != nil {
		return outputdata.GetProfessorProjects{}, err
	}

	for _, project := range projects {
		student, err := p.studentRepo.GetStudentById(project.StudentId)
		if err != nil {
			return outputdata.GetProfessorProjects{}, err
		}

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
	repo, err := p.projectRepo.GetProjectRepository(fmt.Sprint(input.ProjectId))
	if err != nil {
		return outputdata.GetProjectCommits{}, err
	}

	// getting professor repo hub info, should be checked for existance later
	repohubInfo, err := p.accountRepo.GetAccountRepoHubData(fmt.Sprint(input.ProfessorId))
	if err != nil {
		return outputdata.GetProjectCommits{}, err
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////
	// check for access token first????????????????????????????????????????????
	token := &oauth2.Token{
		RefreshToken: repohubInfo.ApiKey,
	}

	err = gitRepositoryHub.Authentificate(token)
	if err != nil {
		return outputdata.GetProjectCommits{}, err
	}

	// update token
	repohubInfo.ApiKey = token.RefreshToken
	// save new token
	err = p.accountRepo.UpdateAccountRepoHubIntegration(repohubInfo)
	if err != nil {
		return outputdata.GetProjectCommits{}, err
	}

	// get from repo
	commits, err := gitRepositoryHub.GetRepositoryCommitsFromTime(repo, input.From)
	if err != nil {
		return outputdata.GetProjectCommits{}, err
	}

	output := outputdata.MapToGetProjectCommits(commits)
	return output, nil
}

// returns detailed project data (with student data and ed programme)
func (p *ProjectInteractor) GetProjectById(input inputdata.GetProjectById) (outputdata.GetProjectById, error) {
	// get project by id
	project, err := p.projectRepo.GetProjectById(fmt.Sprint(input.ProjectId))
	if err != nil {
		return outputdata.GetProjectById{}, err
	}

	cloudFolder, err := p.projectRepo.GetProjectFolderLink(fmt.Sprint(input.ProjectId))
	if err != nil {
		if !errors.Is(err, models.ErrProjectCloudFolderNotFound) {
			return outputdata.GetProjectById{}, err
		}
		cloudFolder = "" // поменять на нил
	}

	// getting student info
	student, err := p.studentRepo.GetStudentById(project.StudentId)
	if err != nil {
		return outputdata.GetProjectById{}, err
	}

	edProg, err := p.uniRepo.GetEducationalProgrammeById(student.EducationalProgrammeId)
	if err != nil {
		return outputdata.GetProjectById{}, err
	}

	output := outputdata.MapToGetProjectsById(project, student, edProg, cloudFolder)
	return output, nil
}

// returns project statistics
func (p *ProjectInteractor) GetProjectStatsById(input inputdata.GetProjectStatsById) (outputdata.GetProjectStatsById, error) {
	stats := models.ProjectStats{}
	projId := fmt.Sprint(input.ProjectId)

	projectGrading, err := p.projectRepo.GetProjectGradingById(projId)
	if err != nil {
		return outputdata.GetProjectStatsById{}, err
	}
	stats.ProjectGrading = projectGrading

	stats.MeetingInfo, err = p.projectRepo.GetProjectMeetingInfoById(projId)
	if err != nil {
		return outputdata.GetProjectStatsById{}, err
	}

	stats.TasksInfo, err = p.projectRepo.GetProjectTaskInfoById(projId)
	if err != nil {
		return outputdata.GetProjectStatsById{}, err
	}

	output := outputdata.MapToGetProjectStatsById(stats)
	return output, nil
}

func (p *ProjectInteractor) AddProject(input inputdata.AddProject, cloudDrive interfaces.ICloudDrive) (outputdata.AddProject, error) {
	// add to db with repository
	proj, err := p.projectRepo.CreateProjectWithRepository(input.MapToProjectEntity(), input.MapToRepositoryEntity())
	if err != nil {
		return outputdata.AddProject{}, err
	}

	// getting professor drive info, should be checked for existance later
	found := true
	driveInfo, err := p.accountRepo.GetAccountDriveData(fmt.Sprint(input.ProfessorId))
	if err != nil {
		if !errors.Is(err, models.ErrAccountDriveDataNotFound) {
			return outputdata.AddProject{}, err
		}
		found = false
	}

	if found {
		//////////////////////////////////////////////////////////////////////////////////////////////////////
		// check for access token first????????????????????????????????????????????
		token := &oauth2.Token{
			RefreshToken: driveInfo.ApiKey,
		}
		err := cloudDrive.Authentificate(token)
		if err != nil {
			return outputdata.AddProject{}, err
		}

		// add folder to cloud
		driveProject, err := cloudDrive.AddProjectFolder(proj.Project, driveInfo)
		if err != nil {
			return outputdata.AddProject{}, err
		}

		// add folder id from drive
		err = p.projectRepo.AssignDriveFolder(driveProject)
		if err != nil {
			return outputdata.AddProject{}, err
		}
	}

	// returning id
	output := outputdata.MapToAddProject(proj.Project)
	return output, nil
}
