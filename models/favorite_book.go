package models

import "time"

type FavoriteBook struct {
	UserID    int       `gorm:"primaryKey;not null;index"`
	BookID    int       `gorm:"primaryKey;not null;index"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Book Book `gorm:"foreignKey:BookID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (FavoriteBook) TableName() string {
	return "favorite_books"
}
