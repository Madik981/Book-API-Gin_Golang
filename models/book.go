package models

type Book struct {
	ID         int     `json:"id"`
	Title      string  `json:"title"`
	AuthorID   int     `json:"author_id"`
	CategoryID int     `json:"category_id"`
	Price      float64 `json:"price"`
}

type CreateBookInput struct {
	Title      string  `json:"title" binding:"required"`
	AuthorID   int     `json:"author_id" binding:"required,gt=0"`
	CategoryID int     `json:"category_id" binding:"required,gt=0"`
	Price      float64 `json:"price" binding:"required,gte=0.01"`
}

type UpdateBookInput struct {
	Title      string  `json:"title" binding:"required"`
	AuthorID   int     `json:"author_id" binding:"required,gt=0"`
	CategoryID int     `json:"category_id" binding:"required,gt=0"`
	Price      float64 `json:"price" binding:"required,gte=0.01"`
}
