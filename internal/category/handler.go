
package category

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

	group := router.Group("/category")
	{
		group.GET("", h.GetAllCategory)
		group.POST("", h.CreateCategory)
		group.GET("/:id", h.GetCategoryByID)
		group.GET("/code/:code", h.GetCategoryByCode)
		group.PUT("/:id", h.UpdateCategory)
		group.DELETE("/:id", h.SoftDeleteCategory)
	}
}

// GetAllCategory
// GetCategoryByID
// GetCategoryByCode

// CreateCategory
// UpdateCategory
// SoftDeleteCategory


type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAllCategory(c *gin.Context) {
	category, err := h.service.GetAllCategory()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, category)
}

func (h *Handler) GetCategoryByID(c *gin.Context) {
	id := c.Param("id")

	category, err := h.service.GetCategoryByID(id)
	if err != nil {
		if errors.Is(err, ErrCategoryNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, category)
}

func (h *Handler) GetCategoryByCode(c *gin.Context) {
	code := c.Param("code")

	category, err := h.service.GetCategoryByCode(code)
	if err != nil {
		if errors.Is(err, ErrCategoryNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, category)
}
func (h *Handler) CreateCategory(c *gin.Context) {
	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.service.CreateCategory(req)
	if err != nil {
		if errors.Is(err, ErrCategoryCodeDuplicate) {
			response.Error(c, http.StatusConflict, err.Error())
			return
		}
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Created(c, category)
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	var req UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	category, err := h.service.UpdateCategory(id, req)
	if err != nil {
		if errors.Is(err, ErrCategoryNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Success(c, category)
		return
	}

	response.Success(c, category)
}

func (h *Handler) SoftDeleteCategory(c *gin.Context) {
	id := c.Param("id")

	err := h.service.SoftDeleteCategory(id)
	if err != nil {
		if errors.Is(err, ErrCategoryNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.NoContent(c)
}

