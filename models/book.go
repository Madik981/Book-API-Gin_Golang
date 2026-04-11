package models

type Book struct {
	ID         int      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title      string   `json:"title" gorm:"size:255;not null"`
	AuthorID   int      `json:"author_id" gorm:"not null;index"`
	CategoryID int      `json:"category_id" gorm:"not null;index"`
	Price      float64  `json:"price" gorm:"not null"`
	Author     Author   `json:"-" gorm:"foreignKey:AuthorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Category   Category `json:"-" gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
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
