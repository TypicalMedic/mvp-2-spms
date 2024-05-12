package managestudents

import (
	"bytes"
	"crypto/sha512"
	"fmt"
	entities "mvp-2-spms/domain-aggregate"
	"mvp-2-spms/services/interfaces"
	"mvp-2-spms/services/manage-accounts/inputdata"
	"mvp-2-spms/services/manage-accounts/outputdata"
	"mvp-2-spms/services/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/oauth2"
)

const pbkdf2Iterations int = 4096
const pbkdf2HashSize int = 32

type AccountInteractor struct {
	accountRepo interfaces.IAccountRepository
	uniRepo     interfaces.IUniversityRepository
}

func InitAccountInteractor(accRepo interfaces.IAccountRepository, uniRepo interfaces.IUniversityRepository) *AccountInteractor {
	return &AccountInteractor{
		accountRepo: accRepo,
		uniRepo:     uniRepo,
	}
}

func (a *AccountInteractor) GetAccountProfessorId(login string) string {
	res, _ := a.accountRepo.GetAccountByLogin(login)
	return res.Id
}

func (a *AccountInteractor) GetProfessorInfo(input inputdata.GetProfessorInfo) outputdata.GetProfessorInfo {
	profInfo, _ := a.accountRepo.GetProfessorById(fmt.Sprint(input.AccountId))
	uni, _ := a.uniRepo.GetUniversityById(profInfo.UniversityId)
	// add get account login
	output := outputdata.MapToGetAccountInfo(profInfo, uni)
	return output
}

func (a *AccountInteractor) GetPlannerIntegration(input inputdata.GetPlannerIntegration) outputdata.GetPlannerIntegration {
	planner, _ := a.accountRepo.GetAccountPlannerData(fmt.Sprint(input.AccountId))
	output := outputdata.MapToGetPlannerIntegration(planner)
	return output
}

func (a *AccountInteractor) GetDriveIntegration(input inputdata.GetDriveIntegration) outputdata.GetDriveIntegration {
	drive, _ := a.accountRepo.GetAccountDriveData(fmt.Sprint(input.AccountId))
	output := outputdata.MapToGetDriveIntegration(drive)
	return output
}
func (a *AccountInteractor) SetProfessorPlanner(plannerId, profId string) {
	plannerInfo, _ := a.accountRepo.GetAccountPlannerData(profId)
	plannerInfo.PlannerData.Id = plannerId
	a.accountRepo.UpdateAccountPlannerIntegration(plannerInfo)
}

func (a *AccountInteractor) GetProfessorIntegrPlanners(profId string, planner interfaces.IPlannerService) outputdata.GetProfessorIntegrPlanners {
	plannerInfo, _ := a.accountRepo.GetAccountPlannerData(profId)
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	// check for access token first????????????????????????????????????????????
	token := &oauth2.Token{
		RefreshToken: plannerInfo.ApiKey,
	}
	planner.Authentificate(token)

	planners := planner.GetAllPlanners()
	return outputdata.MapToGetProfessorIntegrPlanners(planners)
}

func (a *AccountInteractor) GetDriveBaseFolderName(folderId, profId string, cloudDrive interfaces.ICloudDrive) string {
	driveInfo, _ := a.accountRepo.GetAccountDriveData(fmt.Sprint(profId))

	//////////////////////////////////////////////////////////////////////////////////////////////////////
	// check for access token first????????????????????????????????????????????
	token := &oauth2.Token{
		RefreshToken: driveInfo.ApiKey,
	}
	cloudDrive.Authentificate(token)
	folderName := cloudDrive.GetFolderNameById(folderId)
	return folderName
}

func (a *AccountInteractor) GetRepoHubIntegration(input inputdata.GetRepoHubIntegration) outputdata.GetRepoHubIntegration {
	repoHub, _ := a.accountRepo.GetAccountRepoHubData(fmt.Sprint(input.AccountId))
	output := outputdata.MapToGetRepoHubIntegration(repoHub)
	return output
}

func (a *AccountInteractor) SetPlannerIntegration(input inputdata.SetPlannerIntegration, planner interfaces.IPlannerService) outputdata.SetPlannerIntegration {
	token := planner.GetToken(input.AuthCode)
	refreshTok := token.RefreshToken
	accessTok := token.AccessToken
	expires := token.Expiry

	integr := models.PlannerIntegration{
		BaseIntegration: models.BaseIntegration{
			AccountId: fmt.Sprint(input.AccountId),
			ApiKey:    refreshTok,
			Type:      input.Type,
		},
		PlannerData: models.PlannerData{},
	}
	a.accountRepo.AddAccountPlannerIntegration(integr)

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

	drive.Authentificate(token)
	baseFolder := drive.AddProfessorBaseFolder()

	integr := models.CloudDriveIntegration{
		BaseIntegration: models.BaseIntegration{
			AccountId: fmt.Sprint(input.AccountId),
			ApiKey:    refreshTok,
			Type:      input.Type,
		},
		DriveData: baseFolder,
	}
	a.accountRepo.AddAccountDriveIntegration(integr)

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
	integr := models.BaseIntegration{
		AccountId: fmt.Sprint(input.AccountId),
		ApiKey:    refreshTok,
		Type:      input.Type,
	}
	a.accountRepo.AddAccountRepoHubIntegration(integr)

	return outputdata.SetRepoHubIntegration{
		AccessToken: accessTok,
		Expiry:      expires,
	}
}

func (a *AccountInteractor) GetAccountIntegrations(input inputdata.GetAccountIntegrations) outputdata.GetAccountIntegrations {
	drive, _ := a.accountRepo.GetAccountDriveData(fmt.Sprint(input.AccountId))
	planner, _ := a.accountRepo.GetAccountPlannerData(fmt.Sprint(input.AccountId))
	repohub, _ := a.accountRepo.GetAccountRepoHubData(fmt.Sprint(input.AccountId))

	var (
		outputDrive   *outputdata.GetAccountIntegrationsDrive
		outputPlanner *outputdata.GetAccountIntegrationsPlanner
		outputRepos   []outputdata.GetAccountIntegrationsIntegr = []outputdata.GetAccountIntegrationsIntegr{}
	)
	if drive.AccountId != "0" {
		outputDrive = &outputdata.GetAccountIntegrationsDrive{
			Type: outputdata.GetAccountIntegrationsIntegr{
				Id:   drive.Type,
				Name: drive.GetTypeAsString(),
			},
			BaseFolderId: drive.BaseFolderId, ///////////////////////////////////////change
		}
	}
	if planner.AccountId != "0" {
		outputPlanner = &outputdata.GetAccountIntegrationsPlanner{
			Type: outputdata.GetAccountIntegrationsIntegr{
				Id:   planner.Type,
				Name: planner.GetTypeAsString(),
			},
			PlannerName: planner.PlannerData.Id, ///////////////////////////////////////change
		}
	}
	if repohub.AccountId != "0" {
		outputRepos = append(outputRepos, outputdata.GetAccountIntegrationsIntegr{
			Id:   repohub.Type,
			Name: repohub.GetRepoHubTypeAsString(),
		})
	}
	return outputdata.MapToGetAccountIntegrations(outputDrive, outputPlanner, outputRepos)
}

func (a *AccountInteractor) CheckCredsValidity(input inputdata.CheckCredsValidity) bool {
	account, _ := a.accountRepo.GetAccountByLogin(input.Login)
	key := pbkdf2.Key([]byte(input.Password), []byte(account.Salt), pbkdf2Iterations, pbkdf2HashSize, sha512.New)

	return bytes.Equal(key, account.Hash)
}
func (a *AccountInteractor) CheckUsernameExists(input inputdata.CheckUsernameExists) bool {
	account, _ := a.accountRepo.GetAccountByLogin(input.Login)
	return account.Login == input.Login
}
func (a *AccountInteractor) SignUp(input inputdata.SignUp) outputdata.SignUp {
	salt := uuid.NewString()
	passHash := pbkdf2.Key([]byte(input.Password), []byte(salt), pbkdf2Iterations, pbkdf2HashSize, sha512.New)

	prof := entities.Professor{
		Person: entities.Person{
			Name:       input.Name,
			Surname:    input.Surname,
			Middlename: input.Middlename,
		},
		ScienceDegree: input.ScienceDegree,
		UniversityId:  fmt.Sprint(input.UniId),
	}
	prof, _ = a.accountRepo.AddProfessor(prof)
	account := models.Account{
		Login: input.Login,
		Hash:  passHash,
		Salt:  salt,
		Id:    prof.Id,
	}
	a.accountRepo.AddAccount(account)
	return outputdata.SignUp{
		Id:    account.Id,
		Login: account.Login,
	}
}
