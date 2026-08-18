package request

type BookRequest struct {
	Title       string `json:"title"`
	CategoryId  int    `json:"category_id"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	ReleaseYear int    `json:"release_year"`
	Price       int64  `json:"price"`
	TotalPage   int    `json:"total_page"`
}
