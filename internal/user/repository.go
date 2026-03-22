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
		A.ID, A.Username, A.Password, 
		COALESCE(A.FullName, ''), 
		COALESCE(A.Email, ''), 
		A.IsActive, 
		COALESCE(A.RefreshToken, ''),
		A.CreatedAt, A.CreatedBy, 
		COALESCE(A.UpdatedAt, A.CreatedAt), 
		COALESCE(A.UpdatedBy, '')
	FROM initial_users A
	WHERE A.Username = ? AND A.IsActive = true
	`

	var u User
	err := r.db.QueryRow(query, username).Scan(
		&u.ID,
		&u.Username,
		&u.Password,
		&u.FullName,
		&u.Email,
		&u.IsActive,
		&u.RefreshToken,
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
		INSERT INTO initial_users (ID, Username, Password, FullName, Email, IsActive, RefreshToken, CreatedAt, CreatedBy)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(query,
		u.ID,
		u.Username,
		u.Password,
		u.FullName,
		u.Email,
		u.IsActive,
		u.RefreshToken,
		u.CreatedAt,
		u.CreatedBy,
	)

	if err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}

func (r *Repository) UpdateRefreshToken(userID, token string) error {
	query := `UPDATE initial_users SET RefreshToken = ? WHERE ID = ?`
	_, err := r.db.Exec(query, token, userID)
	if err != nil {
		return fmt.Errorf("failed to update refresh token: %w", err)
	}
	return nil
}

func (r *Repository) FindByRefreshToken(token string) (User, error) {
	query := `
	SELECT 
		A.ID, A.Username, A.Password, 
		COALESCE(A.FullName, ''), 
		COALESCE(A.Email, ''), 
		A.IsActive, 
		COALESCE(A.RefreshToken, ''),
		A.CreatedAt, A.CreatedBy, 
		COALESCE(A.UpdatedAt, A.CreatedAt), 
		COALESCE(A.UpdatedBy, '')
	FROM initial_users A
	WHERE A.RefreshToken = ? AND A.IsActive = true
	`

	var u User
	err := r.db.QueryRow(query, token).Scan(
		&u.ID,
		&u.Username,
		&u.Password,
		&u.FullName,
		&u.Email,
		&u.IsActive,
		&u.RefreshToken,
		&u.CreatedAt,
		&u.CreatedBy,
		&u.UpdatedAt,
		&u.UpdatedBy,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, ErrUserNotFound
		}
		return User{}, fmt.Errorf("failed to find user by refresh token: %w", err)
	}

	return u, nil
}
