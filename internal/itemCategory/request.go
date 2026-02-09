package itemCategory

import "time"

type CreateItemCategoryRequest struct {
	ID string `json:"iD"` 
	Code string `json:"code"` 
	Name string `json:"name"` 
	Description string `json:"description"` 
	IsActive bool `json:"isActive"` 
	CreatedAt time.Time `json:"createdAt"` 
	CreatedBy string `json:"createdBy"` 
	UpdatedAt time.Time `json:"updatedAt"` 
	UpdatedBy string `json:"updatedBy"` 
} 

type UpdateItemCategoryRequest struct {
	ID string `json:"iD"` 
	Code string `json:"code"` 
	Name string `json:"name"` 
	Description string `json:"description"` 
	IsActive bool `json:"isActive"` 
	CreatedAt time.Time `json:"createdAt"` 
	CreatedBy string `json:"createdBy"` 
	UpdatedAt time.Time `json:"updatedAt"` 
	UpdatedBy string `json:"updatedBy"` 
}