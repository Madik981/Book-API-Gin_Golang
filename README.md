# Book API (Gin + Go)

This is a small REST API project for managing books, authors, and categories.
Data is stored in memory (`map` + `sync.RWMutex`), so everything resets after restart.

## What this project does

- Book endpoints:
  - `GET /books` (with pagination + filters)
  - `POST /books`
  - `GET /books/:id`
  - `PUT /books/:id`
  - `DELETE /books/:id`
- Author endpoints:
  - `GET /authors`
  - `POST /authors`
- Category endpoints:
  - `GET /categories`
  - `POST /categories`
- Validation:
  - required fields
  - `author_id > 0`
  - `category_id > 0`
  - `price >= 0.01`

## How to run

```bash
go mod tidy
go run .
```

Server runs on `:8080`.

## Books query params

For `GET /books` you can use:

- `page` (default: `1`)
- `limit` (default: `10`, max: `100`)
- `author_id`
- `category_id`
- `title` (partial match, case-insensitive)

Example URL:

`/books?category_id=1&page=1&limit=10&title=harry`

## JSON examples

### 1) Create author

**Endpoint:** `POST /authors`

**Request body**
```json
{
  "name": "J.K. Rowling"
}
```

**Response (201)**
```json
{
  "id": 1,
  "name": "J.K. Rowling"
}
```

### 2) Create category

**Endpoint:** `POST /categories`

**Request body**
```json
{
  "name": "Fantasy"
}
```

**Response (201)**
```json
{
  "id": 1,
  "name": "Fantasy"
}
```

### 3) Create book

**Endpoint:** `POST /books`

**Request body**
```json
{
  "title": "Harry Potter",
  "author_id": 1,
  "category_id": 1,
  "price": 19.99
}
```

**Response (201)**
```json
{
  "id": 1,
  "title": "Harry Potter",
  "author_id": 1,
  "category_id": 1,
  "price": 19.99
}
```

### 4) List books (with filters/pagination)

**Endpoint:** `GET /books?category_id=1&page=1&limit=10&title=harry`

**Response (200)**
```json
{
  "data": [
    {
      "id": 1,
      "title": "Harry Potter",
      "author_id": 1,
      "category_id": 1,
      "price": 19.99
    }
  ],
  "total": 1,
  "page": 1,
  "limit": 10,
  "total_pages": 1
}
```

### 5) Get book by ID

**Endpoint:** `GET /books/1`

**Response (200)**
```json
{
  "id": 1,
  "title": "Harry Potter",
  "author_id": 1,
  "category_id": 1,
  "price": 19.99
}
```

### 6) Update book

**Endpoint:** `PUT /books/1`

**Request body**
```json
{
  "title": "Harry Potter and the Chamber of Secrets",
  "author_id": 1,
  "category_id": 1,
  "price": 21.50
}
```

**Response (200)**
```json
{
  "id": 1,
  "title": "Harry Potter and the Chamber of Secrets",
  "author_id": 1,
  "category_id": 1,
  "price": 21.50
}
```

### 7) Delete book

**Endpoint:** `DELETE /books/1`

**Response (204)**
No body.

## Notes

- This project uses in-memory storage, so no database setup is needed.
- IDs are auto-incremented.
- If `author_id` or `category_id` does not exist, creating/updating a book returns a `400` error.
