package outputdata

type GetPfofessorProjects struct {
	Projects []ProjectData `json:"projects"`
}

type ProjectData struct {
	Id          uint   `json:"id"`
	Theme       string `json:"theme"`
	StudentName string `json:"student_name"`
	Cource      uint   `json:"cource"`
	Status      string `json:"status"`
	Stage       string `json:"stage"`
	Year        uint   `json:"year"`
}
