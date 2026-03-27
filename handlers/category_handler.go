package handlers

import (
	"net/http"

	"Book-API-Gin_Golang/models"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	store *models.Store
}

func NewCategoryHandler(store *models.Store) *CategoryHandler {
	return &CategoryHandler{store: store}
}

func (h *CategoryHandler) ListCategories(c *gin.Context) {
	c.JSON(http.StatusOK, h.store.ListCategories())
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var input models.CreateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := h.store.CreateCategory(input)
	c.JSON(http.StatusCreated, category)
}
