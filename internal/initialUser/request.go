package initialUser

import "time"

type CreateInitialUserRequest struct {
	ID string `json:"iD"` 
	Username string `json:"username"` 
	Password string `json:"password"` 
	FullName string `json:"fullName"` 
	Email string `json:"email"` 
	IsActive bool `json:"isActive"` 
	RefreshToken string `json:"refreshToken"` 
	CreatedAt time.Time `json:"createdAt"` 
	CreatedBy string `json:"createdBy"` 
	UpdatedAt time.Time `json:"updatedAt"` 
	UpdatedBy string `json:"updatedBy"` 
} 

type UpdateInitialUserRequest struct {
	ID string `json:"iD"` 
	Username string `json:"username"` 
	Password string `json:"password"` 
	FullName string `json:"fullName"` 
	Email string `json:"email"` 
	IsActive bool `json:"isActive"` 
	RefreshToken string `json:"refreshToken"` 
	CreatedAt time.Time `json:"createdAt"` 
	CreatedBy string `json:"createdBy"` 
	UpdatedAt time.Time `json:"updatedAt"` 
	UpdatedBy string `json:"updatedBy"` 
}