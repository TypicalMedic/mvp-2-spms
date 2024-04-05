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

func InitAccountInteractor(accRepo interfaces.IAccountRepository) *AccountInteractor {
	return &AccountInteractor{
		accountRepo: accRepo,
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

func (a *AccountInteractor) SetPlannerIntegration(input inputdata.SetPlannerIntegration, planner interfaces.IPlannerService) outputdata.SetPlannerIntegration {
	token := planner.GetToken(input.AuthCode)
	refreshTok := token.RefreshToken
	accessTok := token.AccessToken
	expires := token.Expiry
	a.accountRepo.AddAccountPlannerIntegration(fmt.Sprint(input.AccountId), refreshTok, input.Type)

	return outputdata.SetPlannerIntegration{
		AccessToken: accessTok,
		Expiry:      expires,
	}
}

func (a *AccountInteractor) SetDriveIntegration(input inputdata.SetDriveIntegration, drive interfaces.ICloudDrive) outputdata.SetDriveIntegration {
	token := drive.GetToken(input.AuthCode)
	refreshTok := token.RefreshToken
	accessTok := token.AccessToken
	expires := token.Expiry
	a.accountRepo.AddAccountDriveIntegration(fmt.Sprint(input.AccountId), refreshTok, input.Type)

	return outputdata.SetDriveIntegration{
		AccessToken: accessTok,
		Expiry:      expires,
	}
}

func (a *AccountInteractor) SetRepoHubIntegration(input inputdata.SetRepoHubIntegration, planner interfaces.IGitRepositoryHub) outputdata.SetRepoHubIntegration {
	token := planner.GetToken(input.AuthCode)
	refreshTok := token.RefreshToken
	accessTok := token.AccessToken
	expires := token.Expiry
	a.accountRepo.AddAccountRepoHubIntegration(fmt.Sprint(input.AccountId), refreshTok, input.Type)

	return outputdata.SetRepoHubIntegration{
		AccessToken: accessTok,
		Expiry:      expires,
	}
}
