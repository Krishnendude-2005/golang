package user

import (
	"SQLTaskmanager_3layer/models"
	servicePkg "SQLTaskmanager_3layer/service/user"
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

// Create handles POST /user/add
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to read request body"))
		return
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON"))
		return
	}

	created, err := h.service.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User creation failed: " + err.Error()))
		return
	}

	respBytes, _ := json.Marshal(created)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
}

// GetByID handles GET /user/find?id=...
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing user ID"))
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user ID"))
		return
	}

	user, err := h.service.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	respBytes, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
}
