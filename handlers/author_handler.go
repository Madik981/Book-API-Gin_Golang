package handlers

import (
	"net/http"

	"Book-API-Gin_Golang/models"
	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	store *models.Store
}

func NewAuthorHandler(store *models.Store) *AuthorHandler {
	return &AuthorHandler{store: store}
}

func (h *AuthorHandler) ListAuthors(c *gin.Context) {
	c.JSON(http.StatusOK, h.store.ListAuthors())
}

func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var input models.CreateAuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author := h.store.CreateAuthor(input)
	c.JSON(http.StatusCreated, author)
}
