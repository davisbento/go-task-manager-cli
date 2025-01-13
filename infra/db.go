package infra

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "tasks.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create the tasks table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			description TEXT NOT NULL,
			completed BOOLEAN NOT NULL DEFAULT 0
		)
	`)

	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	log.Println("Database initialized")

	return db
}
