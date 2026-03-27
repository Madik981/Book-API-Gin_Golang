package models

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateCategoryInput struct {
	Name string `json:"name" binding:"required"`
}
