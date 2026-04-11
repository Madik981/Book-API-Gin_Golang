package models

type Author struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"size:255;not null"`
}

type CreateAuthorInput struct {
	Name string `json:"name" binding:"required"`
}
