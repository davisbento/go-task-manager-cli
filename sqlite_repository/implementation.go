package sqlite_repository

import (
	"database/sql"
	"davisbento/go-task-manager-cli/repository"
	"log"
)

type SqliteRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *SqliteRepository {
	return &SqliteRepository{
		db: db,
	}
}

func (r *SqliteRepository) AddTask(description string) error {
	_, err := r.db.Exec("INSERT INTO tasks (description) VALUES (?)", description)

	if err != nil {
		log.Fatalf("Failed to add task: %v", err)
		return err
	}

	return nil
}

func (r *SqliteRepository) ListTasks(status string) []repository.Task {
	query := "SELECT id, description, completed FROM tasks"

	if status == "completed" {
		query += " WHERE completed = 1"
	}

	if status == "pending" {
		query += " WHERE completed = 0"
	}

	rows, err := r.db.Query(query)

	if err != nil {
		log.Fatalf("Failed to query tasks: %v", err)
		return nil
	}

	defer rows.Close()

	var tasks []repository.Task

	for rows.Next() {
		var id int
		var description string
		var completed bool

		if err := rows.Scan(&id, &description, &completed); err != nil {
			log.Fatalf("Failed to read row: %v", err)
			return nil
		}

		tasks = append(tasks, repository.Task{
			ID:          id,
			Description: description,
			Completed:   completed,
		})

	}

	return tasks
}

func (r *SqliteRepository) CompleteTask(id int) error {
	_, err := r.db.Exec("UPDATE tasks SET completed = 1 WHERE id = ?", id)

	if err != nil {
		log.Fatalf("Failed to mark task as completed: %v", err)
		return err
	}

	return nil
}

func (r *SqliteRepository) DeleteTask(id int) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", id)

	if err != nil {
		log.Fatalf("Failed to delete task: %v", err)
		return err
	}

	return nil
}
