package category

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

// GetAllCategory
// GetCategoryByID
// GetCategoryByCode
// CreateCategory
// UpdateCategory
// DeleteCategory
func (s *Service) GetAllCategory() ([]Category, error) {
	return s.repository.FindAll()
}

func (s *Service) GetCategoryByID(id string) (Category, error) {
	return s.repository.FindById(id)
}

func (s *Service) GetCategoryByCode(code string) (Category, error) {
	return s.repository.FindByCode(code)
}

func (s *Service) CreateCategory(req CreateCategoryRequest) (Category, error) {
	if err := s.validateCreateRequest(req); err != nil {
		return Category{}, err
	}
	category := Category{
		ID: uuid.New().String(),
		Domain: strings.TrimSpace(req.Domain),
		Code: strings.ToUpper(strings.TrimSpace(req.Code)),
		Name: strings.TrimSpace(req.Name),
		Description: strings.TrimSpace(req.Description),
		IsActive: true,
		AuditTrails: shared.AuditTrails{
			CreatedAt: time.Now(),
			CreatedBy: "system",
		},
	}
	err := s.repository.Save(category)
	return category, err
}

func (s *Service) UpdateCategory(id string, req UpdateCategoryRequest) (Category, error) {
	existing, err := s.repository.FindById(id)
	if err != nil {
		return Category{}, err
	}
	if err := s.validateUpdateRequest(req); err != nil {
		return Category{}, err
	}
	existing.Domain = strings.TrimSpace(req.Domain)
	existing.Name = strings.TrimSpace(req.Name)
	existing.Description = strings.TrimSpace(req.Description)
	existing.IsActive = true
	existing.UpdatedAt = time.Now()
	existing.UpdatedBy = "system"
err = s.repository.Update(existing)
return existing, err
}

func (s *Service) SoftDeleteCategory(id string) error {
	return s.repository.SoftDelete(id)
}

func (s *Service) DeleteCategory(id string) error {
	return s.repository.Delete(id)
}



func (s *Service) validateCreateRequest(req CreateCategoryRequest) error { 
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

func (s *Service) validateUpdateRequest(req UpdateCategoryRequest) error {
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

