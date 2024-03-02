package main

import (
	"mvp-2-spms/database"
	accountepository "mvp-2-spms/database/account-repository"
	meetingrepository "mvp-2-spms/database/meeting-repository"
	projectrepository "mvp-2-spms/database/project-repository"
	studentrepository "mvp-2-spms/database/student-repository"
	unirepository "mvp-2-spms/database/university-repository"
	googleDrive "mvp-2-spms/integrations/cloud-drive/google-drive"
	"mvp-2-spms/integrations/git-repository-hub/github"
	googleapi "mvp-2-spms/integrations/google-api"
	googleCalendar "mvp-2-spms/integrations/planner-service/google-calendar"
	"mvp-2-spms/internal"
	managemeetings "mvp-2-spms/services/manage-meetings"
	manageprojects "mvp-2-spms/services/manage-projects"
	managestudents "mvp-2-spms/services/manage-students"
	"mvp-2-spms/web_server/routes"
	"net/http"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/drive/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/student_project_management?parseTime=true"
	gdb, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	db := database.InitDatabade(gdb)
	repos := internal.Repositories{
		Projects:     projectrepository.InitProjectRepository(*db),
		Students:     studentrepository.InitStudentRepository(*db),
		Universities: unirepository.InitUniversityRepository(*db),
		Meetings:     meetingrepository.InitMeetingRepository(*db),
		Accounts:     accountepository.InitAccountRepository(*db),
	}

	repoHub := github.InitGithub(github.InitGithubAPI())

	scopes := []string{calendar.CalendarScope, drive.DriveScope}

	gCalendarApi := googleCalendar.InitCalendarApi(googleapi.InitGoogleAPI(scopes...))
	gCalendar := googleCalendar.InitGoogleCalendar(gCalendarApi)
	gDriveApi := googleDrive.InitDriveApi(googleapi.InitGoogleAPI(scopes...))
	gDrive := googleDrive.InitGoogleDrive(gDriveApi)

	interactors := internal.Intercators{
		ProjectManager: manageprojects.InitProjectInteractor(repos.Projects, repos.Students, &repoHub, repos.Universities, &gDrive, repos.Accounts),
		StudentManager: managestudents.InitStudentInteractor(repos.Students),
		MeetingManager: managemeetings.InitMeetingInteractor(repos.Meetings, &gCalendar, repos.Accounts),
	}
	app := internal.StudentsProjectsManagementApp{
		Intercators: interactors,
	}
	router := routes.SetupRouter(&app)
	http.ListenAndServe(":8080", router.Router())
}
