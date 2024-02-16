package people

import "mvp-2-spms/domain/university"

type Professor struct {
	Person
	ScienceDegree   string
	AvailablePlaces []AvailablePlace
}

type AvailablePlace struct {
	EducationalProgramme university.EducationalProgramme
	Cource               uint
	PlacesCount          uint
}
