package routes

import (
	"mvp-2-spms/internal"
	"mvp-2-spms/web_server/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Router struct {
	router *chi.Mux
	app    *internal.StudentsProjectsManagementApp
}

func (r *Router) Router() *chi.Mux {
	return r.router
}

func InitRouter(app *internal.StudentsProjectsManagementApp) Router {
	r := chi.NewRouter()
	return Router{
		router: r,
		app:    app,
	}
}

func SetupRouter(app *internal.StudentsProjectsManagementApp) Router {
	r := InitRouter(app)
	r.SetupMiddleware()
	r.SetupRoutes()
	return r
}

// middleware for all routes
func (r *Router) SetupMiddleware() {
	r.router.Use(middleware.Logger)
	r.router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Professor-Id"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
}

func (r *Router) SetupRoutes() {
	r.router.Get("/", handlers.Ping)
	r.setupMeetingRoutes()
	r.setupProjectRoutes()
	r.setupStudentRoutes()
	r.setupUniversityRoutes()
	r.setupAccountRoutes()
	r.setupAuthentificationRoutes()
}

func (r *Router) setupProjectRoutes() {
	projH := handlers.InitProjectHandler(r.app.Intercators.ProjectManager, r.app.Intercators.AccountManager, r.app.Integrations.CloudDrives, r.app.Integrations.GitRepositoryHubs)
	taskH := handlers.InitTaskHandler(r.app.Intercators.TaskManager, r.app.Intercators.AccountManager, r.app.Integrations.CloudDrives)

	// setup middleware for checking if professor is authorized and it's his projects?
	r.router.Route("/projects", func(r chi.Router) {
		r.With().Get("/", projH.GetAllProfProjects) // GET /projects with middleware (currently empty)
		r.Get("/filter", dummyHandler)              // GET /projects/filter?student_id=1 query params are accessed with r.URL.Query().Get("student_id")
		r.Post("/add", projH.AddProject)            // POST /projects/add
		// Subrouters:
		r.Route("/{projectID}", func(r chi.Router) {
			// r.Use(///) --> context (for handling not found errors for example)
			r.Get("/", projH.GetProject)                     // GET /projects/123
			r.Get("/statistics", projH.GetProjectStatistics) // GET /projects/123/statistics
			r.Put("/", dummyHandler)                         // PUT /projects/123
			r.Delete("/", dummyHandler)                      // DELETE /projects/123
			r.Get("/commits", projH.GetProjectCommits)       // GET /projects/123/commits?from=2006-01-02T15:04:05.000Z
			r.Route("/tasks", func(r chi.Router) {
				r.Get("/", taskH.GetAllProjectTasks) // GET /projects/123/tasks
				r.Post("/add", taskH.AddTask)        // POST /projects/123/tasks/add
			})
		})
	})
}

func (r *Router) setupStudentRoutes() {
	studH := handlers.InitStudentHandler(r.app.Intercators.StudentManager)

	// setup middleware for checking if professor is authorized and it's his projects?
	r.router.Route("/students", func(r chi.Router) {
		r.With().Get("/", studH.GetStudents) // GET /students with middleware (currently empty)
		r.Post("/add", studH.AddStudent)     // POST /students/add
	})
}
func (r *Router) setupAccountRoutes() {
	accH := handlers.InitAccountHandler(r.app.Intercators.AccountManager)

	// setup middleware for checking if professor is authorized and it's his projects?
	r.router.Route("/accounts", func(r chi.Router) {
		r.Route("/{accID}", func(r chi.Router) {
			r.Get("/", accH.GetAccountInfo)                     // GET /accounts/123
			r.Get("/integrations", accH.GetAccountIntegrations) // GET /accounts/123/integrations
		})
	})
}

func (r *Router) setupUniversityRoutes() {
	uniH := handlers.InitUniversityHandler(r.app.Intercators.UnversityManager)

	// setup middleware for checking if professor is authorized and it's his projects?
	r.router.Route("/universities", func(r chi.Router) {
		r.With().Get("/", dummyHandler) // GET /universities with middleware (currently empty)
		r.Route("/{uniID}", func(r chi.Router) {
			r.Get("/", dummyHandler) // GET /universities/123
			r.Route("/edprogrammes", func(r chi.Router) {
				r.Get("/", uniH.GetAllUniEdProgrammes) // GET /universities/123/edprogrammes
				r.Post("/add", dummyHandler)           // POST /universities/123/edprogrammes/add
			})
		})
	})
}

func (r *Router) setupAuthentificationRoutes() {
	calendarH := handlers.InitPlannerIntegrationHandler(r.app.Integrations.Planners, r.app.Intercators.AccountManager)
	driveH := handlers.InitCloudDriveHandler(r.app.Integrations.CloudDrives, r.app.Intercators.AccountManager)
	repoHubH := handlers.InitGitRepoHandler(r.app.Integrations.GitRepositoryHubs, r.app.Intercators.AccountManager)

	r.router.Route("/auth", func(r chi.Router) {
		r.Route("/integration", func(r chi.Router) {
			r.Route("/authlink", func(r chi.Router) {
				r.Get("/googlecalendar", calendarH.GetGoogleCalendarLink) // GET /auth/integration/authlink/googlecalendar
				r.Get("/googledrive", driveH.GetGoogleDriveLink)          // GET /auth/integration/authlink/googledrive
				r.Get("/github", repoHubH.GetGitHubLink)                  // GET /auth/integration/authlink/github
			})
			r.Route("/access", func(r chi.Router) {
				r.Get("/googlecalendar", calendarH.OAuthCallbackGoogleCalendar) // GET /auth/integration/access/googlecalendar
				r.Get("/googledrive", driveH.OAuthCallbackGoogleDrive)          // GET /auth/integration/access/googledrive
				r.Get("/github", repoHubH.OAuthCallbackGitHub)                  // GET /auth/integration/access/github
			})
		})
	})
}

func (r *Router) setupMeetingRoutes() {
	meetH := handlers.InitMeetingHandler(r.app.Intercators.MeetingManager, r.app.Intercators.AccountManager, r.app.Integrations.Planners)
	// RESTy routes for "meetings" resource
	// setup middleware for checking professor?
	r.router.Route("/meetings", func(r chi.Router) {
		r.With().Get("/", meetH.GetProfessorMeetings) // GET /meetings?from=2006-01-02T15:04:05.000Z with middleware (currently empty)
		r.Get("/filter", dummyHandler)                // GET /meetings/filter?student_id=1&status=planned query params are accessed with r.URL.Query().Get("student_id")
		r.Post("/add", meetH.AddMeeting)              // POST /meetings/add
		// Subrouters:
		r.Route("/{meetingID}", func(r chi.Router) {
			// r.Use(///) --> context (for handling not found errors for example)
			r.Get("/", dummyHandler)    // GET /meetings/123
			r.Put("/", dummyHandler)    // PUT /meetings/123
			r.Delete("/", dummyHandler) // DELETE /meetings/123
		})
	})
}

func dummyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
