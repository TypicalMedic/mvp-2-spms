package outputdata

type GetAccountInfo struct {
	Id            int    `json:"id"`
	Login         string `json:"login"`
	Name          string `json:"name"`
	ScienceDegree string `json:"science_degree"`
	University    string `json:"university"`
}

// func MapToGetAccountInfo(integr models.CloudDriveIntegration) GetAccountInfo {
// 	return GetDriveIntegration{
// 		BaseGetIntegration: BaseGetIntegration{
// 			APIKey: integr.ApiKey,
// 			Type:   integr.Type,
// 		},
// 	}
// }
