package task

import (
"strings"
"time"

"github.com/google/uuid"
"scaffoldy/internal/shared"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

// GetAllTask
// GetTaskByID
// GetTaskByCode
// CreateTask
// UpdateTask
// DeleteTask
func (s *Service) GetAllTask() ([]Task, error) {
	return s.repository.FindAll()
}

func (s *Service) GetTaskByID(id string) (Task, error) {
	return s.repository.FindById(id)
}

func (s *Service) GetTaskByCode(code string) (Task, error) {
	return s.repository.FindByCode(code)
}

func (s *Service) CreateTask(req CreateTaskRequest) (Task, error) {
	if err := s.validateCreateRequest(req); err != nil {
		return Task{}, err
	}
	task := Task{
		ID: uuid.New().String(),
		Title: strings.TrimSpace(req.Title),
		Description: req.Description,
		Status: strings.TrimSpace(req.Status),
		AuditTrails: shared.AuditTrails{
			CreatedAt: time.Now(),
			CreatedBy: "system",
		},
	}
	err := s.repository.Save(task)
	return task, err
}

func (s *Service) UpdateTask(id string, req UpdateTaskRequest) (Task, error) {
	existing, err := s.repository.FindById(id)
	if err != nil {
		return Task{}, err
	}
	if err := s.validateUpdateRequest(req); err != nil {
		return Task{}, err
	}
	existing.Title = strings.TrimSpace(req.Title)
	existing.Description = req.Description
	existing.Status = strings.TrimSpace(req.Status)
	existing.UpdatedAt = time.Now()
	existing.UpdatedBy = "system"
err = s.repository.Update(existing)
return existing, err
}

func (s *Service) SoftDeleteTask(id string) error {
	return s.repository.SoftDelete(id)
}

func (s *Service) DeleteTask(id string) error {
	return s.repository.Delete(id)
}



func (s *Service) validateCreateRequest(req CreateTaskRequest) error { 
	// if strings.TrimSpace(req.Code) == "" {
	// 	return fmt.Errorf("code is required")
	// }
	// if strings.TrimSpace(req.Name) == "" {
	// 	return fmt.Errorf("name is required")
	// }
	// if req.Price < 0 {
	// 	return fmt.Errorf("price cannot be negative")
	// }
	// if req.Stock < 0 {
	// 	return fmt.Errorf("stock cannot be negative")
	// }
	return nil
}

func (s *Service) validateUpdateRequest(req UpdateTaskRequest) error {
	// if strings.TrimSpace(req.Name) == "" {
	// 	return fmt.Errorf("name is required")
	// }
	// if req.Price < 0 {
	// 	return fmt.Errorf("price cannot be negative")
	// }
	// if req.Stock < 0 {
	// 	return fmt.Errorf("stock cannot be negative")
	// }
	return nil
}

