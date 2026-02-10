
package task

import (
	"errors"
	"scaffoldy/pkg/response"
	"net/http"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup, db *sql.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	h := NewHandler(svc)

	group := router.Group("/task")
	{
		group.GET("", h.GetAllTask)
		group.POST("", h.CreateTask)
		group.GET("/:id", h.GetTaskByID)
		group.GET("/code/:code", h.GetTaskByCode)
		group.PUT("/:id", h.UpdateTask)
		group.DELETE("/:id", h.SoftDeleteTask)
	}
}

// GetAllTask
// GetTaskByID
// GetTaskByCode

// CreateTask
// UpdateTask
// SoftDeleteTask


type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAllTask(c *gin.Context) {
	task, err := h.service.GetAllTask()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, task)
}

func (h *Handler) GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		if errors.Is(err, ErrTaskNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, task)
}

func (h *Handler) GetTaskByCode(c *gin.Context) {
	code := c.Param("code")

	task, err := h.service.GetTaskByCode(code)
	if err != nil {
		if errors.Is(err, ErrTaskNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, task)
}
func (h *Handler) CreateTask(c *gin.Context) {
	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.service.CreateTask(req)
	if err != nil {
		if errors.Is(err, ErrTaskCodeDuplicate) {
			response.Error(c, http.StatusConflict, err.Error())
			return
		}
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Created(c, task)
}

func (h *Handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var req UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	task, err := h.service.UpdateTask(id, req)
	if err != nil {
		if errors.Is(err, ErrTaskNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Success(c, task)
		return
	}

	response.Success(c, task)
}

func (h *Handler) SoftDeleteTask(c *gin.Context) {
	id := c.Param("id")

	err := h.service.SoftDeleteTask(id)
	if err != nil {
		if errors.Is(err, ErrTaskNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.NoContent(c)
}

