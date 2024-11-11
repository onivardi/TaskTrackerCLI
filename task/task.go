package task

import (
	"flag"
	"time"
)

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func Main() {
	task := flag.String("add", "", "Add a task")
	delete := flag.Int("delete", -1, "Delete a task on given ID")

	// Flags for updating a task
	taskID := flag.Int("id", -1, "Task ID to update. Work with -description and -status")
	newDescription := flag.String("description", "", "New description for the task")
	status := flag.Int("status", -1, "Update the status on the task.: Use 0 for todo, 1 for in-progress, 2 for done")

	flag.Parse()
}
