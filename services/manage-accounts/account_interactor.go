package managestudents

import (
	"fmt"
	"mvp-2-spms/services/interfaces"
	"mvp-2-spms/services/manage-accounts/inputdata"
	"mvp-2-spms/services/manage-accounts/outputdata"
)

type AccountInteractor struct {
	accountRepo interfaces.IAccountRepository
}

func InitAccountInteractor(stRepo interfaces.IAccountRepository) *AccountInteractor {
	return &AccountInteractor{
		accountRepo: stRepo,
	}
}

func (a *AccountInteractor) GetPlannerIntegration(input inputdata.GetPlannerIntegration) outputdata.GetPlannerIntegration {
	planner := a.accountRepo.GetAccountPlannerData(fmt.Sprint(input.AccountId))
	output := outputdata.MapToGetPlannerIntegration(planner)
	return output
}

func (a *AccountInteractor) GetDriveIntegration(input inputdata.GetDriveIntegration) outputdata.GetDriveIntegration {
	drive := a.accountRepo.GetAccountDriveData(fmt.Sprint(input.AccountId))
	output := outputdata.MapToGetDriveIntegration(drive)
	return output
}

func (a *AccountInteractor) GetRepoHubIntegration(input inputdata.GetRepoHubIntegration) outputdata.GetRepoHubIntegration {
	repoHub := a.accountRepo.GetAccountRepoHubData(fmt.Sprint(input.AccountId))
	output := outputdata.MapToGetRepoHubIntegration(repoHub)
	return output
}
