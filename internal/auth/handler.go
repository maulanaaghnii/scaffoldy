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
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type AuthHandler struct {
	userRepo *user.Repository
}

func Register(router *gin.RouterGroup, db *sql.DB) {
	h := &AuthHandler{
		userRepo: user.NewRepository(db),
	}
	router.POST("/login", h.Login)
	router.POST("/refresh", h.RefreshToken)
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

	// 3. Generate Access Token
	accessToken, err := utils.GenerateToken(u.ID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate access token")
		return
	}

	// 4. Generate Refresh Token
	refreshToken, err := utils.GenerateRefreshToken(u.ID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate refresh token")
		return
	}

	// 5. Store Refresh Token in DB
	err = h.userRepo.UpdateRefreshToken(u.ID, refreshToken)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to store refresh token")
		return
	}

	response.Success(c, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user": gin.H{
			"username":  u.Username,
			"full_name": u.FullName,
			"role":      u.Role,
			"email":     u.Email,
			"roleCode":  u.Role,
		},
	})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// 1. Validate JWT structure and expiration
	token, err := utils.ValidateToken(req.RefreshToken)
	if err != nil || !token.Valid {
		response.Error(c, http.StatusUnauthorized, "Invalid refresh token")
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["type"] != "refresh" {
		response.Error(c, http.StatusUnauthorized, "Invalid refresh token type")
		return
	}

	// 2. Verify token exists in database (Anti-revocation)
	u, err := h.userRepo.FindByRefreshToken(req.RefreshToken)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Refresh token not found or expired")
		return
	}

	// 3. Generate New Access Token
	newAccessToken, err := utils.GenerateToken(u.ID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate access token")
		return
	}

	// 4. Generate New Refresh Token (Rotation)
	newRefreshToken, err := utils.GenerateRefreshToken(u.ID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate refresh token")
		return
	}

	// 5. Update Refresh Token in DB
	err = h.userRepo.UpdateRefreshToken(u.ID, newRefreshToken)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update refresh token")
		return
	}

	response.Success(c, gin.H{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}
