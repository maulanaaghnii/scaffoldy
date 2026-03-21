package productWithAuth

import "time"

type CreateProductWithAuthRequest struct {
	ID string `json:"iD"` 
	Name string `json:"name"` 
	Description string `json:"description"` 
	Price float64 `json:"price"` 
	Stock int `json:"stock"` 
	IsActive bool `json:"isActive"` 
	CreatedAt time.Time `json:"createdAt"` 
	CreatedBy string `json:"createdBy"` 
	UpdatedAt time.Time `json:"updatedAt"` 
	UpdatedBy string `json:"updatedBy"` 
} 

type UpdateProductWithAuthRequest struct {
	ID string `json:"iD"` 
	Name string `json:"name"` 
	Description string `json:"description"` 
	Price float64 `json:"price"` 
	Stock int `json:"stock"` 
	IsActive bool `json:"isActive"` 
	CreatedAt time.Time `json:"createdAt"` 
	CreatedBy string `json:"createdBy"` 
	UpdatedAt time.Time `json:"updatedAt"` 
	UpdatedBy string `json:"updatedBy"` 
}