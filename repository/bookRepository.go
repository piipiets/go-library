package repository

import (
	"database/sql"

	"github.com/piipiets/go-library/model"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) CreateBook(book model.Book, username string) error {

	query := `
        INSERT INTO books (
            title,
			category_id,
			description,
			image_url,
			release_year,
			price,
			total_page,
			thickness,
            created_at,
            created_by,
            modified_at,
            modified_by
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
    `

	_, err := r.db.Exec(
		query,
		book.Title,
		book.CategoryId,
		book.Description,
		book.ImageUrl,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.CreatedAt,
		book.CreatedBy,
		book.ModifiedAt,
		book.ModifiedBy,
	)

	if err != nil {
		return err
	}

	return nil
}
