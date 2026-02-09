package item

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

// GetAllItem
// GetItemByID
// GetItemByCode
// CreateItem
// UpdateItem
// DeleteItem
func (s *Service) GetAllItem() ([]Item, error) {
	return s.repository.FindAll()
}

func (s *Service) GetItemByID(id string) (Item, error) {
	return s.repository.FindById(id)
}

func (s *Service) GetItemByCode(code string) (Item, error) {
	return s.repository.FindByCode(code)
}

func (s *Service) CreateItem(req CreateItemRequest) (Item, error) {
	if err := s.validateCreateRequest(req); err != nil {
		return Item{}, err
	}
	item := Item{
		ID: uuid.New().String(),
		Code: strings.ToUpper(strings.TrimSpace(req.Code)),
		Name: strings.TrimSpace(req.Name),
		Description: strings.TrimSpace(req.Description),
		Category: strings.TrimSpace(req.Category),
		Unit: strings.TrimSpace(req.Unit),
		Price: req.Price,
		Stock: req.Stock,
		IsActive: true,
		AuditTrails: shared.AuditTrails{
			CreatedAt: time.Now(),
			CreatedBy: "system",
		},
	}
	err := s.repository.Save(item)
	return item, err
}

func (s *Service) UpdateItem(id string, req UpdateItemRequest) (Item, error) {
	existing, err := s.repository.FindById(id)
	if err != nil {
		return Item{}, err
	}
	if err := s.validateUpdateRequest(req); err != nil {
		return Item{}, err
	}
	existing.Name = strings.TrimSpace(req.Name)
	existing.Description = strings.TrimSpace(req.Description)
	existing.Category = strings.TrimSpace(req.Category)
	existing.Unit = strings.TrimSpace(req.Unit)
	existing.Price = req.Price
	existing.Stock = req.Stock
	existing.IsActive = true
	existing.UpdatedAt = time.Now()
	existing.UpdatedBy = "system"
err = s.repository.Update(existing)
return existing, err
}

func (s *Service) SoftDeleteItem(id string) error {
	return s.repository.SoftDelete(id)
}

func (s *Service) DeleteItem(id string) error {
	return s.repository.Delete(id)
}



func (s *Service) validateCreateRequest(req CreateItemRequest) error { 
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

func (s *Service) validateUpdateRequest(req UpdateItemRequest) error {
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

