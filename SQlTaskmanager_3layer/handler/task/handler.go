package task

import (
	"SQLTaskmanager_3layer/models"
	servicePkg "SQLTaskmanager_3layer/service/task"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type Handler struct {
	service servicePkg.Service
}

func New(service servicePkg.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to read request body"))
		return
	}

	var task models.Task
	if err := json.Unmarshal(body, &task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON"))
		return
	}

	created, err := h.service.Create(task, task.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating task"))
		return
	}

	resp, _ := json.Marshal(created)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	tasks, err := h.service.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No tasks found"))
		return
	}

	resp, _ := json.Marshal(tasks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
