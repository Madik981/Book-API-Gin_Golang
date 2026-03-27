package models

import (
	"sort"
	"strings"
	"sync"
)

type Store struct {
	mu sync.RWMutex

	books      map[int]Book
	authors    map[int]Author
	categories map[int]Category

	nextBookID     int
	nextAuthorID   int
	nextCategoryID int
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

func NewStore() *Store {
	return &Store{
		books:          make(map[int]Book),
		authors:        make(map[int]Author),
		categories:     make(map[int]Category),
		nextBookID:     1,
		nextAuthorID:   1,
		nextCategoryID: 1,
	}
}

func (s *Store) ListBooks(filter BookFilter) PaginatedBooks {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.Limit < 1 {
		filter.Limit = 10
	}
	if filter.Limit > 100 {
		filter.Limit = 100
	}

	filtered := make([]Book, 0)
	titleNeedle := strings.ToLower(strings.TrimSpace(filter.Title))

	for _, b := range s.books {
		if filter.AuthorID > 0 && b.AuthorID != filter.AuthorID {
			continue
		}
		if filter.CategoryID > 0 && b.CategoryID != filter.CategoryID {
			continue
		}
		if titleNeedle != "" && !strings.Contains(strings.ToLower(b.Title), titleNeedle) {
			continue
		}
		filtered = append(filtered, b)
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].ID < filtered[j].ID
	})

	total := len(filtered)
	totalPages := 0
	if total > 0 {
		totalPages = (total + filter.Limit - 1) / filter.Limit
	}

	start := (filter.Page - 1) * filter.Limit
	if start > total {
		start = total
	}
	end := start + filter.Limit
	if end > total {
		end = total
	}

	pageData := filtered[start:end]

	return PaginatedBooks{
		Data:       pageData,
		Total:      total,
		Page:       filter.Page,
		Limit:      filter.Limit,
		TotalPages: totalPages,
	}
}

func (s *Store) GetBook(id int) (Book, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	b, ok := s.books[id]
	return b, ok
}

func (s *Store) CreateBook(input CreateBookInput) (Book, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.authors[input.AuthorID]; !ok {
		return Book{}, false
	}
	if _, ok := s.categories[input.CategoryID]; !ok {
		return Book{}, false
	}

	book := Book{
		ID:         s.nextBookID,
		Title:      strings.TrimSpace(input.Title),
		AuthorID:   input.AuthorID,
		CategoryID: input.CategoryID,
		Price:      input.Price,
	}
	s.books[book.ID] = book
	s.nextBookID++
	return book, true
}

func (s *Store) UpdateBook(id int, input UpdateBookInput) (Book, string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.books[id]; !ok {
		return Book{}, "not_found"
	}
	if _, ok := s.authors[input.AuthorID]; !ok {
		return Book{}, "author_not_found"
	}
	if _, ok := s.categories[input.CategoryID]; !ok {
		return Book{}, "category_not_found"
	}

	book := Book{
		ID:         id,
		Title:      strings.TrimSpace(input.Title),
		AuthorID:   input.AuthorID,
		CategoryID: input.CategoryID,
		Price:      input.Price,
	}
	s.books[id] = book
	return book, ""
}

func (s *Store) DeleteBook(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.books[id]; !ok {
		return false
	}
	delete(s.books, id)
	return true
}

func (s *Store) ListAuthors() []Author {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]Author, 0, len(s.authors))
	for _, a := range s.authors {
		out = append(out, a)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}

func (s *Store) CreateAuthor(input CreateAuthorInput) Author {
	s.mu.Lock()
	defer s.mu.Unlock()
	a := Author{ID: s.nextAuthorID, Name: strings.TrimSpace(input.Name)}
	s.authors[a.ID] = a
	s.nextAuthorID++
	return a
}

func (s *Store) ListCategories() []Category {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]Category, 0, len(s.categories))
	for _, c := range s.categories {
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}

func (s *Store) CreateCategory(input CreateCategoryInput) Category {
	s.mu.Lock()
	defer s.mu.Unlock()
	c := Category{ID: s.nextCategoryID, Name: strings.TrimSpace(input.Name)}
	s.categories[c.ID] = c
	s.nextCategoryID++
	return c
}
