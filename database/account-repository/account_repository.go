package accountrepository

import (
	"mvp-2-spms/database"
	"mvp-2-spms/database/models"
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
