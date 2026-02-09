package itemCategory

import "scaffoldy/internal/shared"

type ItemCategory struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
	shared.AuditTrails
}
