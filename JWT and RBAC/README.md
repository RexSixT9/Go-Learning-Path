# JWT and RBAC Project

This folder contains a Go API project focused on authentication and authorization.
It is a practical example for learning how JWT-based login and role-based access control work in a real Go service.

## What This Project Covers

- User authentication with JWT
- Role-based access control using middleware
- Protected and public API routes
- MongoDB-backed user data access
- Modular project structure with `cmd` and `internal`

## Main Project Files

- [`cmd/api/main.go`](./cmd/api/main.go) - application entry point
- [`internal/app/app.go`](./internal/app/app.go) - application wiring and startup logic
- [`internal/httpserver/router.go`](./internal/httpserver/router.go) - route setup
- [`internal/httpserver/health.go`](./internal/httpserver/health.go) - health check endpoint
- [`internal/auth/jwt.go`](./internal/auth/jwt.go) - JWT creation and verification helpers
- [`internal/middleware/auth.go`](./internal/middleware/auth.go) - authentication middleware
- [`internal/middleware/roles.go`](./internal/middleware/roles.go) - role-based access middleware
- [`internal/user/handler.go`](./internal/user/handler.go) - user request handlers
- [`internal/user/service.go`](./internal/user/service.go) - user business logic
- [`internal/user/repo.go`](./internal/user/repo.go) - user data access layer
- [`internal/user/model.go`](./internal/user/model.go) - user data model
- [`internal/config/config.go`](./internal/config/config.go) - environment and app config
- [`internal/db/mongo.go`](./internal/db/mongo.go) - MongoDB connection setup

## Important Folders

- `cmd/`
  - Contains the runnable API entry point
- `internal/app/`
  - Bootstraps dependencies and starts the app
- `internal/auth/`
  - Handles JWT-related helper functions
- `internal/httpserver/`
  - Holds health and routing setup
- `internal/middleware/`
  - Contains auth and role enforcement middleware
- `internal/user/`
  - Contains user handlers, models, services, and repository code
- `internal/config/`
  - Loads configuration values
- `internal/db/`
  - Manages database connection setup

## Example Learning Path

If you want to study this project in order, use this flow:

1. Start with [`cmd/api/main.go`](./cmd/api/main.go)
2. Read [`internal/app/app.go`](./internal/app/app.go) to see how the app is assembled
3. Check [`internal/config/config.go`](./internal/config/config.go) for environment values
4. Study [`internal/db/mongo.go`](./internal/db/mongo.go) for database setup
5. Review [`internal/user/model.go`](./internal/user/model.go), [`internal/user/repo.go`](./internal/user/repo.go), and [`internal/user/service.go`](./internal/user/service.go)
6. Learn how [`internal/auth/jwt.go`](./internal/auth/jwt.go) creates and validates tokens
7. Understand [`internal/middleware/auth.go`](./internal/middleware/auth.go) and [`internal/middleware/roles.go`](./internal/middleware/roles.go)
8. Finish with [`internal/httpserver/router.go`](./internal/httpserver/router.go) to see the full route wiring

## Files You May See In Practice

- `.env`
  - Local environment configuration
- `.env.example`
  - Sample environment file
- `.air.toml`
  - Air live-reload configuration
- `go.mod`
  - Module definition and dependencies
- `go.sum`
  - Dependency checksums

## Why This Project Is Useful

- It shows how authentication fits into a Go API
- It demonstrates middleware-based access control
- It keeps code organized in a production-style folder layout
- It gives a strong base for building secured APIs

## Notes

- This project is separate from the main Go learning roadmap in the repository root
- Use the root [README](../README.md) for the full topic-by-topic Go learning path

