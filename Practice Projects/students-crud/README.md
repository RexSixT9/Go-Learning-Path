# Students API

A small Go CRUD REST API for learning how a layered backend works end to end.

This project shows how to wire together:

- HTTP routing with the standard library
- Request validation
- JSON response helpers
- SQLite persistence
- Configuration loading from YAML
- Graceful shutdown

## What You Can Learn Here

This repo is a good study path if you want to understand how a Go REST API is structured.

You can follow the request flow in this order:

1. `cmd/students-api/main.go` starts the app and registers routes.
2. `internal/http/handlers/student/student.go` handles HTTP requests.
3. `internal/storage/storage.go` defines the storage contract.
4. `internal/storage/sqlite/sqlite.go` implements that contract using SQLite.
5. `internal/types/types.go` defines the `Student` model.
6. `internal/utils/response/response.go` formats JSON responses.
7. `internal/config/config.go` loads config from `config/local.yaml`.

## Project Structure

```text
students-api/
├── cmd/
│   └── students-api/
│       └── main.go              # app entrypoint and route registration
├── config/
│   └── local.yaml               # local runtime config
├── internal/
│   ├── config/                  # config loading and parsing
│   ├── http/handlers/student/   # CRUD HTTP handlers
│   ├── storage/                 # storage interface
│   ├── storage/sqlite/          # SQLite implementation
│   ├── types/                   # domain models
│   └── utils/response/          # JSON response helpers
├── storage/
│   └── storage.db               # local SQLite database file
├── LICENSE
└── README.md
```

If you are learning Go, this tree is intentionally split by responsibility:

- `cmd/` is the application start point
- `internal/` keeps implementation details private to the repo
- `config/` holds runtime settings
- `storage/` is the local database file used for development

## How The API Works

The application is built in layers:

- The `main` package loads config, creates the SQLite storage, and registers routes.
- The handler layer validates input and turns HTTP requests into storage calls.
- The storage layer hides the database details behind an interface.
- The SQLite implementation executes SQL queries.

That separation is useful because you can replace SQLite later without rewriting the handlers.

## Entry Point

See [`cmd/students-api/main.go`](./cmd/students-api/main.go).

This file does three important things:

1. Loads config with `config.MustLoad()`.
2. Creates the SQLite storage with `sqlite.NewSQLiteStorage(cfg)`.
3. Registers HTTP routes on `http.NewServeMux()`.

It also handles graceful shutdown with `os/signal` and `server.Shutdown(...)`.

## Routes

These are the routes currently exposed by the app:

| Method | Path | Purpose |
| --- | --- | --- |
| `POST` | `/api/students` | Create a student |
| `GET` | `/api/students` | List all students |
| `GET` | `/api/students/{id}` | Get one student by ID |
| `PUT` | `/api/students/{id}` | Update a student |
| `DELETE` | `/api/students/{id}` | Delete a student |

Route registration happens in [`cmd/students-api/main.go`](./cmd/students-api/main.go).

## Student Model

See [`internal/types/types.go`](./internal/types/types.go).

```go
type Student struct {
	ID    int64  `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"required"`
}
```

Notes:

- `ID` is returned in responses.
- `Name`, `Email`, and `Age` are validated on create and update.
- `Email` must be a valid email address.

## Request Validation

Validation happens in the handler layer using `go-playground/validator`.

Relevant code:

- [`internal/http/handlers/student/student.go`](./internal/http/handlers/student/student.go)
- [`internal/types/types.go`](./internal/types/types.go)
- [`internal/utils/response/response.go`](./internal/utils/response/response.go)

Validation behavior:

- Empty request body returns `400 Bad Request`.
- Invalid JSON returns `400 Bad Request`.
- Missing required fields return a formatted validation error.

## Storage Contract

See [`internal/storage/storage.go`](./internal/storage/storage.go).

The interface keeps the handler code independent from the database engine:

```go
type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentByID(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
	UpdateStudent(id int64, name string, email string, age int) error
	DeleteStudent(id int64) error
}
```

This is a classic Go pattern:

- define behavior as an interface
- implement it in a concrete package
- inject it into handlers

## SQLite Implementation

See [`internal/storage/sqlite/sqlite.go`](./internal/storage/sqlite/sqlite.go).

This package:

- opens the SQLite database
- creates the `students` table if it does not exist
- runs SQL for create, read, update, and delete operations

Table schema:

```sql
CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	age INTEGER NOT NULL
)
```

Important behavior:

- `email` is unique
- `id` is auto-incremented
- queries use prepared statements

## HTTP Handlers

See [`internal/http/handlers/student/student.go`](./internal/http/handlers/student/student.go).

The handler file contains five functions:

- `New(storage)` creates a student
- `GetByID(storage)` fetches a student by ID
- `List(storage)` returns all students
- `Update(storage)` updates a student
- `Delete(storage)` removes a student

Each handler follows the same pattern:

1. Read input from the request.
2. Validate or parse it.
3. Call the storage layer.
4. Return JSON using the response helper.

## Response Format

See [`internal/utils/response/response.go`](./internal/utils/response/response.go).

Success responses are written as JSON with the correct HTTP status code.

Error responses are normalized into:

```json
{
  "status": "Error",
  "error": "message"
}
```

## Configuration

See [`internal/config/config.go`](./internal/config/config.go) and [`config/local.yaml`](./config/local.yaml).

The app expects a config file path from either:

- `CONFIG_PATH` environment variable
- `--config` command-line flag

Example config:

```yaml
env: "dev"
storage_path: "storage/storage.db"
http_server:
  address: "localhost:8080"
```

## Run Locally

1. Make sure Go is installed.
2. Set the config path.
3. Start the app.

Example for PowerShell:

```powershell
$env:CONFIG_PATH="config/local.yaml"
go run ./cmd/students-api
```

If you prefer a flag:

```powershell
go run ./cmd/students-api --config config/local.yaml
```

The server should start on `localhost:8080` with the default config.

## Example Requests

### Create a student

```bash
curl -X POST http://localhost:8080/api/students ^
  -H "Content-Type: application/json" ^
  -d "{\"name\":\"Asha\",\"email\":\"asha@example.com\",\"age\":21}"
```

### List students

```bash
curl http://localhost:8080/api/students
```

### Get a student

```bash
curl http://localhost:8080/api/students/1
```

### Update a student

```bash
curl -X PUT http://localhost:8080/api/students/1 ^
  -H "Content-Type: application/json" ^
  -d "{\"name\":\"Asha Patel\",\"email\":\"asha@example.com\",\"age\":22}"
```

### Delete a student

```bash
curl -X DELETE http://localhost:8080/api/students/1
```

## Learning Roadmap

If you are using this repo as part of a Go learning path, study it in this order:

1. `main.go`
2. `config.go`
3. `types.go`
4. `storage.go`
5. `sqlite.go`
6. `student.go`
7. `response.go`

Suggested questions to answer while reading:

- Where does the program start?
- How is config loaded?
- Why is the storage layer behind an interface?
- Where is validation performed?
- How does the handler know which student ID to use?
- How are JSON responses written?

## Possible Improvements

These are good next exercises for practice:

- add `PATCH` support for partial updates
- add unit tests for handlers and storage
- handle duplicate email errors with `409 Conflict`
- return `404 Not Found` when a student does not exist
- add pagination or filtering to `GET /api/students`
- add a Makefile or task runner for common commands

## Contributing

This repo can be used as a learning sandbox, so contributions are welcome if they improve clarity or behavior.

Good contribution ideas:

- fix a bug in validation or routing
- improve error handling
- add tests for a handler or storage method
- document a package more clearly
- add a new feature that fits the current CRUD flow

Suggested contribution flow:

1. Create a branch for your change.
2. Make one focused improvement.
3. Run the app and verify the behavior manually.
4. Add or update tests if possible.
5. Open a pull request with a short explanation of what changed and why.

If you are contributing as part of a learning path, it helps to keep changes small and explain the reasoning in the commit message or PR description.

## Before You Push

A simple pre-push checklist for this repo:

- review the README and make sure the language still matches the code
- confirm the route list matches `cmd/students-api/main.go`
- confirm the model fields match `internal/types/types.go`
- confirm the SQLite schema still matches `internal/storage/sqlite/sqlite.go`
- run formatting and tests if you changed code

Suggested commands:

```powershell
gofmt -w .\cmd\students-api\main.go .\internal\config\config.go .\internal\http\handlers\student\student.go .\internal\storage\storage.go .\internal\storage\sqlite\sqlite.go .\internal\types\types.go .\internal\utils\response\response.go
go test ./...
```

## Notes

- The project uses the standard library router `http.NewServeMux`.
- SQLite database file is stored locally in `storage/storage.db`.
- The repo is intentionally small so each layer is easy to understand.
- The project is licensed under MIT. See [LICENSE](./LICENSE).
