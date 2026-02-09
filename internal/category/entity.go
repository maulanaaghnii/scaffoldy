package category

import "scaffoldy/internal/shared"

type Category struct {
	ID          string `json:"id"`
	Domain      string `json:"domain"`
	Code        string `json:"Code"`
	Name        string `json:"Name"`
	Description string `json:"description"`
	IsActive    bool   `json:"isactive"`
	shared.AuditTrails
}
