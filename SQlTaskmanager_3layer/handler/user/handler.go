package user

import (
	"SQLTaskmanager_3layer/models"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type Service interface {
	Create(user models.User) (models.User, error)
	GetById(id int) (models.User, error)
}
type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{service: service}
}

// Create handles POST /user/add
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, errSms := w.Write([]byte("Failed to read request body"))
		if errSms != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, errMe := w.Write([]byte("Invalid JSON"))
		if errMe != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	created, err := h.service.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, errMsg := w.Write([]byte("User creation failed: " + err.Error()))
		if errMsg != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	respBytes, _ := json.Marshal(created)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, errorSms := w.Write(respBytes)
	if errorSms != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetByID handles GET /user/find?id=...
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Missing user ID"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, errorMsg := w.Write([]byte("Invalid user ID"))
		if errorMsg != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	user, err := h.service.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, errSms := w.Write([]byte("User not found"))
		if errSms != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	respBytes, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, errMsg := w.Write(respBytes)
	if errMsg != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
