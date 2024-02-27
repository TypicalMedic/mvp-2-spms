package main

import (
	"mvp-2-spms/database"
	projectrepository "mvp-2-spms/database/project-repository"
	studentrepository "mvp-2-spms/database/student-repository"
	unirepository "mvp-2-spms/database/university-repository"
	"mvp-2-spms/integrations/git-repository-hub/github"
	"mvp-2-spms/internal"
	manageprojects "mvp-2-spms/services/manage-projects"
	managestudents "mvp-2-spms/services/manage-students"
	"mvp-2-spms/web_server/routes"
	"net/http"

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
	}

	repoHub := github.InitGithub(github.InitGithubAPI())

	interactors := internal.Intercators{
		ProjectManager: manageprojects.InitProjectInteractor(repos.Projects, repos.Students, &repoHub, repos.Universities),
		StudentManager: managestudents.InitStudentInteractor(repos.Students),
	}
	app := internal.StudentsProjectsManagementApp{
		Intercators: interactors,
	}
	router := routes.SetupRouter(&app)
	http.ListenAndServe(":8080", router.Router())
}
