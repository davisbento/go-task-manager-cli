package repository

type Task struct {
	ID          int
	Description string
	Completed   bool
}

type Repository interface {
	AddTask(description string) error
	ListTasks(status string) []Task
	CompleteTask(id int) error
	DeleteTask(id int) error
}
