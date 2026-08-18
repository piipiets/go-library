package repository

import (
	"database/sql"
	"time"

	"github.com/piipiets/go-library/model"
	"github.com/piipiets/go-library/model/request"
	"github.com/piipiets/go-library/model/response"
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
			name
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

func (r *CategoryRepository) GetCategoryById(id int) (model.Categories, error) {
	var category model.Categories

	query := `
		SELECT id, name
		FROM categories
		WHERE id = $1
	`

	err := r.db.QueryRow(
		query,
		id,
	).Scan(
		&category.ID,
		&category.Name,
	)

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *CategoryRepository) DeleteCategory(id int) error {
	query := `
		DELETE
		FROM categories
		WHERE id = $1
	`

	result, err := r.db.Exec(query, id)

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *CategoryRepository) GetBooksByCategory(id int) ([]response.BookResponse, error) {
	var books []response.BookResponse

	query := `
        SELECT
            b.title,
			b.description,
			b.image_url,
			b.release_year,
			b.price,
			b.total_page,
			b.thickness,
			c.name as category
		FROM books b
		JOIN categories c ON b.category_id = c.id
		WHERE b.category_id = $1
		ORDER BY b.id
    `

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book response.BookResponse

		err := rows.Scan(
			&book.Title,
			&book.Description,
			&book.ImageUrl,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&book.Category,
		)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *CategoryRepository) UpdateCategory(id int, category model.Categories) error {
	query := `
		UPDATE categories
		SET name = $1,
		    modified_at = $2,
		    modified_by = $3
		WHERE id = $4
		RETURNING id, name
	`

	result, err := r.db.Exec(
		query,
		category.Name,
		category.ModifiedAt,
		category.ModifiedBy,
		id,
	)

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
