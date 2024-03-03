package routes

import (
	"mvp-2-spms/internal"
	"mvp-2-spms/web_server/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
}

func (r *Router) SetupRoutes() {
	r.router.Get("/", handlers.Ping)
	r.setupMeetingRoutes()
	r.setupProjectRoutes()
	r.setupStudentRoutes()
}

func (r *Router) setupProjectRoutes() {
	projH := handlers.InitProjectHandler(r.app.Intercators.ProjectManager)
	taskH := handlers.InitTaskHandler(r.app.Intercators.TaskManager)

	// setup middleware for checking if professor is authorized and it's his projects?
	r.router.Route("/projects", func(r chi.Router) {
		r.With().Get("/", projH.GetAllProfProjects) // GET /projects with middleware (currently empty)
		r.Get("/filter", dummyHandler)              // GET /projects/filter?student_id=1 query params are accessed with r.URL.Query().Get("student_id")
		r.Post("/add", projH.AddProject)            // POST /projects/add
		// Subrouters:
		r.Route("/{projectID}", func(r chi.Router) {
			// r.Use(///) --> context (for handling not found errors for example)
			r.Get("/", projH.GetProject)               // GET /projects/123
			r.Put("/", dummyHandler)                   // PUT /projects/123
			r.Delete("/", dummyHandler)                // DELETE /projects/123
			r.Get("/commits", projH.GetProjectCommits) // GET /projects/123/commits
			r.Route("/tasks", func(r chi.Router) {
				r.Get("/", dummyHandler)      // GET /projects/123/tasks
				r.Post("/add", taskH.AddTask) // POST /projects/123/tasks/add
			})
		})
	})
}
func (r *Router) setupStudentRoutes() {
	studH := handlers.InitStudentHandler(r.app.Intercators.StudentManager)

	// setup middleware for checking if professor is authorized and it's his projects?
	r.router.Route("/students", func(r chi.Router) {
		r.Post("/add", studH.AddStudent) // POST /students/add
	})
}

func (r *Router) setupMeetingRoutes() {
	meetH := handlers.InitMeetingHandler(r.app.Intercators.MeetingManager)
	// RESTy routes for "meetings" resource
	// setup middleware for checking professor?
	r.router.Route("/meetings", func(r chi.Router) {
		r.With().Get("/", dummyHandler)  // GET /meetings with middleware (currently empty)
		r.Get("/filter", dummyHandler)   // GET /meetings/filter?student_id=1 query params are accessed with r.URL.Query().Get("student_id")
		r.Post("/add", meetH.AddMeeting) // POST /meetings/add
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
