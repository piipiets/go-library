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

# API cURL Examples

### 1. Create User
```
curl -X POST https://go-library-production.up.railway.app/user/add ^
  -H "Content-Type: application/json" ^
  -d "{\"username\":\"pipit\",\"password\":\"123456\"}"
```

### Login
``` 
curl -X POST https://go-library-production.up.railway.app/login ^
  -H "Content-Type: application/json" ^
  -d "{\"username\":\"pipit\",\"password\":\"123456\"}"
  ```

### Create Category
```
curl -X POST https://go-library-production.up.railway.app/api/categories ^
  -H "Content-Type: application/json" ^
  -H "Authorization: Bearer <JWT_TOKEN>" ^
  -d "{\"name\":\"Programming\"}"
```

### Get All Categories
```
curl -X GET https://go-library-production.up.railway.app/api/categories ^
  -H "Authorization: Bearer <JWT_TOKEN>"
```

### Get Category By ID
```
curl -X GET https://go-library-production.up.railway.app/api/categories/1 ^
  -H "Authorization: Bearer <JWT_TOKEN>"
```