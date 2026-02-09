package item

import "scaffoldy/internal/shared"

type Item struct {
	ID          string  `json:"id"`
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Unit        string  `json:"unit"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	IsActive    bool    `json:"is_active"`
	shared.AuditTrails
}
