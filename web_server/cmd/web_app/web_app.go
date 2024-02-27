package main

import (
	"mvp-2-spms/database"
	projectrepository "mvp-2-spms/database/project-repository"
	studentrepository "mvp-2-spms/database/student-repository"
	"mvp-2-spms/integrations/git-repository-hub/github"
	"mvp-2-spms/internal"
	manageprojects "mvp-2-spms/services/manage-projects"
	"mvp-2-spms/web_server/routes"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/student_project_management?parseTime=true"
	gdb, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db := database.InitDatabade(gdb)
	repos := internal.Repositories{
		Projects: projectrepository.InitProjectRepository(*db),
		Students: studentrepository.InitStudentRepository(*db),
	}

	repoHub := github.InitGithub(github.InitGithubAPI())

	interactors := internal.Intercators{
		ProjectManager: manageprojects.InitProjectInteractor(repos.Projects, repos.Students, &repoHub),
	}
	app := internal.StudentsProjectsManagementApp{
		Intercators: interactors,
	}
	router := routes.SetupRouter(&app)
	http.ListenAndServe(":8080", router.Router())
}
