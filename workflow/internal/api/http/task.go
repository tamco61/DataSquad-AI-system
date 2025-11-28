package http

import (
	"encoding/json"
	"net/http"
	"workflow/internal/model"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *HTTPHandler) GetTasksFiltering(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	assignee := r.URL.Query().Get("assignee_id")

	filters := model.TaskFilter{
		Status:     status,
		AssigneeID: assignee,
	}

	tasks, err := h.service.GetTasksFiltering(r.Context(), filters)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "db_error", err.Error())
		return
	}

	if tasks == nil {
		empty := []model.TaskDTO{}
		writeJSON(w, http.StatusOK, empty)
		return
	}

	writeJSON(w, http.StatusOK, tasks)
}

func (h *HTTPHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	var req model.TaskDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "bad_json", "invalid JSON body")
		return
	}

	if req.Title == "" || req.RawText == "" {
		writeError(w, http.StatusBadRequest, "validation_error", "title and raw_text are required")
		return
	}

	if req.ID == "" {
		req.ID = uuid.NewString()
	}

	if req.Status == "" {
		req.Status = "new"
	}

	err := h.service.CreateTask(r.Context(), &req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "db_error", err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, req)
}

func (h *HTTPHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req model.TaskDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "bad_json", "invalid JSON body")
		return
	}

	if req.Title == "" || req.RawText == "" {
		writeError(w, http.StatusBadRequest, "validation_error", "title and raw_text are required")
		return
	}

	if req.ID == "" {
		req.ID = uuid.NewString()
	}

	if req.Status == "" {
		req.Status = "new"
	}

	err := h.service.CreateTask(r.Context(), &req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "db_error", err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, req)
}

func (h *HTTPHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if id == "" {
		writeError(w, http.StatusBadRequest, "bad_request", "missing task id")
		return
	}

	task, err := h.service.GetTask(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusNotFound, "not_found", "task not found")
		return
	}

	writeJSON(w, http.StatusOK, task)
}
