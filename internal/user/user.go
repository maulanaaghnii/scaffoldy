package user

import (
	"scaffoldy/internal/shared"
)

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"-"` // Hashed password, never export to JSON
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	IsActive     bool   `json:"is_active"`
	RefreshToken string `json:"-"`
	shared.AuditTrails
}
