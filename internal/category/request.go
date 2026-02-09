package category

import "time"

type CreateCategoryRequest struct {
	ID string `json:"iD"` 
	Domain string `json:"domain"` 
	Code string `json:"code"` 
	Name string `json:"name"` 
	Description string `json:"description"` 
	IsActive bool `json:"isActive"` 
	CreatedAt time.Time `json:"createdAt"` 
	CreatedBy string `json:"createdBy"` 
} 

type UpdateCategoryRequest struct {
	ID string `json:"iD"` 
	Domain string `json:"domain"` 
	Code string `json:"code"` 
	Name string `json:"name"` 
	Description string `json:"description"` 
	IsActive bool `json:"isActive"` 
	UpdatedAt time.Time `json:"updatedAt"` 
	UpdatedBy string `json:"updatedBy"` 
}