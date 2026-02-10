package category

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
	ErrCategoryNotFound = errors.New("Category not found")

	ErrCategoryCodeDuplicate = errors.New("Category code already exists")
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(category Category) error {
	query := `
		INSERT INTO tblcategory (
			ID,
			Domain,
			Code,
			Name,
			Description,
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
		category.ID,
		category.Domain,
		category.Code,
		category.Name,
		category.Description,
		category.IsActive,
		category.CreatedAt,
		category.CreatedBy,
		category.UpdatedAt,
		category.UpdatedBy,
	)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return ErrCategoryCodeDuplicate
			}
		}
		
		return fmt.Errorf("failed to save Category: %w", err)
	}

	return nil
}
func (r *Repository) Update(category Category) error {
	query := `
		UPDATE tblcategory SET 
			ID = ?,
			Domain = ?,
			Code = ?,
			Name = ?,
			Description = ?,
			IsActive = ?,
			CreatedAt = ?,
			CreatedBy = ?,
			UpdatedAt = ?,
			UpdatedBy = ?
		WHERE ID = ?
	`

	result, err := r.db.Exec(query,
		category.ID,
		category.Domain,
		category.Code,
		category.Name,
		category.Description,
		category.IsActive,
		category.CreatedAt,
		category.CreatedBy,
		category.UpdatedAt,
		category.UpdatedBy,
	)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return ErrCategoryCodeDuplicate
			}
		}
		
		return fmt.Errorf("failed to update Category: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrCategoryNotFound
	}

	return nil
}

func (r *Repository) FindAll() ([]Category, error) {
	query := `
		SELECT ID, Domain, Code, Name, Description, IsActive, CreatedAt, CreatedBy, UpdatedAt, UpdatedBy
		FROM tblcategory
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query Category: %w", err)
	}
	defer rows.Close()

	 categoryList := make([]Category, 0)
	for rows.Next() {
		var category Category
		err := rows.Scan(
			&category.ID,
			&category.Domain,
			&category.Code,
			&category.Name,
			&category.Description,
			&category.IsActive,
			&category.CreatedAt,
			&category.CreatedBy,
			&category.UpdatedAt,
			&category.UpdatedBy		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan Category: %w", err)
		}
		categoryList = append(categoryList, category)
	}

	return categoryList , nil
}

func (r *Repository) FindById(id string) (Category, error) {
	query := `
		SELECT ID, Domain, Code, Name, Description, IsActive, CreatedAt, CreatedBy, UpdatedAt, UpdatedBy
		FROM tblcategory
		WHERE ID = ?
	`

	row := r.db.QueryRow(query, id)
	var category Category
	err := row.Scan(
		&category.ID,
		&category.Domain,
		&category.Code,
		&category.Name,
		&category.Description,
		&category.IsActive,
		&category.CreatedAt,
		&category.CreatedBy,
		&category.UpdatedAt,
		&category.UpdatedBy	)
	if err != nil {
		if err == sql.ErrNoRows {
			return Category{}, ErrCategoryNotFound
		}
		return Category{}, fmt.Errorf("failed to scan Category: %w", err)
	}
	return category, nil
}

func (r *Repository) FindByCode(code string) (Category, error) {
	query := `
		SELECT ID, Domain, Code, Name, Description, IsActive, CreatedAt, CreatedBy, UpdatedAt, UpdatedBy
		FROM tblcategory
		WHERE Code = ?
	`

	row := r.db.QueryRow(query, code)
	var category Category
	err := row.Scan(
		&category.ID,
		&category.Domain,
		&category.Code,
		&category.Name,
		&category.Description,
		&category.IsActive,
		&category.CreatedAt,
		&category.CreatedBy,
		&category.UpdatedAt,
		&category.UpdatedBy	)
	if err != nil {
		if err == sql.ErrNoRows {
			return Category{}, ErrCategoryNotFound
		}
		return Category{}, fmt.Errorf("failed to scan Category: %w", err)
	}
	return category, nil
}

func (r *Repository) SoftDelete(id string) error {
	query := `
		UPDATE tblcategory SET IsActive = false
		WHERE ID = ?
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete Category: %w", err)
	}
	return nil
}

func (r *Repository) Delete(id string) error {
	query := `
		DELETE FROM tblcategory
		WHERE ID = ?
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete Category: %w", err)
	}
	return nil
}

