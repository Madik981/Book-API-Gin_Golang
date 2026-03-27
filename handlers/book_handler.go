package handlers

import (
	"net/http"
	"strconv"

	"Book-API-Gin_Golang/models"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	store *models.Store
}

func NewBookHandler(store *models.Store) *BookHandler {
	return &BookHandler{store: store}
}

func (h *BookHandler) ListBooks(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	authorID, _ := strconv.Atoi(c.DefaultQuery("author_id", "0"))
	categoryID, _ := strconv.Atoi(c.DefaultQuery("category_id", "0"))

	filter := models.BookFilter{
		Title:      c.Query("title"),
		AuthorID:   authorID,
		CategoryID: categoryID,
		Page:       page,
		Limit:      limit,
	}

	result := h.store.ListBooks(filter)
	c.JSON(http.StatusOK, result)
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	book, ok := h.store.GetBook(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, ok := h.store.CreateBook(input)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "author_id or category_id does not exist"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, updateErr := h.store.UpdateBook(id, input)
	switch updateErr {
	case "":
		c.JSON(http.StatusOK, book)
	case "not_found":
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
	case "author_not_found":
		c.JSON(http.StatusBadRequest, gin.H{"error": "author_id does not exist"})
	case "category_not_found":
		c.JSON(http.StatusBadRequest, gin.H{"error": "category_id does not exist"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	ok := h.store.DeleteBook(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
