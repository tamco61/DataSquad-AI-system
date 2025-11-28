package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *HTTPHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if id == "" {
		writeError(w, http.StatusBadRequest, "bad_request", "missing task id")
		return
	}

	user, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusNotFound, "not_found", "user not found")
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (h *HTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {}

func (h *HTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {}
