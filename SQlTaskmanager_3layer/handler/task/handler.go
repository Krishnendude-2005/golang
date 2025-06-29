package task

import (
	"SQLTaskmanager_3layer/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Service interface {
	Create(task models.Task, userID int) (models.Task, error)
	GetById(userID int) ([]models.Task, error)
	DeleteTaskById(id int) (int, error)
	Update(task models.Task, taskID int) (models.Task, error)
	GetAll() ([]models.Task, error)
}
type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed reading body", http.StatusBadRequest)
		return
	}

	var task models.Task
	if err := json.Unmarshal(body, &task); err != nil {
		http.Error(w, "Failed parsing body", http.StatusBadRequest)
		return
	}

	created, err := h.service.Create(task, task.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonResp, _ := json.Marshal(created)
	w.Header().Set("Content-Type", "application/json")
	val, err := w.Write(jsonResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Println("bytes used", val)
}

func (h *Handler) GetByUserID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("user_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	tasks, err := h.service.GetById(id)
	if err != nil {
		http.Error(w, "Tasks not found", http.StatusNotFound)
		return
	}

	jsonResp, _ := json.Marshal(tasks)
	w.Header().Set("Content-Type", "application/json")
	val, err := w.Write(jsonResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("bytes used", val)
}

func (h *Handler) DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if _, err := h.service.DeleteTaskById(id); err != nil {
		http.Error(w, "Delete failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	val, err := w.Write([]byte("Task deleted"))
	if err != nil {
		http.Error(w, "Delete failed", http.StatusInternalServerError)
	}
	fmt.Println("bytes used", val)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	taskIDStr := r.URL.Query().Get("id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed reading body", http.StatusBadRequest)
		return
	}

	var task models.Task
	if err := json.Unmarshal(body, &task); err != nil {
		http.Error(w, "Failed parsing body", http.StatusBadRequest)
		return
	}

	updated, err := h.service.Update(task, taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, _ := json.Marshal(updated)
	w.Header().Set("Content-Type", "application/json")
	val, err := w.Write(jsonResp)
	if err != nil {
		http.Error(w, "Failed writing body", http.StatusInternalServerError)
	} else {
		fmt.Println("bytes used", val)
	}

}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
		return
	}

	jsonResp, _ := json.Marshal(tasks)
	w.Header().Set("Content-Type", "application/json")
	val, err := w.Write(jsonResp)
	if err != nil {
		http.Error(w, "Failed writing body", http.StatusInternalServerError)
	}
	fmt.Println("bytes used", val)
}
