package user

import (
	"SQLTaskmanager_3layer/models"
	"gofr.dev/pkg/gofr"
	"strconv"
)

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gofr.Context) (interface{}, error) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return nil, err
	}
	return h.service.Create(c, user)
}

func (h *Handler) GetById(c *gofr.Context) (interface{}, error) {
	idStr := c.PathParam("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return nil, err
	}
	return h.service.GetById(c, id)
}
