package scaffoldcomponents

func HandlerContent() string {
	handlerContent := `
package {{.DomainNameLower}}

import (
	"errors"
	"scaffoldy/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAll{{.DomainName}}
// Get{{.DomainName}}ByID
// Get{{.DomainName}}ByCode

// Create{{.DomainName}}
// Update{{.DomainName}}
// SoftDelete{{.DomainName}}


type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAll{{.DomainName}}(c *gin.Context) {
	{{.DomainNameLower}}, err := h.service.GetAll{{.DomainName}}()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, {{.DomainNameLower}})
}

func (h *Handler) Get{{.DomainName}}ByID(c *gin.Context) {
	id := c.Param("id")

	{{.DomainNameLower}}, err := h.service.Get{{.DomainName}}ByID(id)
	if err != nil {
		if errors.Is(err, Err{{.DomainName}}NotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, {{.DomainNameLower}})
}

func (h *Handler) Get{{.DomainName}}ByCode(c *gin.Context) {
	code := c.Param("code")

	{{.DomainNameLower}}, err := h.service.Get{{.DomainName}}ByCode(code)
	if err != nil {
		if errors.Is(err, Err{{.DomainName}}NotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, {{.DomainNameLower}})
}
func (h *Handler) Create{{.DomainName}}(c *gin.Context) {
	var req Create{{.DomainName}}Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	{{.DomainNameLower}}, err := h.service.Create{{.DomainName}}(req)
	if err != nil {
		if errors.Is(err, Err{{.DomainName}}CodeDuplicate) {
			response.Error(c, http.StatusConflict, err.Error())
			return
		}
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Created(c, {{.DomainNameLower}})
}

func (h *Handler) Update{{.DomainName}}(c *gin.Context) {
	id := c.Param("id")

	var req Update{{.DomainName}}Request
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	{{.DomainNameLower}}, err := h.service.Update{{.DomainName}}(id, req)
	if err != nil {
		if errors.Is(err, Err{{.DomainName}}NotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Success(c, {{.DomainNameLower}})
		return
	}

	response.Success(c, {{.DomainNameLower}})
}

func (h *Handler) SoftDelete{{.DomainName}}(c *gin.Context) {
	id := c.Param("id")

	err := h.service.SoftDelete{{.DomainName}}(id)
	if err != nil {
		if errors.Is(err, Err{{.DomainName}}NotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.NoContent(c)
}

`

	return handlerContent
}
