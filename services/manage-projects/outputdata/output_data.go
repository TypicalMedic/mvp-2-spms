package outputdata

type GetPfofessorProjects struct {
	Projects []ProjectData
}

type ProjectData struct {
	Id          uint
	Theme       string
	StudentName string
	Cource      uint
	Status      string
	Stage       string
	Year        uint
}
