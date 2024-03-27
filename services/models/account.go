package models

type Account struct {
	Id    string
	Login string
	Hash  string //?
}

type BaseIntegration struct {
	AccountId string
	ApiKey    string
	Type      int
}

type PlannerIntegration struct {
	BaseIntegration
	PlannerData
}

type CloudDriveIntegration struct {
	BaseIntegration
	DriveData
}
