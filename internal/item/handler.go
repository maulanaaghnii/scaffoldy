
package item

import (
	"errors"
	"scaffoldy/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllItem
// GetItemByID
// GetItemByCode

// CreateItem
// UpdateItem
// SoftDeleteItem


type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAllItem(c *gin.Context) {
	item, err := h.service.GetAllItem()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, item)
}

func (h *Handler) GetItemByID(c *gin.Context) {
	id := c.Param("id")

	item, err := h.service.GetItemByID(id)
	if err != nil {
		if errors.Is(err, ErrItemNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, item)
}

func (h *Handler) GetItemByCode(c *gin.Context) {
	code := c.Param("code")

	item, err := h.service.GetItemByCode(code)
	if err != nil {
		if errors.Is(err, ErrItemNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, item)
}
func (h *Handler) CreateItem(c *gin.Context) {
	var req CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.service.CreateItem(req)
	if err != nil {
		if errors.Is(err, ErrItemCodeDuplicate) {
			response.Error(c, http.StatusConflict, err.Error())
			return
		}
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Created(c, item)
}

func (h *Handler) UpdateItem(c *gin.Context) {
	id := c.Param("id")

	var req UpdateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	item, err := h.service.UpdateItem(id, req)
	if err != nil {
		if errors.Is(err, ErrItemNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Success(c, item)
		return
	}

	response.Success(c, item)
}

func (h *Handler) SoftDeleteItem(c *gin.Context) {
	id := c.Param("id")

	err := h.service.SoftDeleteItem(id)
	if err != nil {
		if errors.Is(err, ErrItemNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.NoContent(c)
}

