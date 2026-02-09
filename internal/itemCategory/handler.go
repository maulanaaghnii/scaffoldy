
package itemCategory

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

	group := router.Group("/itemCategory")
	{
		group.GET("", h.GetAllItemCategory)
		group.POST("", h.CreateItemCategory)
		group.GET("/:id", h.GetItemCategoryByID)
		group.GET("/code/:code", h.GetItemCategoryByCode)
		group.PUT("/:id", h.UpdateItemCategory)
		group.DELETE("/:id", h.SoftDeleteItemCategory)
	}
}

// GetAllItemCategory
// GetItemCategoryByID
// GetItemCategoryByCode

// CreateItemCategory
// UpdateItemCategory
// SoftDeleteItemCategory


type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAllItemCategory(c *gin.Context) {
	itemCategory, err := h.service.GetAllItemCategory()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, itemCategory)
}

func (h *Handler) GetItemCategoryByID(c *gin.Context) {
	id := c.Param("id")

	itemCategory, err := h.service.GetItemCategoryByID(id)
	if err != nil {
		if errors.Is(err, ErrItemCategoryNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, itemCategory)
}

func (h *Handler) GetItemCategoryByCode(c *gin.Context) {
	code := c.Param("code")

	itemCategory, err := h.service.GetItemCategoryByCode(code)
	if err != nil {
		if errors.Is(err, ErrItemCategoryNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, itemCategory)
}
func (h *Handler) CreateItemCategory(c *gin.Context) {
	var req CreateItemCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	itemCategory, err := h.service.CreateItemCategory(req)
	if err != nil {
		if errors.Is(err, ErrItemCategoryCodeDuplicate) {
			response.Error(c, http.StatusConflict, err.Error())
			return
		}
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Created(c, itemCategory)
}

func (h *Handler) UpdateItemCategory(c *gin.Context) {
	id := c.Param("id")

	var req UpdateItemCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	itemCategory, err := h.service.UpdateItemCategory(id, req)
	if err != nil {
		if errors.Is(err, ErrItemCategoryNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Success(c, itemCategory)
		return
	}

	response.Success(c, itemCategory)
}

func (h *Handler) SoftDeleteItemCategory(c *gin.Context) {
	id := c.Param("id")

	err := h.service.SoftDeleteItemCategory(id)
	if err != nil {
		if errors.Is(err, ErrItemCategoryNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.NoContent(c)
}

