
package productWithoutAuth

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

	group := router.Group("/product-without-auth")
	{
		group.GET("", h.GetAllProductWithoutAuth)
		group.POST("", h.CreateProductWithoutAuth)
		group.GET("/:id", h.GetProductWithoutAuthByID)
		group.GET("/code/:code", h.GetProductWithoutAuthByCode)
		group.PUT("/:id", h.UpdateProductWithoutAuth)
		group.DELETE("/:id", h.SoftDeleteProductWithoutAuth)
	}
}

// GetAllProductWithoutAuth
// GetProductWithoutAuthByID
// GetProductWithoutAuthByCode

// CreateProductWithoutAuth
// UpdateProductWithoutAuth
// SoftDeleteProductWithoutAuth


type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAllProductWithoutAuth(c *gin.Context) {
	productWithoutAuth, err := h.service.GetAllProductWithoutAuth()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, productWithoutAuth)
}

func (h *Handler) GetProductWithoutAuthByID(c *gin.Context) {
	id := c.Param("id")

	productWithoutAuth, err := h.service.GetProductWithoutAuthByID(id)
	if err != nil {
		if errors.Is(err, ErrProductWithoutAuthNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, productWithoutAuth)
}

func (h *Handler) GetProductWithoutAuthByCode(c *gin.Context) {
	code := c.Param("code")

	productWithoutAuth, err := h.service.GetProductWithoutAuthByCode(code)
	if err != nil {
		if errors.Is(err, ErrProductWithoutAuthNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, productWithoutAuth)
}
func (h *Handler) CreateProductWithoutAuth(c *gin.Context) {
	var req CreateProductWithoutAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productWithoutAuth, err := h.service.CreateProductWithoutAuth(req)
	if err != nil {
		if errors.Is(err, ErrProductWithoutAuthCodeDuplicate) {
			response.Error(c, http.StatusConflict, err.Error())
			return
		}
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Created(c, productWithoutAuth)
}

func (h *Handler) UpdateProductWithoutAuth(c *gin.Context) {
	id := c.Param("id")

	var req UpdateProductWithoutAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	productWithoutAuth, err := h.service.UpdateProductWithoutAuth(id, req)
	if err != nil {
		if errors.Is(err, ErrProductWithoutAuthNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Success(c, productWithoutAuth)
		return
	}

	response.Success(c, productWithoutAuth)
}

func (h *Handler) SoftDeleteProductWithoutAuth(c *gin.Context) {
	id := c.Param("id")

	err := h.service.SoftDeleteProductWithoutAuth(id)
	if err != nil {
		if errors.Is(err, ErrProductWithoutAuthNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.NoContent(c)
}

