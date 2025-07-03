package task

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
	var task models.Task
	if err := c.Bind(&task); err != nil {
		return nil, err
	}
	return h.service.Create(c, task, task.UserID)
}

func (h *Handler) GetById(c *gofr.Context) (interface{}, error) {
	idStr := c.PathParam("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return nil, err
	}

	return h.service.GetById(c, id)
}

func (h *Handler) DeleteTaskById(c *gofr.Context) (interface{}, error) {
	idStr := c.PathParam("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return nil, err
	}

	return h.service.DeleteTaskById(c, id)
}

func (h *Handler) Update(c *gofr.Context) (interface{}, error) {
	idStr := c.PathParam("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return nil, err
	}

	var task models.Task
	if err := c.Bind(&task); err != nil {
		return nil, err
	}

	return h.service.Update(c, task, id)
}

func (h *Handler) GetAll(c *gofr.Context) (interface{}, error) {
	return h.service.GetAll(c)
}
