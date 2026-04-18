package handlers

import (
	"Book-API-Gin_Golang/models"
	"os"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, store *models.Store) {
	bookHandler := NewBookHandler(store)
	authorHandler := NewAuthorHandler(store)
	categoryHandler := NewCategoryHandler(store)
	jwtSecret := resolveJWTSecret()
	authHandler := NewAuthHandler(store, jwtSecret)

	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)

	r.GET("/books", bookHandler.ListBooks)

	r.GET("/authors", authorHandler.ListAuthors)

	r.GET("/categories", categoryHandler.ListCategories)

	protected := r.Group("/")
	protected.Use(AuthMiddleware(jwtSecret))
	{
		protected.GET("/books/favorites", bookHandler.ListFavoriteBooks)
		protected.PUT("/books/:id/favorites", bookHandler.AddFavoriteBook)
		protected.DELETE("/books/:id/favorites", bookHandler.RemoveFavoriteBook)

		protected.POST("/books", bookHandler.CreateBook)
		protected.PUT("/books/:id", bookHandler.UpdateBook)
		protected.DELETE("/books/:id", bookHandler.DeleteBook)

		protected.POST("/authors", authorHandler.CreateAuthor)
		protected.POST("/categories", categoryHandler.CreateCategory)
	}

	r.GET("/books/:id", bookHandler.GetBookByID)
}

func resolveJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "dev-secret-change-me"
	}
	return []byte(secret)
}
