package initialUser

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

// GetAllInitialUser
// GetInitialUserByID
// GetInitialUserByCode
// CreateInitialUser
// UpdateInitialUser
// DeleteInitialUser
func (s *Service) GetAllInitialUser() ([]InitialUser, error) {
	return s.repository.FindAll()
}

func (s *Service) GetInitialUserByID(id string) (InitialUser, error) {
	return s.repository.FindById(id)
}

func (s *Service) GetInitialUserByCode(code string) (InitialUser, error) {
	return s.repository.FindByCode(code)
}

func (s *Service) CreateInitialUser(req CreateInitialUserRequest) (InitialUser, error) {
	if err := s.validateCreateRequest(req); err != nil {
		return InitialUser{}, err
	}
	initialUser := InitialUser{
		ID: uuid.New().String(),
		Username: strings.TrimSpace(req.Username),
		Password: strings.TrimSpace(req.Password),
		FullName: strings.TrimSpace(req.FullName),
		Email: strings.TrimSpace(req.Email),
		IsActive: true,
		RefreshToken: strings.TrimSpace(req.RefreshToken),
		AuditTrails: shared.AuditTrails{
			CreatedAt: time.Now(),
			CreatedBy: "system",
		},
	}
	err := s.repository.Save(initialUser)
	return initialUser, err
}

func (s *Service) UpdateInitialUser(id string, req UpdateInitialUserRequest) (InitialUser, error) {
	existing, err := s.repository.FindById(id)
	if err != nil {
		return InitialUser{}, err
	}
	if err := s.validateUpdateRequest(req); err != nil {
		return InitialUser{}, err
	}
	existing.Username = strings.TrimSpace(req.Username)
	existing.Password = strings.TrimSpace(req.Password)
	existing.FullName = strings.TrimSpace(req.FullName)
	existing.Email = strings.TrimSpace(req.Email)
	existing.IsActive = true
	existing.RefreshToken = strings.TrimSpace(req.RefreshToken)
	existing.UpdatedAt = time.Now()
	existing.UpdatedBy = "system"
err = s.repository.Update(existing)
return existing, err
}

func (s *Service) SoftDeleteInitialUser(id string) error {
	return s.repository.SoftDelete(id)
}

func (s *Service) DeleteInitialUser(id string) error {
	return s.repository.Delete(id)
}



func (s *Service) validateCreateRequest(req CreateInitialUserRequest) error { 
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

func (s *Service) validateUpdateRequest(req UpdateInitialUserRequest) error {
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

