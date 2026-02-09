package itemCategory

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
	ErrItemCategoryNotFound = errors.New("ItemCategory not found")

	ErrItemCategoryCodeDuplicate = errors.New("ItemCategory code already exists")
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(itemCategory ItemCategory) error {
	query := `
		INSERT INTO tblitemcategory (
			ID,
			Code,
			Name,
			Description,
			IsActive,
			CreatedAt,
			CreatedBy,
			UpdatedAt,
			UpdatedBy
		)
		VALUES (?,?,?,?,?,?,?,?,?
		)
	`
	_, err := r.db.Exec(query,
		itemCategory.ID,
		itemCategory.Code,
		itemCategory.Name,
		itemCategory.Description,
		itemCategory.IsActive,
		itemCategory.CreatedAt,
		itemCategory.CreatedBy,
		itemCategory.UpdatedAt,
		itemCategory.UpdatedBy,
	)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return ErrItemCategoryCodeDuplicate
			}
		}
		
		return fmt.Errorf("failed to save ItemCategory: %w", err)
	}

	return nil
}
func (r *Repository) Update(itemCategory ItemCategory) error {
	query := `
		UPDATE tblitemcategory SET 
			ID = ?,
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
		itemCategory.ID,
		itemCategory.Code,
		itemCategory.Name,
		itemCategory.Description,
		itemCategory.IsActive,
		itemCategory.CreatedAt,
		itemCategory.CreatedBy,
		itemCategory.UpdatedAt,
		itemCategory.UpdatedBy,
	)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return ErrItemCategoryCodeDuplicate
			}
		}
		
		return fmt.Errorf("failed to update ItemCategory: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrItemCategoryNotFound
	}

	return nil
}

func (r *Repository) FindAll() ([]ItemCategory, error) {
	query := `
		SELECT ID, Code, Name, Description, IsActive, CreatedAt, CreatedBy, UpdatedAt, UpdatedBy
		FROM tblitemcategory
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query ItemCategory: %w", err)
	}
	defer rows.Close()

	 itemCategoryList := make([]ItemCategory, 0)
	for rows.Next() {
		var itemCategory ItemCategory
		err := rows.Scan(
			&itemCategory.ID,
			&itemCategory.Code,
			&itemCategory.Name,
			&itemCategory.Description,
			&itemCategory.IsActive,
			&itemCategory.CreatedAt,
			&itemCategory.CreatedBy,
			&itemCategory.UpdatedAt,
			&itemCategory.UpdatedBy		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan ItemCategory: %w", err)
		}
		itemCategoryList = append(itemCategoryList, itemCategory)
	}

	return itemCategoryList , nil
}

func (r *Repository) FindById(id string) (ItemCategory, error) {
	query := `
		SELECT ID, Code, Name, Description, IsActive, CreatedAt, CreatedBy, UpdatedAt, UpdatedBy
		FROM tblitemcategory
		WHERE ID = ?
	`

	row := r.db.QueryRow(query, id)
	var itemCategory ItemCategory
	err := row.Scan(
		&itemCategory.ID,
		&itemCategory.Code,
		&itemCategory.Name,
		&itemCategory.Description,
		&itemCategory.IsActive,
		&itemCategory.CreatedAt,
		&itemCategory.CreatedBy,
		&itemCategory.UpdatedAt,
		&itemCategory.UpdatedBy	)
	if err != nil {
		if err == sql.ErrNoRows {
			return ItemCategory{}, ErrItemCategoryNotFound
		}
		return ItemCategory{}, fmt.Errorf("failed to scan ItemCategory: %w", err)
	}
	return itemCategory, nil
}

func (r *Repository) FindByCode(code string) (ItemCategory, error) {
	query := `
		SELECT ID, Code, Name, Description, IsActive, CreatedAt, CreatedBy, UpdatedAt, UpdatedBy
		FROM tblitemcategory
		WHERE Code = ?
	`

	row := r.db.QueryRow(query, code)
	var itemCategory ItemCategory
	err := row.Scan(
		&itemCategory.ID,
		&itemCategory.Code,
		&itemCategory.Name,
		&itemCategory.Description,
		&itemCategory.IsActive,
		&itemCategory.CreatedAt,
		&itemCategory.CreatedBy,
		&itemCategory.UpdatedAt,
		&itemCategory.UpdatedBy	)
	if err != nil {
		if err == sql.ErrNoRows {
			return ItemCategory{}, ErrItemCategoryNotFound
		}
		return ItemCategory{}, fmt.Errorf("failed to scan ItemCategory: %w", err)
	}
	return itemCategory, nil
}

func (r *Repository) SoftDelete(id string) error {
	query := `
		UPDATE tblitemcategory SET IsActive = false
		WHERE ID = ?
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete ItemCategory: %w", err)
	}
	return nil
}

func (r *Repository) Delete(id string) error {
	query := `
		DELETE FROM tblitemcategory
		WHERE ID = ?
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete ItemCategory: %w", err)
	}
	return nil
}

