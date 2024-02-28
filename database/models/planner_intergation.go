package models

import (
	"fmt"
	"mvp-2-spms/services/manage-meetings/models"
	"strconv"
)

type PlannerIntegration struct {
	AccountId uint   `gorm:"column:account_id"`
	PlannerId string `gorm:"column:planner_id"`
	ApiKey    string `gorm:"column:api_key"`
}

func (*PlannerIntegration) TableName() string {
	return "planner_integration"
}

func (pi *PlannerIntegration) MapToUseCaseModel() models.PlannerIntegration {
	return models.PlannerIntegration{
		BaseIntegration: models.BaseIntegration{
			AccountId: fmt.Sprint(pi.AccountId),
			ApiKey:    pi.ApiKey,
		},
		PlannerData: models.PlannerData{
			Id: pi.PlannerId,
		},
	}
}

func (pi *PlannerIntegration) MapUseCaseModelToThis(model models.PlannerIntegration) {
	sId, _ := strconv.Atoi(model.AccountId)
	pi.AccountId = uint(sId)
	pi.PlannerId = model.PlannerData.Id
	pi.ApiKey = model.ApiKey
}
