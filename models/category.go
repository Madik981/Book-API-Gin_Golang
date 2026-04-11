package models

type Category struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"size:255;not null"`
}

type CreateCategoryInput struct {
	Name string `json:"name" binding:"required"`
}
