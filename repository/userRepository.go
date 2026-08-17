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
