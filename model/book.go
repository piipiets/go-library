package model

import "time"

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	CategoryId  int       `json:"category_id"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	ReleaseYear int       `json:"release_year"`
	Price       int64     `json:"price"`
	TotalPage   int       `json:"total_page"`
	Thickness   string    `json:"thickness"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	ModifiedAt  time.Time `json:"modified_at"`
	ModifiedBy  string    `json:"modified_by"`
}
