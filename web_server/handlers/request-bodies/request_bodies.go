package requestbodies

import "time"

type AddStudent struct {
	Name                   string `json:"name"`
	Surname                string `json:"surname"`
	Middlename             string `json:"middlename"`
	Cource                 int    `json:"cource"`
	EducationalProgrammeId int    `json:"education_programme_id"`
}

type AddMeeting struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	MeetingTime time.Time `json:"meeting_time"`
	StudentId   int       `json:"student_participant_id"`
	IsOnline    bool      `json:"is_online"`
}

type AddProject struct {
	Theme          string `json:"theme"`
	StudentId      int    `json:"student_id"`
	Year           int    `json:"year"`
	RepoOwner      string `json:"repository_owner_login"`
	RepositoryName string `json:"repository_name"`
}
