package main

import (
	"bufio"
	"davisbento/go-task-manager-cli/infra"
	"davisbento/go-task-manager-cli/repository"
	"davisbento/go-task-manager-cli/sqlite_repository"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := infra.InitDB()

	defer db.Close()

	var repository repository.Repository = sqlite_repository.NewRepository(db)

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

			err := repository.AddTask(description)

			if err != nil {
				fmt.Println("Failed to add task:", err)
				continue
			}

			fmt.Println("Task added successfully!")

		case "2":
			fmt.Println("List of Tasks:")
			tasks := repository.ListTasks("")

			if len(tasks) == 0 {
				fmt.Println("======================")
				fmt.Println("No tasks found.")
				fmt.Println("======================")
				continue
			}

			// create a tabular format
			fmt.Println("ID\tDescription\tCompleted")
			fmt.Println("----\t-----------\t--------")

			for _, task := range tasks {
				completed := "No"

				if task.Completed {
					completed = "Yes"
				}

				// create a tabular format position the characters in the table
				fmt.Printf("%d\t%s\t\t%s\n", task.ID, task.Description, completed)
			}

			fmt.Println()

		case "3":
			fmt.Print("Enter the task ID to mark as completed: ")

			fmt.Scanln(&input)

			id, err := strconv.Atoi(strings.TrimSpace(input))

			if err != nil {
				fmt.Println("Invalid ID. Please enter a valid number.")
				continue
			}

			if err := repository.CompleteTask(id); err != nil {
				fmt.Println("Failed to mark task as completed:", err)
				continue
			}

			fmt.Println("Task marked as completed successfully!")

		case "4":
			fmt.Print("Enter the task ID to delete: ")

			fmt.Scanln(&input)

			id, err := strconv.Atoi(strings.TrimSpace(input))

			if err != nil {
				fmt.Println("Invalid ID. Please enter a valid number.")
				continue
			}

			if err := repository.DeleteTask(id); err != nil {
				fmt.Println("Failed to delete task:", err)
				continue
			}

			fmt.Println("Task deleted successfully!")

		case "q":
			fmt.Println("Goodbye!")
			return // Exit the program

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
