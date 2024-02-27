package requestbodies

type AddStudent struct {
	Name                   string `json:"name"`
	Surname                string `json:"surname"`
	Middlename             string `json:"middlename"`
	Cource                 int    `json:"cource"`
	EducationalProgrammeId int    `json:"education_programme_id"`
}
