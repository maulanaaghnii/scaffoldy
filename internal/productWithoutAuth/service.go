package productWithoutAuth

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

// GetAllProductWithoutAuth
// GetProductWithoutAuthByID
// GetProductWithoutAuthByCode
// CreateProductWithoutAuth
// UpdateProductWithoutAuth
// DeleteProductWithoutAuth
func (s *Service) GetAllProductWithoutAuth() ([]ProductWithoutAuth, error) {
	return s.repository.FindAll()
}

func (s *Service) GetProductWithoutAuthByID(id string) (ProductWithoutAuth, error) {
	return s.repository.FindById(id)
}

func (s *Service) GetProductWithoutAuthByCode(code string) (ProductWithoutAuth, error) {
	return s.repository.FindByCode(code)
}

func (s *Service) CreateProductWithoutAuth(req CreateProductWithoutAuthRequest) (ProductWithoutAuth, error) {
	if err := s.validateCreateRequest(req); err != nil {
		return ProductWithoutAuth{}, err
	}
	productWithoutAuth := ProductWithoutAuth{
		ID: uuid.New().String(),
		Name: strings.TrimSpace(req.Name),
		Description: strings.TrimSpace(req.Description),
		Price: req.Price,
		Stock: req.Stock,
		IsActive: true,
		AuditTrails: shared.AuditTrails{
			CreatedAt: time.Now(),
			CreatedBy: "system",
		},
	}
	err := s.repository.Save(productWithoutAuth)
	return productWithoutAuth, err
}

func (s *Service) UpdateProductWithoutAuth(id string, req UpdateProductWithoutAuthRequest) (ProductWithoutAuth, error) {
	existing, err := s.repository.FindById(id)
	if err != nil {
		return ProductWithoutAuth{}, err
	}
	if err := s.validateUpdateRequest(req); err != nil {
		return ProductWithoutAuth{}, err
	}
	existing.Name = strings.TrimSpace(req.Name)
	existing.Description = strings.TrimSpace(req.Description)
	existing.Price = req.Price
	existing.Stock = req.Stock
	existing.IsActive = true
	existing.UpdatedAt = time.Now()
	existing.UpdatedBy = "system"
err = s.repository.Update(existing)
return existing, err
}

func (s *Service) SoftDeleteProductWithoutAuth(id string) error {
	return s.repository.SoftDelete(id)
}

func (s *Service) DeleteProductWithoutAuth(id string) error {
	return s.repository.Delete(id)
}



func (s *Service) validateCreateRequest(req CreateProductWithoutAuthRequest) error { 
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

func (s *Service) validateUpdateRequest(req UpdateProductWithoutAuthRequest) error {
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

