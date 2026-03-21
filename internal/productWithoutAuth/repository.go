package productWithoutAuth

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
	ErrProductWithoutAuthNotFound = errors.New("ProductWithoutAuth not found")

	ErrProductWithoutAuthCodeDuplicate = errors.New("ProductWithoutAuth code already exists")
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(productWithoutAuth ProductWithoutAuth) error {
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
		productWithoutAuth.ID,
		productWithoutAuth.Name,
		productWithoutAuth.Description,
		productWithoutAuth.Price,
		productWithoutAuth.Stock,
		productWithoutAuth.IsActive,
		productWithoutAuth.CreatedAt,
		productWithoutAuth.CreatedBy,
		productWithoutAuth.UpdatedAt,
		productWithoutAuth.UpdatedBy,
	)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return ErrProductWithoutAuthCodeDuplicate
			}
		}
		
		return fmt.Errorf("failed to save ProductWithoutAuth: %w", err)
	}

	return nil
}
func (r *Repository) Update(productWithoutAuth ProductWithoutAuth) error {
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
		productWithoutAuth.ID,
		productWithoutAuth.Name,
		productWithoutAuth.Description,
		productWithoutAuth.Price,
		productWithoutAuth.Stock,
		productWithoutAuth.IsActive,
		productWithoutAuth.CreatedAt,
		productWithoutAuth.CreatedBy,
		productWithoutAuth.UpdatedAt,
		productWithoutAuth.UpdatedBy,
	)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return ErrProductWithoutAuthCodeDuplicate
			}
		}
		
		return fmt.Errorf("failed to update ProductWithoutAuth: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrProductWithoutAuthNotFound
	}

	return nil
}

func (r *Repository) FindAll() ([]ProductWithoutAuth, error) {
	query := `
		SELECT COALESCE(ID, ''), COALESCE(Name, ''), COALESCE(Description, ''), Price, Stock, IsActive, CreatedAt, COALESCE(CreatedBy, ''), COALESCE(UpdatedAt, CreatedAt), COALESCE(UpdatedBy, '')
		FROM product
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query ProductWithoutAuth: %w", err)
	}
	defer rows.Close()

	 productWithoutAuthList := make([]ProductWithoutAuth, 0)
	for rows.Next() {
		var productWithoutAuth ProductWithoutAuth
		err := rows.Scan(
			&productWithoutAuth.ID,
			&productWithoutAuth.Name,
			&productWithoutAuth.Description,
			&productWithoutAuth.Price,
			&productWithoutAuth.Stock,
			&productWithoutAuth.IsActive,
			&productWithoutAuth.CreatedAt,
			&productWithoutAuth.CreatedBy,
			&productWithoutAuth.UpdatedAt,
			&productWithoutAuth.UpdatedBy		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan ProductWithoutAuth: %w", err)
		}
		productWithoutAuthList = append(productWithoutAuthList, productWithoutAuth)
	}

	return productWithoutAuthList , nil
}

func (r *Repository) FindById(id string) (ProductWithoutAuth, error) {
	query := `
		SELECT COALESCE(ID, ''), COALESCE(Name, ''), COALESCE(Description, ''), Price, Stock, IsActive, CreatedAt, COALESCE(CreatedBy, ''), COALESCE(UpdatedAt, CreatedAt), COALESCE(UpdatedBy, '')
		FROM product
		WHERE ID = ?
	`

	row := r.db.QueryRow(query, id)
	var productWithoutAuth ProductWithoutAuth
	err := row.Scan(
		&productWithoutAuth.ID,
		&productWithoutAuth.Name,
		&productWithoutAuth.Description,
		&productWithoutAuth.Price,
		&productWithoutAuth.Stock,
		&productWithoutAuth.IsActive,
		&productWithoutAuth.CreatedAt,
		&productWithoutAuth.CreatedBy,
		&productWithoutAuth.UpdatedAt,
		&productWithoutAuth.UpdatedBy	)
	if err != nil {
		if err == sql.ErrNoRows {
			return ProductWithoutAuth{}, ErrProductWithoutAuthNotFound
		}
		return ProductWithoutAuth{}, fmt.Errorf("failed to scan ProductWithoutAuth: %w", err)
	}
	return productWithoutAuth, nil
}

func (r *Repository) FindByCode(code string) (ProductWithoutAuth, error) {
	query := `
		SELECT COALESCE(ID, ''), COALESCE(Name, ''), COALESCE(Description, ''), Price, Stock, IsActive, CreatedAt, COALESCE(CreatedBy, ''), COALESCE(UpdatedAt, CreatedAt), COALESCE(UpdatedBy, '')
		FROM product
		WHERE Code = ?
	`

	row := r.db.QueryRow(query, code)
	var productWithoutAuth ProductWithoutAuth
	err := row.Scan(
		&productWithoutAuth.ID,
		&productWithoutAuth.Name,
		&productWithoutAuth.Description,
		&productWithoutAuth.Price,
		&productWithoutAuth.Stock,
		&productWithoutAuth.IsActive,
		&productWithoutAuth.CreatedAt,
		&productWithoutAuth.CreatedBy,
		&productWithoutAuth.UpdatedAt,
		&productWithoutAuth.UpdatedBy	)
	if err != nil {
		if err == sql.ErrNoRows {
			return ProductWithoutAuth{}, ErrProductWithoutAuthNotFound
		}
		return ProductWithoutAuth{}, fmt.Errorf("failed to scan ProductWithoutAuth: %w", err)
	}
	return productWithoutAuth, nil
}

func (r *Repository) SoftDelete(id string) error {
	query := `
		UPDATE product SET IsActive = false
		WHERE ID = ?
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete ProductWithoutAuth: %w", err)
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
		return fmt.Errorf("failed to delete ProductWithoutAuth: %w", err)
	}
	return nil
}

