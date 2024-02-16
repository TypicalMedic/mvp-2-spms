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
}

func (r *Router) setupProjectRoutes() {
	projH := handlers.InitProjectHandler(r.app.Intercators.ProjectManager)

	// setup middleware for checking if professor is authorized and it's his projects?
	r.router.Route("/projects", func(r chi.Router) {
		r.With().Get("/", projH.GetAllProfProjects) // GET /projects with middleware (currently empty)
		r.Get("/filter", dummyHandler)              // GET /projects/filter?student_id=1 query params are accessed with r.URL.Query().Get("student_id")
		r.Post("/add", dummyHandler)                // POST /projects/add
		// Subrouters:
		r.Route("/{projectID}", func(r chi.Router) {
			// r.Use(///) --> context (for handling not found errors for example)
			r.Get("/", dummyHandler)        // GET /projects/123
			r.Put("/", dummyHandler)        // PUT /projects/123
			r.Delete("/", dummyHandler)     // DELETE /projects/123
			r.Get("/commits", dummyHandler) // GET /projects/123/commits
			r.Route("/tasks", func(r chi.Router) {
				r.Get("/", dummyHandler)     // GET /projects/123/tasks
				r.Post("/add", dummyHandler) // POST /projects/123/tasks/add
				r.Put("/", dummyHandler)     // PUT /projects/123/tasks
				r.Delete("/", dummyHandler)  // DELETE /projects/123/tasks
			})
		})
	})
}

func (r *Router) setupMeetingRoutes() {
	// RESTy routes for "meetings" resource
	// setup middleware for checking professor?
	r.router.Route("/meetings", func(r chi.Router) {
		r.With().Get("/", dummyHandler) // GET /meetings with middleware (currently empty)
		r.Get("/filter", dummyHandler)  // GET /meetings/filter?student_id=1 query params are accessed with r.URL.Query().Get("student_id")
		r.Post("/add", dummyHandler)    // POST /meetings/add
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
