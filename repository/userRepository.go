package repository

import (
	"database/sql"

	"github.com/piipiets/go-library/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByUsername(username string) (model.Users, error) {
	var user model.Users

	query := `
        SELECT id, username, password
        FROM users
        WHERE username = $1
    `

	err := r.db.QueryRow(
		query,
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
	)

	return user, err
}

func (r *UserRepository) CreateUser(user model.Users) (model.Users, error) {
	query := `
		INSERT INTO users (
			username,
			password,
			created_at,
			created_by,
			modified_at,
			modified_by
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		user.Username,
		user.Password,
		user.CreatedAt,
		user.CreatedBy,
		user.ModifiedAt,
		user.ModifiedBy,
	).Scan(&user.ID)

	if err != nil {
		return user, err
	}

	return user, nil
}
