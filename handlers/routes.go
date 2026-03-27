package handlers

import (
	"Book-API-Gin_Golang/models"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, store *models.Store) {
	bookHandler := NewBookHandler(store)
	authorHandler := NewAuthorHandler(store)
	categoryHandler := NewCategoryHandler(store)

	r.GET("/books", bookHandler.ListBooks)
	r.POST("/books", bookHandler.CreateBook)
	r.GET("/books/:id", bookHandler.GetBookByID)
	r.PUT("/books/:id", bookHandler.UpdateBook)
	r.DELETE("/books/:id", bookHandler.DeleteBook)

	r.GET("/authors", authorHandler.ListAuthors)
	r.POST("/authors", authorHandler.CreateAuthor)

	r.GET("/categories", categoryHandler.ListCategories)
	r.POST("/categories", categoryHandler.CreateCategory)
}
