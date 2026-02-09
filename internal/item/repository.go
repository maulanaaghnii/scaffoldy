package item

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
	ErrItemNotFound = errors.New("Item not found")

	ErrItemCodeDuplicate = errors.New("Item code already exists")
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(item Item) error {
	query := `
		INSERT INTO tblitem (
			ID,
			Code,
			Name,
			Description,
			Category,
			Unit,
			Price,
			Stock,
			IsActive,
			CreatedAt,
			CreatedBy,
			UpdatedAt,
			UpdatedBy
		)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?
		)
	`
	_, err := r.db.Exec(query,
		item.ID,
		item.Code,
		item.Name,
		item.Description,
		item.Category,
		item.Unit,
		item.Price,
		item.Stock,
		item.IsActive,
		item.CreatedAt,
		item.CreatedBy,
		item.UpdatedAt,
		item.UpdatedBy,
	)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return ErrItemCodeDuplicate
			}
		}
		
		return fmt.Errorf("failed to save Item: %w", err)
	}

	return nil
}
func (r *Repository) Update(item Item) error {
	query := `
		UPDATE tblitem SET 
			ID = ?,
			Code = ?,
			Name = ?,
			Description = ?,
			Category = ?,
			Unit = ?,
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
		item.ID,
		item.Code,
		item.Name,
		item.Description,
		item.Category,
		item.Unit,
		item.Price,
		item.Stock,
		item.IsActive,
		item.CreatedAt,
		item.CreatedBy,
		item.UpdatedAt,
		item.UpdatedBy,
	)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return ErrItemCodeDuplicate
			}
		}
		
		return fmt.Errorf("failed to update Item: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrItemNotFound
	}

	return nil
}

func (r *Repository) FindAll() ([]Item, error) {
	query := `
		SELECT ID, Code, Name, Description, Category, Unit, Price, Stock, IsActive, CreatedAt, CreatedBy, UpdatedAt, UpdatedBy
		FROM tblitem
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query Item: %w", err)
	}
	defer rows.Close()

	 itemList := make([]Item, 0)
	for rows.Next() {
		var item Item
		err := rows.Scan(
			&item.ID,
			&item.Code,
			&item.Name,
			&item.Description,
			&item.Category,
			&item.Unit,
			&item.Price,
			&item.Stock,
			&item.IsActive,
			&item.CreatedAt,
			&item.CreatedBy,
			&item.UpdatedAt,
			&item.UpdatedBy		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan Item: %w", err)
		}
		itemList = append(itemList, item)
	}

	return itemList , nil
}

func (r *Repository) FindById(id string) (Item, error) {
	query := `
		SELECT ID, Code, Name, Description, Category, Unit, Price, Stock, IsActive, CreatedAt, CreatedBy, UpdatedAt, UpdatedBy
		FROM tblitem
		WHERE ID = ?
	`

	row := r.db.QueryRow(query, id)
	var item Item
	err := row.Scan(
		&item.ID,
		&item.Code,
		&item.Name,
		&item.Description,
		&item.Category,
		&item.Unit,
		&item.Price,
		&item.Stock,
		&item.IsActive,
		&item.CreatedAt,
		&item.CreatedBy,
		&item.UpdatedAt,
		&item.UpdatedBy	)
	if err != nil {
		if err == sql.ErrNoRows {
			return Item{}, ErrItemNotFound
		}
		return Item{}, fmt.Errorf("failed to scan Item: %w", err)
	}
	return item, nil
}

func (r *Repository) FindByCode(code string) (Item, error) {
	query := `
		SELECT ID, Code, Name, Description, Category, Unit, Price, Stock, IsActive, CreatedAt, CreatedBy, UpdatedAt, UpdatedBy
		FROM tblitem
		WHERE Code = ?
	`

	row := r.db.QueryRow(query, code)
	var item Item
	err := row.Scan(
		&item.ID,
		&item.Code,
		&item.Name,
		&item.Description,
		&item.Category,
		&item.Unit,
		&item.Price,
		&item.Stock,
		&item.IsActive,
		&item.CreatedAt,
		&item.CreatedBy,
		&item.UpdatedAt,
		&item.UpdatedBy	)
	if err != nil {
		if err == sql.ErrNoRows {
			return Item{}, ErrItemNotFound
		}
		return Item{}, fmt.Errorf("failed to scan Item: %w", err)
	}
	return item, nil
}

func (r *Repository) SoftDelete(id string) error {
	query := `
		UPDATE tblitem SET IsActive = false
		WHERE ID = ?
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete Item: %w", err)
	}
	return nil
}

func (r *Repository) Delete(id string) error {
	query := `
		DELETE FROM tblitem
		WHERE ID = ?
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete Item: %w", err)
	}
	return nil
}

