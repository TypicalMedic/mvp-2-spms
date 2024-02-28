package models

type Account struct {
	Id    string
	Login string
	Hash  string //?
}

type PlannerIntegration struct {
	BaseIntegration
	PlannerData
}

type BaseIntegration struct {
	AccountId string
	ApiKey    string
}
