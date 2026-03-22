package initialUser

import (
	"scaffoldy/internal/shared"
)

type InitialUser struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"-"` // Hashed password, never export to JSON
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	IsActive     bool   `json:"is_active"`
	RefreshToken string `json:"-"`
	shared.AuditTrails
}
