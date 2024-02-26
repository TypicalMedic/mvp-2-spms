package domainaggregate

type Student struct {
	Person
	//EnrollmentYear uint
	EducationalProgrammeId string
	cource                 currentCource
}

// in DDD this should be gotten through the repository (ala GetStudCource(...))?
func (s *Student) Cource() currentCource {
	return s.cource
}

type currentCource struct {
	CourceNo     uint
	ProjectId    string
	SupervisorId Professor
}

func (s *Student) GetCource() uint {
	return s.cource.CourceNo
}
