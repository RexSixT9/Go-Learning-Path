# Notes CRUD API

A simple Go-based REST API for learning backend fundamentals.

This project is part of a Go learning path for beginners and newcomers who want to understand how to build a CRUD API with Gin and MongoDB.

## Features

- Create, read, update, and delete notes
- Health check endpoint
- MongoDB-backed data storage
- Beginner-friendly project structure

## Tech Stack

- Go
- Gin
- MongoDB
- godotenv

## Project Structure

- `cmd/api` - application entrypoint
- `internal/config` - environment loading
- `internal/db` - MongoDB connection
- `internal/handlers` - HTTP handlers
- `internal/repository` - database access logic
- `internal/models` - request and response models
- `routes` - route registration

## Requirements

- Go 1.26+
- MongoDB
- `.env` file with the required configuration

## Environment Variables

Create a `.env` file in the project root:

```env
MONGO_URI=mongodb://localhost:27017
MONGO_DB_NAME=notes_db
PORT=8080
```

## Run the Project

```bash
go mod tidy
go run ./cmd/api
```

## API Endpoints

- `GET /health` - service health check
- `POST /notes/` - create a note
- `GET /notes/` - list all notes
- `GET /notes/:id` - get a note by ID
- `PUT /notes/:id` - update a note
- `DELETE /notes/:id` - delete a note

## Contribution

Contributions are welcome.

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test your changes
5. Open a pull request

Please keep changes small, clear, and beginner-friendly when possible.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
