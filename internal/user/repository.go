package user

import (
	"database/sql"
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindByUsername(username string) (User, error) {
	query := `
		SELECT 
			ID, Username, Password, 
			COALESCE(FullName, ''), 
			COALESCE(Email, ''), 
			IsActive, CreatedAt, CreatedBy, 
			COALESCE(UpdatedAt, CreatedAt), 
			COALESCE(UpdatedBy, '')
		FROM users
		WHERE Username = ? AND IsActive = true
	`

	var u User
	err := r.db.QueryRow(query, username).Scan(
		&u.ID,
		&u.Username,
		&u.Password,
		&u.FullName,
		&u.Email,
		&u.IsActive,
		&u.CreatedAt,
		&u.CreatedBy,
		&u.UpdatedAt,
		&u.UpdatedBy,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, ErrUserNotFound
		}
		return User{}, fmt.Errorf("failed to find user: %w", err)
	}

	return u, nil
}

func (r *Repository) Save(u User) error {
	query := `
		INSERT INTO users (ID, Username, Password, FullName, Email, IsActive, CreatedAt, CreatedBy)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(query,
		u.ID,
		u.Username,
		u.Password,
		u.FullName,
		u.Email,
		u.IsActive,
		u.CreatedAt,
		u.CreatedBy,
	)

	if err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}
