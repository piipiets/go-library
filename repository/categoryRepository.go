package repository

import (
	"database/sql"
	"time"

	"github.com/piipiets/go-library/model"
	"github.com/piipiets/go-library/model/request"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) Create(categoryRequest request.CategoryRequest, username string) (model.Categories, error) {
	now := time.Now()

	var category model.Categories

	category.Name = categoryRequest.Name
	category.CreatedAt = now
	category.CreatedBy = username
	category.ModifiedAt = now
	category.ModifiedBy = username

	query := `
        INSERT INTO categories (
            name,
            created_at,
            created_by,
            modified_at,
            modified_by
        )
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `

	err := r.db.QueryRow(
		query,
		category.Name,
		category.CreatedAt,
		category.CreatedBy,
		category.ModifiedAt,
		category.ModifiedBy,
	).Scan(&category.ID)

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *CategoryRepository) GetAllCategory() ([]model.Categories, error) {
	var categories []model.Categories

	query := `
        SELECT
            id,
			name,
            created_at,
            created_by,
            modified_at,
            modified_by
		FROM categories
        ORDER BY id
    `

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category model.Categories

		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.CreatedAt,
			&category.CreatedBy,
			&category.ModifiedAt,
			&category.ModifiedBy,
		)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
