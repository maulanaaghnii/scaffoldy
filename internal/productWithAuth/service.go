package productWithAuth

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

// GetAllProductWithAuth
// GetProductWithAuthByID
// GetProductWithAuthByCode
// CreateProductWithAuth
// UpdateProductWithAuth
// DeleteProductWithAuth
func (s *Service) GetAllProductWithAuth() ([]ProductWithAuth, error) {
	return s.repository.FindAll()
}

func (s *Service) GetProductWithAuthByID(id string) (ProductWithAuth, error) {
	return s.repository.FindById(id)
}

func (s *Service) GetProductWithAuthByCode(code string) (ProductWithAuth, error) {
	return s.repository.FindByCode(code)
}

func (s *Service) CreateProductWithAuth(req CreateProductWithAuthRequest) (ProductWithAuth, error) {
	if err := s.validateCreateRequest(req); err != nil {
		return ProductWithAuth{}, err
	}
	productWithAuth := ProductWithAuth{
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
	err := s.repository.Save(productWithAuth)
	return productWithAuth, err
}

func (s *Service) UpdateProductWithAuth(id string, req UpdateProductWithAuthRequest) (ProductWithAuth, error) {
	existing, err := s.repository.FindById(id)
	if err != nil {
		return ProductWithAuth{}, err
	}
	if err := s.validateUpdateRequest(req); err != nil {
		return ProductWithAuth{}, err
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

func (s *Service) SoftDeleteProductWithAuth(id string) error {
	return s.repository.SoftDelete(id)
}

func (s *Service) DeleteProductWithAuth(id string) error {
	return s.repository.Delete(id)
}



func (s *Service) validateCreateRequest(req CreateProductWithAuthRequest) error { 
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

func (s *Service) validateUpdateRequest(req UpdateProductWithAuthRequest) error {
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

