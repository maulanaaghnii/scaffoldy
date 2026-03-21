package productWithAuth

import "scaffoldy/internal/shared"

type ProductWithAuth struct {
	ID          string  `json:"id"`
	Name        string  `json:"Name"`
	Description string  `json:"Desc"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	IsActive    bool    `json:"is_active"`
	shared.AuditTrails
}
