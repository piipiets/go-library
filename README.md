# Go Library REST API

A simple Library REST API built with Go, Gin, and PostgreSQL.

This project is created as a learning project to practice backend development with Go, including REST API development, layered architecture, JWT authentication, middleware, PostgreSQL integration, and database migration.

## Live API

The API is deployed on Railway:

**Base URL:**  
https://go-library-production.up.railway.app

---

## Features

- RESTful API
- JWT Authentication
- User management
- Category management
- PostgreSQL database
- Repository-Service-Handler architecture
- JWT authentication middleware
- Viper configuration management

---

## Tech Stack

- **Go**
- **Gin** - HTTP web framework
- **PostgreSQL** - Database
- **lib/pq** - PostgreSQL driver
- **JWT** - Authentication
- **Viper** - Configuration management
- **sql-migrate** - Database migration

---

## API Endpoints

### Authentication & User

| Method | Endpoint | Auth | Description |
|---|---|---|---|
| `POST` | `/login` | ❌ | Login and get JWT token |
| `POST` | `/user/add` | ❌ | Create a new user |

### Categories

| Method | Endpoint | Auth | Description |
|---|---|---|---|
| `POST` | `/api/categories` | ✅ | Create a category |
| `GET` | `/api/categories` | ✅ | Get all categories |
| `GET` | `/api/categories/:id` | ✅ | Get category by ID |
| `PUT` | `/api/categories/:id` | ✅ | Update category |
| `DELETE` | `/api/categories/:id` | ✅ | Delete category |
| `GET` | `/api/categories/:id/books` | ✅ | Get all books by category |

### Books
| Method   | Endpoint         | Auth | Description    |
| -------- | ---------------- | ---- | -------------- |
| `POST`   | `/api/books`     | ✅    | Create a book  |
| `GET`    | `/api/books`     | ✅    | Get all books  |
| `GET`    | `/api/books/:id` | ✅    | Get book by ID |
| `PUT`    | `/api/books/:id` | ✅    | Update a book  |
| `DELETE` | `/api/books/:id` | ✅    | Delete a book  |

## Example Authorization Header

```http
Authorization: Bearer <JWT_TOKEN>
```

## Environment Variables

Set these environment variables before running the app:

```env
DATABASE_URL=postgresql://postgres:password@host:5432/railway
DB_ENGINE=postgres
```

