package task

import (
	"scaffoldy/internal/shared"
)

type Task struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Status      string  `json:"status"`
	shared.AuditTrails
}
