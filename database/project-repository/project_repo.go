package projectrepository

import (
	"fmt"
	"mvp-2-spms/database"
	"mvp-2-spms/database/models"
	entities "mvp-2-spms/domain-aggregate"
	usecaseModels "mvp-2-spms/services/models"
)

type ProjectRepository struct {
	dbContext database.Database
}

func InitProjectRepository(dbcxt database.Database) *ProjectRepository {
	return &ProjectRepository{
		dbContext: dbcxt,
	}
}

func (r *ProjectRepository) GetProfessorProjects(profId string) []entities.Project {
	var projects []models.Project
	r.dbContext.DB.Select("*").Where("supervisor_id = ?", profId).Find(&projects)
	result := []entities.Project{}
	for _, pj := range projects {
		// вынести в маппер
		result = append(result, pj.MapToEntity())
	}
	return result
}

func (r *ProjectRepository) GetProjectRepository(projId string) usecaseModels.Repository {
	var project models.Project
	r.dbContext.DB.Select("repo_id").Where("id = ?", projId).Find(&project)
	var repo models.Repository
	r.dbContext.DB.Select("*").Where("id = ?", project.RepoId).Find(&repo)
	return repo.MapToEntity()
}

func (r *ProjectRepository) GetProjectById(projId string) entities.Project {
	var project models.Project
	r.dbContext.DB.Select("*").Where("id = ?", projId).Find(&project)
	return project.MapToEntity()
}

func (r *ProjectRepository) CreateProject(project entities.Project) entities.Project {
	dbProject := models.Project{}
	dbProject.MapEntityToThis(project)
	r.dbContext.DB.Create(&dbProject)
	return dbProject.MapToEntity()
}

func (r *ProjectRepository) CreateProjectWithRepository(project entities.Project, repo usecaseModels.Repository) usecaseModels.ProjectInRepository {
	dbRepo := models.Repository{}
	dbRepo.MapModelToThis(repo)
	r.dbContext.DB.Create(&dbRepo)

	dbProject := models.Project{}
	dbProject.MapEntityToThis(project)
	dbProject.RepoId = dbRepo.Id
	r.dbContext.DB.Create(&dbProject)
	return usecaseModels.ProjectInRepository{
		Project: dbProject.MapToEntity(),
	}
}

func (r *ProjectRepository) AssignDriveFolder(project usecaseModels.DriveProject) {
	dbCloudFolder := models.CloudFolder{}
	dbCloudFolder.MapUseCaseModelToThis(project.DriveFolder)
	r.dbContext.DB.Create(&dbCloudFolder)
	r.dbContext.DB.Model(&models.Project{}).Where("id = ?", project.Project.Id).Update("cloud_id", project.DriveFolder.Id)
}

func (r *ProjectRepository) GetProjectCloudFolderId(projId string) string {
	proj := models.Project{}
	r.dbContext.DB.Select("cloud_id").Where("id = ?", projId).Find(&proj)
	return fmt.Sprint(proj.CloudId)
}

func (r *ProjectRepository) GetProjectFolderLink(projId string) string {
	result := models.CloudFolder{}
	r.dbContext.DB.Select("link").Where("id = ?", r.GetProjectCloudFolderId(projId)).Find(&result)
	return result.Link
}

func (r *ProjectRepository) GetStudentCurrentProjectTheme(studId string) string {
	proj := models.Project{}
	r.dbContext.DB.Select("theme").Where("student_id = ? AND status_id IN(?, ?)",
		studId, entities.ProjectInProgress,
		entities.ProjectNotConfirmed).Order("year desc").Limit(1).Find(&proj)
	return proj.Theme
}
