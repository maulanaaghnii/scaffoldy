package initialUser

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// Save
// Update
// FindAll
// FindById
// FindByCode
// SoftDelete
// Delete
var (
	ErrUserNotFound = errors.New("User not found")

	ErrUserCodeDuplicate = errors.New("User code already exists")
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(initialUser InitialUser) error {
	query := `
		INSERT INTO initial_user (
			ID,
			Username,
			Password,
			FullName,
			Email,
			IsActive,
			RefreshToken,
			CreatedAt,
			CreatedBy,
			UpdatedAt,
			UpdatedBy
		)
		VALUES (?,?,?,?,?,?,?,?,?,?,?
		)
	`
	_, err := r.db.Exec(query,
		initialUser.ID,
		initialUser.Username,
		initialUser.Password,
		initialUser.FullName,
		initialUser.Email,
		initialUser.IsActive,
		initialUser.RefreshToken,
		initialUser.CreatedAt,
		initialUser.CreatedBy,
		initialUser.UpdatedAt,
		initialUser.UpdatedBy,
	)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return ErrUserCodeDuplicate
			}
		}

		return fmt.Errorf("failed to save InitialUser: %w", err)
	}

	return nil
}
func (r *Repository) Update(initialUser InitialUser) error {
	query := `
		UPDATE initial_user SET 
			ID = ?,
			Username = ?,
			Password = ?,
			FullName = ?,
			Email = ?,
			IsActive = ?,
			RefreshToken = ?,
			CreatedAt = ?,
			CreatedBy = ?,
			UpdatedAt = ?,
			UpdatedBy = ?
		WHERE ID = ?
	`

	result, err := r.db.Exec(query,
		initialUser.ID,
		initialUser.Username,
		initialUser.Password,
		initialUser.FullName,
		initialUser.Email,
		initialUser.IsActive,
		initialUser.RefreshToken,
		initialUser.CreatedAt,
		initialUser.CreatedBy,
		initialUser.UpdatedAt,
		initialUser.UpdatedBy,
	)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return ErrUserCodeDuplicate
			}
		}

		return fmt.Errorf("failed to update InitialUser: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}

func (r *Repository) FindAll() ([]InitialUser, error) {
	query := `
		SELECT COALESCE(ID, ''), COALESCE(Username, ''), COALESCE(Password, ''), COALESCE(FullName, ''), COALESCE(Email, ''), IsActive, COALESCE(RefreshToken, ''), CreatedAt, COALESCE(CreatedBy, ''), COALESCE(UpdatedAt, CreatedAt), COALESCE(UpdatedBy, '')
		FROM initial_user
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query InitialUser: %w", err)
	}
	defer rows.Close()

	initialUserList := make([]InitialUser, 0)
	for rows.Next() {
		var initialUser InitialUser
		err := rows.Scan(
			&initialUser.ID,
			&initialUser.Username,
			&initialUser.Password,
			&initialUser.FullName,
			&initialUser.Email,
			&initialUser.IsActive,
			&initialUser.RefreshToken,
			&initialUser.CreatedAt,
			&initialUser.CreatedBy,
			&initialUser.UpdatedAt,
			&initialUser.UpdatedBy)
		if err != nil {
			return nil, fmt.Errorf("failed to scan InitialUser: %w", err)
		}
		initialUserList = append(initialUserList, initialUser)
	}

	return initialUserList, nil
}

func (r *Repository) FindById(id string) (InitialUser, error) {
	query := `
		SELECT COALESCE(ID, ''), COALESCE(Username, ''), COALESCE(Password, ''), COALESCE(FullName, ''), COALESCE(Email, ''), IsActive, COALESCE(RefreshToken, ''), CreatedAt, COALESCE(CreatedBy, ''), COALESCE(UpdatedAt, CreatedAt), COALESCE(UpdatedBy, '')
		FROM initial_user
		WHERE ID = ?
	`

	row := r.db.QueryRow(query, id)
	var initialUser InitialUser
	err := row.Scan(
		&initialUser.ID,
		&initialUser.Username,
		&initialUser.Password,
		&initialUser.FullName,
		&initialUser.Email,
		&initialUser.IsActive,
		&initialUser.RefreshToken,
		&initialUser.CreatedAt,
		&initialUser.CreatedBy,
		&initialUser.UpdatedAt,
		&initialUser.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return InitialUser{}, ErrUserNotFound
		}
		return InitialUser{}, fmt.Errorf("failed to scan InitialUser: %w", err)
	}
	return initialUser, nil
}

func (r *Repository) FindByUsername(username string) (InitialUser, error) {
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
	FROM initial_user A
	WHERE A.Username = ? AND A.IsActive = true
	`

	var u InitialUser
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
			return InitialUser{}, ErrUserNotFound
		}
		return InitialUser{}, fmt.Errorf("failed to find user: %w", err)
	}

	return u, nil
}

func (r *Repository) UpdateRefreshToken(userID, token string) error {
	query := `UPDATE initial_user SET RefreshToken = ? WHERE ID = ?`
	_, err := r.db.Exec(query, token, userID)
	if err != nil {
		return fmt.Errorf("failed to update refresh token: %w", err)
	}
	return nil
}
func (r *Repository) FindByRefreshToken(token string) (InitialUser, error) {
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
	FROM initial_user A
	WHERE A.RefreshToken = ? AND A.IsActive = true
	`

	var u InitialUser
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
			return InitialUser{}, ErrUserNotFound
		}
		return InitialUser{}, fmt.Errorf("failed to find user by refresh token: %w", err)
	}

	return u, nil
}

func (r *Repository) FindByCode(code string) (InitialUser, error) {
	query := `
		SELECT COALESCE(ID, ''), COALESCE(Username, ''), COALESCE(Password, ''), COALESCE(FullName, ''), COALESCE(Email, ''), IsActive, COALESCE(RefreshToken, ''), CreatedAt, COALESCE(CreatedBy, ''), COALESCE(UpdatedAt, CreatedAt), COALESCE(UpdatedBy, '')
		FROM initial_user
		WHERE Code = ?
	`

	row := r.db.QueryRow(query, code)
	var initialUser InitialUser
	err := row.Scan(
		&initialUser.ID,
		&initialUser.Username,
		&initialUser.Password,
		&initialUser.FullName,
		&initialUser.Email,
		&initialUser.IsActive,
		&initialUser.RefreshToken,
		&initialUser.CreatedAt,
		&initialUser.CreatedBy,
		&initialUser.UpdatedAt,
		&initialUser.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return InitialUser{}, ErrUserNotFound
		}
		return InitialUser{}, fmt.Errorf("failed to scan InitialUser: %w", err)
	}
	return initialUser, nil
}

func (r *Repository) SoftDelete(id string) error {
	query := `
		UPDATE initial_user SET IsActive = false
		WHERE ID = ?
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete InitialUser: %w", err)
	}
	return nil
}

func (r *Repository) Delete(id string) error {
	query := `
		DELETE FROM initial_user
		WHERE ID = ?
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete InitialUser: %w", err)
	}
	return nil
}
