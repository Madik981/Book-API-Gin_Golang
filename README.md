# Book API (Gin + Go)

This is a small REST API project for managing books, authors, and categories.
Data is stored in PostgreSQL via GORM.

## What this project does

- Authentication endpoints:
  - `POST /auth/register`
  - `POST /auth/login`
- Book endpoints:
  - `GET /books` (with pagination + filters)
  - `POST /books` (protected)
  - `GET /books/:id`
  - `PUT /books/:id` (protected)
  - `DELETE /books/:id` (protected)
- Author endpoints:
  - `GET /authors`
  - `POST /authors` (protected)
- Category endpoints:
  - `GET /categories`
  - `POST /categories` (protected)
- Validation:
  - required fields
  - `author_id > 0`
  - `category_id > 0`
  - `price >= 0.01`

## How to run

1) Prepare PostgreSQL (local or Docker) and create database, for example `book_api`.

2) Set environment variables (PowerShell example):

```bash
$env:DB_HOST="localhost"
$env:DB_PORT="5432"
$env:DB_USER="postgres"
$env:DB_PASSWORD="postgres"
$env:DB_NAME="book_api"
$env:DB_SSLMODE="disable"
$env:DB_TIMEZONE="UTC"
```

You can also use one connection string:

```bash
$env:DATABASE_URL="host=localhost user=postgres password=postgres dbname=book_api port=5432 sslmode=disable TimeZone=UTC"
```

3) Run app:

```bash
go mod tidy
go run .
```

Server runs on `:8080`.

## Authentication

Set JWT secret (recommended):

```bash
$env:JWT_SECRET="your-super-secret"
```

If `JWT_SECRET` is not set, app uses fallback secret `dev-secret-change-me`.

### Register user

**Endpoint:** `POST /auth/register`

**Request body**
```json
{
  "username": "admin",
  "password": "secret123"
}
```

### Login user

**Endpoint:** `POST /auth/login`

**Request body**
```json
{
  "username": "admin",
  "password": "secret123"
}
```

**Response (200)**
```json
{
  "token": "<jwt>",
  "token_type": "Bearer",
  "expires_in": 86400
}
```

Use token for protected endpoints:

`Authorization: Bearer <jwt>`

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

- GORM runs `AutoMigrate` on startup for `authors`, `categories`, `books`, `users`.
- IDs are auto-incremented.
- If `author_id` or `category_id` does not exist, creating/updating a book returns a `400` error.
