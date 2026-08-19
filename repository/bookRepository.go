package repository

import (
	"database/sql"

	"github.com/piipiets/go-library/model"
	"github.com/piipiets/go-library/model/response"
)

type Repository interface {
	CreateBook(book model.Book) (err error)
	GetAllBooks() (result []response.BookResponse, err error)
	GetBookById(id int) (result response.BookResponse, err error)
	DeleteBook(id int) (err error)
	UpdateBook(id int, book model.Book) (err error)
}

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) CreateBook(book model.Book) error {

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

func (r *BookRepository) GetAllBooks() ([]response.BookResponse, error) {
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
    `

	rows, err := r.db.Query(query)
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

func (r *BookRepository) GetBookById(id int) (response.BookResponse, error) {
	var book response.BookResponse

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
		WHERE b.id = $1
    `

	err := r.db.QueryRow(
		query,
		id,
	).Scan(
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
		return book, err
	}

	return book, nil
}

func (r *BookRepository) UpdateBook(id int, book model.Book) error {
	query := `
				UPDATE books
				SET
					title = $1,
					category_id = $2,
					description = $3,
					image_url = $4,
					release_year = $5,
					price = $6,
					total_page = $7,
					thickness = $8,
					modified_at = $9,
					modified_by = $10
				WHERE id = $11
			`

	result, err := r.db.Exec(
		query,
		book.Title,
		book.CategoryId,
		book.Description,
		book.ImageUrl,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.ModifiedAt,
		book.ModifiedBy,
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

func (r *BookRepository) DeleteBook(id int) error {
	query := `
		DELETE
		FROM books
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
