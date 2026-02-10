package task

import "time"

type CreateTaskRequest struct {
	ID string `json:"iD"` 
	Title string `json:"title"` 
	Description *string `json:"description"` 
	Status string `json:"status"` 
	CreatedAt time.Time `json:"createdAt"` 
	CreatedBy string `json:"createdBy"` 
	UpdatedAt time.Time `json:"updatedAt"` 
	UpdatedBy string `json:"updatedBy"` 
} 

type UpdateTaskRequest struct {
	ID string `json:"iD"` 
	Title string `json:"title"` 
	Description *string `json:"description"` 
	Status string `json:"status"` 
	CreatedAt time.Time `json:"createdAt"` 
	CreatedBy string `json:"createdBy"` 
	UpdatedAt time.Time `json:"updatedAt"` 
	UpdatedBy string `json:"updatedBy"` 
}