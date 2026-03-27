package models

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateAuthorInput struct {
	Name string `json:"name" binding:"required"`
}
