package task

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
	ErrTaskNotFound = errors.New("Task not found")

	ErrTaskCodeDuplicate = errors.New("Task code already exists")
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(task Task) error {
	query := `
		INSERT INTO tbltask (
			ID,
			Title,
			Description,
			Status,
			CreatedAt,
			CreatedBy,
			UpdatedAt,
			UpdatedBy
		)
		VALUES (?,?,?,?,?,?,?,?
		)
	`
	_, err := r.db.Exec(query,
		task.ID,
		task.Title,
		task.Description,
		task.Status,
		task.CreatedAt,
		task.CreatedBy,
		task.UpdatedAt,
		task.UpdatedBy,
	)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return ErrTaskCodeDuplicate
			}
		}
		
		return fmt.Errorf("failed to save Task: %w", err)
	}

	return nil
}
func (r *Repository) Update(task Task) error {
	query := `
		UPDATE tbltask SET 
			ID = ?,
			Title = ?,
			Description = ?,
			Status = ?,
			CreatedAt = ?,
			CreatedBy = ?,
			UpdatedAt = ?,
			UpdatedBy = ?
		WHERE ID = ?
	`

	result, err := r.db.Exec(query,
		task.ID,
		task.Title,
		task.Description,
		task.Status,
		task.CreatedAt,
		task.CreatedBy,
		task.UpdatedAt,
		task.UpdatedBy,
	)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return ErrTaskCodeDuplicate
			}
		}
		
		return fmt.Errorf("failed to update Task: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrTaskNotFound
	}

	return nil
}

func (r *Repository) FindAll() ([]Task, error) {
	query := `
		SELECT COALESCE(ID, ''), COALESCE(Title, ''), Description, COALESCE(Status, ''), CreatedAt, COALESCE(CreatedBy, ''), COALESCE(UpdatedAt, CreatedAt), COALESCE(UpdatedBy, '')
		FROM tbltask
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query Task: %w", err)
	}
	defer rows.Close()

	 taskList := make([]Task, 0)
	for rows.Next() {
		var task Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.CreatedBy,
			&task.UpdatedAt,
			&task.UpdatedBy		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan Task: %w", err)
		}
		taskList = append(taskList, task)
	}

	return taskList , nil
}

func (r *Repository) FindById(id string) (Task, error) {
	query := `
		SELECT COALESCE(ID, ''), COALESCE(Title, ''), Description, COALESCE(Status, ''), CreatedAt, COALESCE(CreatedBy, ''), COALESCE(UpdatedAt, CreatedAt), COALESCE(UpdatedBy, '')
		FROM tbltask
		WHERE ID = ?
	`

	row := r.db.QueryRow(query, id)
	var task Task
	err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
		&task.CreatedBy,
		&task.UpdatedAt,
		&task.UpdatedBy	)
	if err != nil {
		if err == sql.ErrNoRows {
			return Task{}, ErrTaskNotFound
		}
		return Task{}, fmt.Errorf("failed to scan Task: %w", err)
	}
	return task, nil
}

func (r *Repository) FindByCode(code string) (Task, error) {
	query := `
		SELECT COALESCE(ID, ''), COALESCE(Title, ''), Description, COALESCE(Status, ''), CreatedAt, COALESCE(CreatedBy, ''), COALESCE(UpdatedAt, CreatedAt), COALESCE(UpdatedBy, '')
		FROM tbltask
		WHERE Code = ?
	`

	row := r.db.QueryRow(query, code)
	var task Task
	err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
		&task.CreatedBy,
		&task.UpdatedAt,
		&task.UpdatedBy	)
	if err != nil {
		if err == sql.ErrNoRows {
			return Task{}, ErrTaskNotFound
		}
		return Task{}, fmt.Errorf("failed to scan Task: %w", err)
	}
	return task, nil
}

func (r *Repository) SoftDelete(id string) error {
	query := `
		UPDATE tbltask SET IsActive = false
		WHERE ID = ?
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete Task: %w", err)
	}
	return nil
}

func (r *Repository) Delete(id string) error {
	query := `
		DELETE FROM tbltask
		WHERE ID = ?
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete Task: %w", err)
	}
	return nil
}

