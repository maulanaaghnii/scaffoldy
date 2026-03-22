package initialUser

import (
	"database/sql"
	"errors"
	"net/http"
	"scaffoldy/pkg/response"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup, db *sql.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	h := NewHandler(svc)

	group := router.Group("/initial-user")
	{
		group.GET("", h.GetAllInitialUser)
		group.POST("", h.CreateInitialUser)
		group.GET("/:id", h.GetInitialUserByID)
		group.GET("/code/:code", h.GetInitialUserByCode)
		group.PUT("/:id", h.UpdateInitialUser)
		group.DELETE("/:id", h.SoftDeleteInitialUser)
	}
}

// GetAllInitialUser
// GetInitialUserByID
// GetInitialUserByCode

// CreateInitialUser
// UpdateInitialUser
// SoftDeleteInitialUser

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAllInitialUser(c *gin.Context) {
	initialUser, err := h.service.GetAllInitialUser()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, initialUser)
}

func (h *Handler) GetInitialUserByID(c *gin.Context) {
	id := c.Param("id")

	initialUser, err := h.service.GetInitialUserByID(id)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, initialUser)
}

func (h *Handler) GetInitialUserByCode(c *gin.Context) {
	code := c.Param("code")

	initialUser, err := h.service.GetInitialUserByCode(code)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, initialUser)
}
func (h *Handler) CreateInitialUser(c *gin.Context) {
	var req CreateInitialUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initialUser, err := h.service.CreateInitialUser(req)
	if err != nil {
		if errors.Is(err, ErrUserCodeDuplicate) {
			response.Error(c, http.StatusConflict, err.Error())
			return
		}
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Created(c, initialUser)
}

func (h *Handler) UpdateInitialUser(c *gin.Context) {
	id := c.Param("id")

	var req UpdateInitialUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	initialUser, err := h.service.UpdateInitialUser(id, req)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Success(c, initialUser)
		return
	}

	response.Success(c, initialUser)
}

func (h *Handler) SoftDeleteInitialUser(c *gin.Context) {
	id := c.Param("id")

	err := h.service.SoftDeleteInitialUser(id)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.NoContent(c)
}
