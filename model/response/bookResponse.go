package response

type BookResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	ReleaseYear int    `json:"release_year"`
	Price       int64  `json:"price"`
	TotalPage   int    `json:"total_page"`
	Thickness   string `json:"thickness"`
	Category    string
}
