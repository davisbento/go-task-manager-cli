package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func initDB() *sql.DB {
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

func addTask(db *sql.DB, description string) {
	_, err := db.Exec("INSERT INTO tasks (description) VALUES (?)", description)

	if err != nil {
		log.Fatalf("Failed to add task: %v", err)
	}

	fmt.Println("Task added successfully!")
}

func listTasks(db *sql.DB, status string) {
	query := "SELECT id, description, completed FROM tasks"

	if status == "completed" {
		query += " WHERE completed = 1"
	}

	if status == "pending" {
		query += " WHERE completed = 0"
	}

	rows, err := db.Query(query)

	if err != nil {
		log.Fatalf("Failed to query tasks: %v", err)
	}

	defer rows.Close()

	// add a few empty lines
	fmt.Println()

	fmt.Println("========== Tasks ==========")

	for rows.Next() {
		var id int
		var description string
		var completed bool
		var status string

		if err := rows.Scan(&id, &description, &completed); err != nil {
			log.Fatalf("Failed to read row: %v", err)
		}

		if completed {
			status = "Completed"
		} else {
			status = "Pending"
		}

		fmt.Printf("[%d] %s - %s\n", id, description, status)
	}

	fmt.Println("================================")
	fmt.Println()
}

func completeTask(db *sql.DB, id int) {
	_, err := db.Exec("UPDATE tasks SET completed = 1 WHERE id = ?", id)
	if err != nil {
		log.Fatalf("Failed to mark task as completed: %v", err)
	}
	fmt.Println("Task marked as completed!")
}

func main() {
	db := initDB()

	defer db.Close()

	for {
		// Display menu
		fmt.Println("\nTask Manager")
		fmt.Println("======================")
		fmt.Println("1. Add a Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. Delete a Task")
		fmt.Println("q. Exit")
		fmt.Print("Choose an option: ")

		// Read user input
		var input string
		fmt.Scanln(&input)

		// Handle input
		switch strings.TrimSpace(input) {
		case "1":
			fmt.Print("Enter task description: ")

			reader := bufio.NewReader(os.Stdin)

			description, _ := reader.ReadString('\n')

			description = strings.TrimSpace(description)

			addTask(db, description)

		case "2":
			listTasks(db, "")

		case "3":
			fmt.Print("Enter the task ID to mark as completed: ")

			fmt.Scanln(&input)

			id, err := strconv.Atoi(strings.TrimSpace(input))

			if err != nil {
				fmt.Println("Invalid ID. Please enter a valid number.")
				continue
			}

			completeTask(db, id)

		case "4":
			fmt.Print("Enter the task ID to delete: ")

			fmt.Scanln(&input)

			id, err := strconv.Atoi(strings.TrimSpace(input))

			if err != nil {
				fmt.Println("Invalid ID. Please enter a valid number.")
				continue
			}

			fmt.Printf("Are you sure you want to delete task %d? (yes/no): ", id)

			// deleteTask(db, id)

		case "q":
			fmt.Println("Goodbye!")
			return // Exit the program

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
