package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"scaffoldy/internal/user"
	"scaffoldy/pkg/response"
	"scaffoldy/pkg/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthHandler struct {
	userRepo *user.Repository
}

func Register(router *gin.RouterGroup, db *sql.DB) {
	h := &AuthHandler{
		userRepo: user.NewRepository(db),
	}
	router.POST("/login", h.Login)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// 1. Find user in DB
	u, err := h.userRepo.FindByUsername(req.Username)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			response.Error(c, http.StatusUnauthorized, "Invalid username or password")
			return
		}
		// Tambahkan log untuk melihat error aslinya di terminal
		fmt.Printf("[LOGIN ERROR] FindByUsername: %v\n", err)
		response.Error(c, http.StatusInternalServerError, "Internal server error: "+err.Error())
		return
	}

	// 2. Compare password hash
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password))
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	// 3. Generate JWT
	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	response.Success(c, gin.H{
		"token": token,
		"user": gin.H{
			"username":  u.Username,
			"full_name": u.FullName,
			"email":     u.Email,
		},
	})
}
