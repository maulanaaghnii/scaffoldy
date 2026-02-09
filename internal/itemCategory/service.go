package itemCategory

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

// GetAllItemCategory
// GetItemCategoryByID
// GetItemCategoryByCode
// CreateItemCategory
// UpdateItemCategory
// DeleteItemCategory
func (s *Service) GetAllItemCategory() ([]ItemCategory, error) {
	return s.repository.FindAll()
}

func (s *Service) GetItemCategoryByID(id string) (ItemCategory, error) {
	return s.repository.FindById(id)
}

func (s *Service) GetItemCategoryByCode(code string) (ItemCategory, error) {
	return s.repository.FindByCode(code)
}

func (s *Service) CreateItemCategory(req CreateItemCategoryRequest) (ItemCategory, error) {
	if err := s.validateCreateRequest(req); err != nil {
		return ItemCategory{}, err
	}
	itemCategory := ItemCategory{
		ID: uuid.New().String(),
		Code: strings.ToUpper(strings.TrimSpace(req.Code)),
		Name: strings.TrimSpace(req.Name),
		Description: strings.TrimSpace(req.Description),
		IsActive: true,
		AuditTrails: shared.AuditTrails{
			CreatedAt: time.Now(),
			CreatedBy: "system",
		},
	}
	err := s.repository.Save(itemCategory)
	return itemCategory, err
}

func (s *Service) UpdateItemCategory(id string, req UpdateItemCategoryRequest) (ItemCategory, error) {
	existing, err := s.repository.FindById(id)
	if err != nil {
		return ItemCategory{}, err
	}
	if err := s.validateUpdateRequest(req); err != nil {
		return ItemCategory{}, err
	}
	existing.Name = strings.TrimSpace(req.Name)
	existing.Description = strings.TrimSpace(req.Description)
	existing.IsActive = true
	existing.UpdatedAt = time.Now()
	existing.UpdatedBy = "system"
err = s.repository.Update(existing)
return existing, err
}

func (s *Service) SoftDeleteItemCategory(id string) error {
	return s.repository.SoftDelete(id)
}

func (s *Service) DeleteItemCategory(id string) error {
	return s.repository.Delete(id)
}



func (s *Service) validateCreateRequest(req CreateItemCategoryRequest) error { 
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

func (s *Service) validateUpdateRequest(req UpdateItemCategoryRequest) error {
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

