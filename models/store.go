package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

type BookFilter struct {
	Title      string
	AuthorID   int
	CategoryID int
	Page       int
	Limit      int
}

type PaginatedBooks struct {
	Data       []Book `json:"data"`
	Total      int    `json:"total"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	TotalPages int    `json:"total_pages"`
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) ListBooks(filter BookFilter) PaginatedBooks {
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.Limit < 1 {
		filter.Limit = 10
	}
	if filter.Limit > 100 {
		filter.Limit = 100
	}

	query := s.applyBookFilters(s.db.Model(&Book{}), filter)

	var total64 int64
	if err := query.Count(&total64).Error; err != nil {
		total64 = 0
	}
	total := int(total64)
	totalPages := 0
	if total > 0 {
		totalPages = (total + filter.Limit - 1) / filter.Limit
	}

	start := (filter.Page - 1) * filter.Limit
	if start > total {
		start = total
	}

	pageData := make([]Book, 0)
	if err := s.applyBookFilters(s.db, filter).
		Order("id ASC").
		Offset(start).
		Limit(filter.Limit).
		Find(&pageData).Error; err != nil {
		pageData = []Book{}
	}

	return PaginatedBooks{
		Data:       pageData,
		Total:      total,
		Page:       filter.Page,
		Limit:      filter.Limit,
		TotalPages: totalPages,
	}
}

func (s *Store) GetBook(id int) (Book, bool) {
	var book Book
	if err := s.db.First(&book, id).Error; err != nil {
		return Book{}, false
	}
	return book, true
}

func (s *Store) CreateBook(input CreateBookInput) (Book, bool) {
	if !s.recordExists(&Author{}, input.AuthorID) {
		return Book{}, false
	}
	if !s.recordExists(&Category{}, input.CategoryID) {
		return Book{}, false
	}

	book := Book{
		Title:      strings.TrimSpace(input.Title),
		AuthorID:   input.AuthorID,
		CategoryID: input.CategoryID,
		Price:      input.Price,
	}

	if err := s.db.Create(&book).Error; err != nil {
		return Book{}, false
	}

	return book, true
}

func (s *Store) UpdateBook(id int, input UpdateBookInput) (Book, string) {
	var book Book
	if err := s.db.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Book{}, "not_found"
		}
		return Book{}, "not_found"
	}

	if !s.recordExists(&Author{}, input.AuthorID) {
		return Book{}, "author_not_found"
	}
	if !s.recordExists(&Category{}, input.CategoryID) {
		return Book{}, "category_not_found"
	}

	book.Title = strings.TrimSpace(input.Title)
	book.AuthorID = input.AuthorID
	book.CategoryID = input.CategoryID
	book.Price = input.Price

	if err := s.db.Save(&book).Error; err != nil {
		return Book{}, "not_found"
	}

	return book, ""
}

func (s *Store) DeleteBook(id int) bool {
	result := s.db.Delete(&Book{}, id)
	if result.Error != nil {
		return false
	}
	return result.RowsAffected > 0
}

func (s *Store) ListAuthors() []Author {
	out := make([]Author, 0)
	if err := s.db.Order("id ASC").Find(&out).Error; err != nil {
		return []Author{}
	}
	return out
}

func (s *Store) CreateAuthor(input CreateAuthorInput) Author {
	a := Author{Name: strings.TrimSpace(input.Name)}
	_ = s.db.Create(&a).Error
	return a
}

func (s *Store) ListCategories() []Category {
	out := make([]Category, 0)
	if err := s.db.Order("id ASC").Find(&out).Error; err != nil {
		return []Category{}
	}
	return out
}

func (s *Store) CreateCategory(input CreateCategoryInput) Category {
	c := Category{Name: strings.TrimSpace(input.Name)}
	_ = s.db.Create(&c).Error
	return c
}

func (s *Store) CreateUser(input CreateUserInput) (User, bool) {
	username := strings.TrimSpace(input.Username)
	normalizedUsername := strings.ToLower(username)

	var existing int64
	if err := s.db.Model(&User{}).
		Where("LOWER(username) = ?", normalizedUsername).
		Count(&existing).Error; err != nil {
		return User{}, false
	}
	if existing > 0 {
		return User{}, false
	}

	user := User{
		Username: username,
		Password: input.Password,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return User{}, false
	}

	return user, true
}

func (s *Store) AuthenticateUser(input LoginInput) (User, bool) {
	normalizedUsername := strings.ToLower(strings.TrimSpace(input.Username))

	var user User
	if err := s.db.Where("LOWER(username) = ?", normalizedUsername).First(&user).Error; err != nil {
		return User{}, false
	}

	if user.Password != input.Password {
		return User{}, false
	}

	return user, true
}

func (s *Store) applyBookFilters(query *gorm.DB, filter BookFilter) *gorm.DB {
	filtered := query.Model(&Book{})

	if filter.AuthorID > 0 {
		filtered = filtered.Where("author_id = ?", filter.AuthorID)
	}
	if filter.CategoryID > 0 {
		filtered = filtered.Where("category_id = ?", filter.CategoryID)
	}

	titleNeedle := strings.TrimSpace(filter.Title)
	if titleNeedle != "" {
		filtered = filtered.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(titleNeedle)+"%")
	}

	return filtered
}

func (s *Store) recordExists(model interface{}, id int) bool {
	var count int64
	err := s.db.Model(model).Where("id = ?", id).Count(&count).Error
	return err == nil && count > 0
}
