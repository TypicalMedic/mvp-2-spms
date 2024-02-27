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
