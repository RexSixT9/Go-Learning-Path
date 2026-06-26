package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/rexsixt9/students-api/internal/types"

	"github.com/rexsixt9/students-api/internal/config"
	_ "modernc.org/sqlite"
)

type SQLiteStorage struct {
	db *sql.DB
}

func NewSQLiteStorage(cfg *config.Config) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite", cfg.StoragePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open SQLite database: %w", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		age INTEGER NOT NULL
	)`)
	if err != nil {
		return nil, fmt.Errorf("failed to create students table: %w", err)
	}

	return &SQLiteStorage{db: db}, nil
}

func (s *SQLiteStorage) CreateStudent(name string, email string, age int) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, fmt.Errorf("failed to insert student: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve last insert ID: %w", err)
	}

	return id, nil
}

func (s *SQLiteStorage) GetStudentByID(id int64) (types.Student, error) {
	stmt, err := s.db.Prepare("SELECT id, name, email, age FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	var student types.Student
	err = stmt.QueryRow(id).Scan(&student.ID, &student.Name, &student.Email, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("student with ID %d not found: %w", id, err)
		}
		return types.Student{}, fmt.Errorf("failed to query student: %w", err)
	}
	return student, nil
}

func (s *SQLiteStorage) GetStudents() ([]types.Student, error) {
	stmt, err := s.db.Prepare("SELECT id, name, email, age FROM students")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("failed to query students: %w", err)
	}
	defer rows.Close()

	var students []types.Student
	for rows.Next() {
		var student types.Student
		err := rows.Scan(&student.ID, &student.Name, &student.Email, &student.Age)
		if err != nil {
			return nil, fmt.Errorf("failed to scan student: %w", err)
		}
		students = append(students, student)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over students: %w", err)
	}
	return students, nil
}

func (s *SQLiteStorage) UpdateStudent(id int64, name string, email string, age int) error {
	stmt, err := s.db.Prepare("UPDATE students SET name = ?, email = ?, age = ? WHERE id = ?")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, email, age, id)
	if err != nil {
		return fmt.Errorf("failed to update student: %w", err)
	}

	return nil
}

func (s *SQLiteStorage) DeleteStudent(id int64) error {
	stmt, err := s.db.Prepare("DELETE FROM students WHERE id = ?")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("failed to delete student: %w", err)
	}
	return nil
}
