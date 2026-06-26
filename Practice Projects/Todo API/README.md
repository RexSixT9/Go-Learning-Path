# Todo API

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

Minimal Go REST API for user auth and per-user todo management.

## Features

- Register and log in users with JWT authentication
- Create, list, update, and delete todos
- Keep todos scoped to the authenticated user

## Stack

- [Gin](https://github.com/gin-gonic/gin)
- [pgx](https://github.com/jackc/pgx)
- [JWT](https://github.com/golang-jwt/jwt)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)

## Quick Start

### 1. Configure environment

Copy `.env.example` to `.env` and set your values:

```env
PORT=8080
DATABASE_URL=postgresql://user:password@localhost/todo_db
JWT_SECRET=your_jwt_secret_key
```

### 2. Run migrations

```powershell
.\scripts\migrate.ps1 up
```

### 3. Start the server

```powershell
go run ./cmd/api
```

## Endpoints

### Public

- `GET /` - health check
- `POST /auth/register` - create user
- `POST /auth/login` - get JWT token

### Protected

Send `Authorization: Bearer <token>` with all requests below:

- `POST /todos`
- `GET /todos`
- `GET /todos/:id`
- `PUT /todos/:id`
- `DELETE /todos/:id`
- `GET /protected`

## Example

```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}'
```

## Notes

- Passwords are stored as bcrypt hashes.
- JWTs expire after 24 hours.
- Todos can only be accessed by the user who created them.

## License

This project is licensed under the [MIT License](LICENSE).
