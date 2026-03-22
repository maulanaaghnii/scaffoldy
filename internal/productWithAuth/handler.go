package productWithAuth

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

	group := router.Group("/product-with-auth")
	{
		group.GET("", h.GetAllProductWithAuth)
		group.POST("", h.CreateProductWithAuth)
		group.GET("/:id", h.GetProductWithAuthByID)
		group.GET("/code/:code", h.GetProductWithAuthByCode)
		group.PUT("/:id", h.UpdateProductWithAuth)
		group.DELETE("/:id", h.SoftDeleteProductWithAuth)
	}
}

// GetAllProductWithAuth
// GetProductWithAuthByID
// GetProductWithAuthByCode

// CreateProductWithAuth
// UpdateProductWithAuth
// SoftDeleteProductWithAuth

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAllProductWithAuth(c *gin.Context) {
	productWithAuth, err := h.service.GetAllProductWithAuth()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, productWithAuth)
}

func (h *Handler) GetProductWithAuthByID(c *gin.Context) {
	id := c.Param("id")

	productWithAuth, err := h.service.GetProductWithAuthByID(id)
	if err != nil {
		if errors.Is(err, ErrProductWithAuthNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, productWithAuth)
}

func (h *Handler) GetProductWithAuthByCode(c *gin.Context) {
	code := c.Param("code")

	productWithAuth, err := h.service.GetProductWithAuthByCode(code)
	if err != nil {
		if errors.Is(err, ErrProductWithAuthNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, productWithAuth)
}
func (h *Handler) CreateProductWithAuth(c *gin.Context) {
	var req CreateProductWithAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productWithAuth, err := h.service.CreateProductWithAuth(req)
	if err != nil {
		if errors.Is(err, ErrProductWithAuthCodeDuplicate) {
			response.Error(c, http.StatusConflict, err.Error())
			return
		}
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Created(c, productWithAuth)
}

func (h *Handler) UpdateProductWithAuth(c *gin.Context) {
	id := c.Param("id")

	var req UpdateProductWithAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	productWithAuth, err := h.service.UpdateProductWithAuth(id, req)
	if err != nil {
		if errors.Is(err, ErrProductWithAuthNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Success(c, productWithAuth)
		return
	}

	response.Success(c, productWithAuth)
}

func (h *Handler) SoftDeleteProductWithAuth(c *gin.Context) {
	id := c.Param("id")

	err := h.service.SoftDeleteProductWithAuth(id)
	if err != nil {
		if errors.Is(err, ErrProductWithAuthNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.NoContent(c)
}
