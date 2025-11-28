package http

import "github.com/gorilla/mux"

func newRouter(h *HTTPHandler) *mux.Router {
	r := mux.NewRouter()

	// Users
	r.HandleFunc("/workflow/user", h.CreateUser).Methods("POST")
	r.HandleFunc("/workflow/user", h.GetUser).Methods("GET")

	// New: GET /workflow/users/{id}
	r.HandleFunc("/workflow/users/{id}", h.GetUserByID).Methods("GET")

	// Tasks
	r.HandleFunc("/workflow/task", h.CreateTask).Methods("POST")
	r.HandleFunc("/workflow/task", h.GetTask).Methods("GET")

	// New: GET /workflow/tasks/{id}
	r.HandleFunc("/workflow/tasks/{id}", h.GetTaskByID).Methods("GET")

	// Healthcheck
	r.HandleFunc("/health", h.Health).Methods("GET")

	return r
}
