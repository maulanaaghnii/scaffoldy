package item

import "time"

type CreateItemRequest struct {
	ID string `json:"iD"` 
	Code string `json:"code"` 
	Name string `json:"name"` 
	Description string `json:"description"` 
	Category string `json:"category"` 
	Unit string `json:"unit"` 
	Price float64 `json:"price"` 
	Stock int `json:"stock"` 
	IsActive bool `json:"isActive"` 
	CreatedAt time.Time `json:"createdAt"` 
	CreatedBy string `json:"createdBy"` 
} 

type UpdateItemRequest struct {
	ID string `json:"iD"` 
	Code string `json:"code"` 
	Name string `json:"name"` 
	Description string `json:"description"` 
	Category string `json:"category"` 
	Unit string `json:"unit"` 
	Price float64 `json:"price"` 
	Stock int `json:"stock"` 
	IsActive bool `json:"isActive"` 
	UpdatedAt time.Time `json:"updatedAt"` 
	UpdatedBy string `json:"updatedBy"` 
}