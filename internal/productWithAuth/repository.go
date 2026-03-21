package productWithAuth

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql")

// Save
// Update
// FindAll
// FindById
// FindByCode
// SoftDelete
// Delete
var (
	ErrProductWithAuthNotFound = errors.New("ProductWithAuth not found")

	ErrProductWithAuthCodeDuplicate = errors.New("ProductWithAuth code already exists")
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(productWithAuth ProductWithAuth) error {
	query := `
		INSERT INTO product (
			ID,
			Name,
			Description,
			Price,
			Stock,
			IsActive,
			CreatedAt,
			CreatedBy,
			UpdatedAt,
			UpdatedBy
		)
		VALUES (?,?,?,?,?,?,?,?,?,?
		)
	`
	_, err := r.db.Exec(query,
		productWithAuth.ID,
		productWithAuth.Name,
		productWithAuth.Description,
		productWithAuth.Price,
		productWithAuth.Stock,
		productWithAuth.IsActive,
		productWithAuth.CreatedAt,
		productWithAuth.CreatedBy,
		productWithAuth.UpdatedAt,
		productWithAuth.UpdatedBy,
	)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return ErrProductWithAuthCodeDuplicate
			}
		}
		
		return fmt.Errorf("failed to save ProductWithAuth: %w", err)
	}

	return nil
}
func (r *Repository) Update(productWithAuth ProductWithAuth) error {
	query := `
		UPDATE product SET 
			ID = ?,
			Name = ?,
			Description = ?,
			Price = ?,
			Stock = ?,
			IsActive = ?,
			CreatedAt = ?,
			CreatedBy = ?,
			UpdatedAt = ?,
			UpdatedBy = ?
		WHERE ID = ?
	`

	result, err := r.db.Exec(query,
		productWithAuth.ID,
		productWithAuth.Name,
		productWithAuth.Description,
		productWithAuth.Price,
		productWithAuth.Stock,
		productWithAuth.IsActive,
		productWithAuth.CreatedAt,
		productWithAuth.CreatedBy,
		productWithAuth.UpdatedAt,
		productWithAuth.UpdatedBy,
	)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return ErrProductWithAuthCodeDuplicate
			}
		}
		
		return fmt.Errorf("failed to update ProductWithAuth: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrProductWithAuthNotFound
	}

	return nil
}

func (r *Repository) FindAll() ([]ProductWithAuth, error) {
	query := `
		SELECT COALESCE(ID, ''), COALESCE(Name, ''), COALESCE(Description, ''), Price, Stock, IsActive, CreatedAt, COALESCE(CreatedBy, ''), COALESCE(UpdatedAt, CreatedAt), COALESCE(UpdatedBy, '')
		FROM product
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query ProductWithAuth: %w", err)
	}
	defer rows.Close()

	 productWithAuthList := make([]ProductWithAuth, 0)
	for rows.Next() {
		var productWithAuth ProductWithAuth
		err := rows.Scan(
			&productWithAuth.ID,
			&productWithAuth.Name,
			&productWithAuth.Description,
			&productWithAuth.Price,
			&productWithAuth.Stock,
			&productWithAuth.IsActive,
			&productWithAuth.CreatedAt,
			&productWithAuth.CreatedBy,
			&productWithAuth.UpdatedAt,
			&productWithAuth.UpdatedBy		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan ProductWithAuth: %w", err)
		}
		productWithAuthList = append(productWithAuthList, productWithAuth)
	}

	return productWithAuthList , nil
}

func (r *Repository) FindById(id string) (ProductWithAuth, error) {
	query := `
		SELECT COALESCE(ID, ''), COALESCE(Name, ''), COALESCE(Description, ''), Price, Stock, IsActive, CreatedAt, COALESCE(CreatedBy, ''), COALESCE(UpdatedAt, CreatedAt), COALESCE(UpdatedBy, '')
		FROM product
		WHERE ID = ?
	`

	row := r.db.QueryRow(query, id)
	var productWithAuth ProductWithAuth
	err := row.Scan(
		&productWithAuth.ID,
		&productWithAuth.Name,
		&productWithAuth.Description,
		&productWithAuth.Price,
		&productWithAuth.Stock,
		&productWithAuth.IsActive,
		&productWithAuth.CreatedAt,
		&productWithAuth.CreatedBy,
		&productWithAuth.UpdatedAt,
		&productWithAuth.UpdatedBy	)
	if err != nil {
		if err == sql.ErrNoRows {
			return ProductWithAuth{}, ErrProductWithAuthNotFound
		}
		return ProductWithAuth{}, fmt.Errorf("failed to scan ProductWithAuth: %w", err)
	}
	return productWithAuth, nil
}

func (r *Repository) FindByCode(code string) (ProductWithAuth, error) {
	query := `
		SELECT COALESCE(ID, ''), COALESCE(Name, ''), COALESCE(Description, ''), Price, Stock, IsActive, CreatedAt, COALESCE(CreatedBy, ''), COALESCE(UpdatedAt, CreatedAt), COALESCE(UpdatedBy, '')
		FROM product
		WHERE Code = ?
	`

	row := r.db.QueryRow(query, code)
	var productWithAuth ProductWithAuth
	err := row.Scan(
		&productWithAuth.ID,
		&productWithAuth.Name,
		&productWithAuth.Description,
		&productWithAuth.Price,
		&productWithAuth.Stock,
		&productWithAuth.IsActive,
		&productWithAuth.CreatedAt,
		&productWithAuth.CreatedBy,
		&productWithAuth.UpdatedAt,
		&productWithAuth.UpdatedBy	)
	if err != nil {
		if err == sql.ErrNoRows {
			return ProductWithAuth{}, ErrProductWithAuthNotFound
		}
		return ProductWithAuth{}, fmt.Errorf("failed to scan ProductWithAuth: %w", err)
	}
	return productWithAuth, nil
}

func (r *Repository) SoftDelete(id string) error {
	query := `
		UPDATE product SET IsActive = false
		WHERE ID = ?
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete ProductWithAuth: %w", err)
	}
	return nil
}

func (r *Repository) Delete(id string) error {
	query := `
		DELETE FROM product
		WHERE ID = ?
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete ProductWithAuth: %w", err)
	}
	return nil
}

