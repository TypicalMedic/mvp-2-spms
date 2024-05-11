package accountrepository

import (
	"mvp-2-spms/database"
	"mvp-2-spms/database/models"
	entities "mvp-2-spms/domain-aggregate"
	usecasemodels "mvp-2-spms/services/models"
)

type AccountRepository struct {
	dbContext database.Database
}

func InitAccountRepository(dbcxt database.Database) *AccountRepository {
	return &AccountRepository{
		dbContext: dbcxt,
	}
}

func (r *AccountRepository) GetAccountByLogin(login string) usecasemodels.Account {
	acc := models.Account{}
	r.dbContext.DB.Select("*").Where("login = ?", login).Find(&acc)
	return acc.MapToUseCaseModel()
}
func (r *AccountRepository) AddProfessor(prof entities.Professor) entities.Professor {
	dbProf := models.Professor{}
	dbProf.MapEntityToThis(prof)
	r.dbContext.DB.Create(&dbProf)
	return dbProf.MapToEntity()
}

func (r *AccountRepository) AddAccount(account usecasemodels.Account) {
	dbAcc := models.Account{}
	dbAcc.MapUseCaseModelToThis(account)
	r.dbContext.DB.Create(&dbAcc)
}

func (r *AccountRepository) GetProfessorById(id string) entities.Professor {
	prof := models.Professor{}
	r.dbContext.DB.Select("*").Where("id = ?", id).Find(&prof)
	return prof.MapToEntity()
}

func (r *AccountRepository) GetAccountPlannerData(id string) usecasemodels.PlannerIntegration {
	dbPlanner := models.PlannerIntegration{}
	r.dbContext.DB.Select("*").Where("account_id = ?", id).Find(&dbPlanner)
	return dbPlanner.MapToUseCaseModel()
}

func (r *AccountRepository) GetAccountDriveData(id string) usecasemodels.CloudDriveIntegration {
	dbDrive := models.DriveIntegration{}
	r.dbContext.DB.Select("*").Where("account_id = ?", id).Find(&dbDrive)
	return dbDrive.MapToUseCaseModel()
}

// can return multiple for 1 account, should consider this
func (r *AccountRepository) GetAccountRepoHubData(id string) usecasemodels.BaseIntegration {
	dbRHub := models.GitRepositoryIntegration{}
	r.dbContext.DB.Select("*").Where("account_id = ?", id).Find(&dbRHub)
	return dbRHub.MapToUseCaseModel()
}

func (r *AccountRepository) AddAccountPlannerIntegration(integr usecasemodels.PlannerIntegration) {
	dbPlanner := models.PlannerIntegration{}
	dbPlanner.MapUseCaseModelToThis(integr)
	r.dbContext.DB.Create(&dbPlanner)
}
func (r *AccountRepository) AddAccountDriveIntegration(integr usecasemodels.CloudDriveIntegration) {
	dbDrive := models.DriveIntegration{}
	dbDrive.MapUseCaseModelToThis(integr)
	r.dbContext.DB.Create(&dbDrive)
}
func (r *AccountRepository) AddAccountRepoHubIntegration(integr usecasemodels.BaseIntegration) {
	dbRepoHub := models.GitRepositoryIntegration{}
	dbRepoHub.MapUseCaseModelToThis(integr)
	r.dbContext.DB.Create(&dbRepoHub)
}

func (r *AccountRepository) UpdateAccountPlannerIntegration(integr usecasemodels.PlannerIntegration) {
	plannerDb := models.PlannerIntegration{}
	plannerDb.MapUseCaseModelToThis(integr)
	r.dbContext.DB.Where("account_id = ?", integr.AccountId).Save(&plannerDb)
}
func (r *AccountRepository) UpdateAccountDriveIntegration(integr usecasemodels.CloudDriveIntegration) {
	r.dbContext.DB.Model(&models.DriveIntegration{}).Where("account_id = ?", integr.AccountId).Update("api_key", integr.ApiKey)
}
func (r *AccountRepository) UpdateAccountRepoHubIntegration(integr usecasemodels.BaseIntegration) {
	r.dbContext.DB.Model(&models.GitRepositoryIntegration{}).Where("account_id = ?", integr.AccountId).Update("api_key", integr.ApiKey)
}
