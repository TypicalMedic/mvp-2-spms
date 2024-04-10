package domainaggregate

type Professor struct {
	Person
	ScienceDegree   string
	AvailablePlaces []AvailablePlace // make private?
	UniversityId    string
}

type AvailablePlace struct {
	EducationalProgrammeId string
	Cource                 uint
	PlacesCount            uint
}

// ???
type Account struct {
	Id           string
	Email        string
	Integrations Integrations
}

type Integrations struct {
	// api keys?
}
