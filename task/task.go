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

type ListTask struct {
	Tasks []Task
}

func (lt *ListTask) Add(description string) {
	t := Task{
		Id:          len(lt.Tasks) + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	lt.Tasks = append(lt.Tasks, t)
}

func Main() {
	// task := flag.String("add", "", "Add a task")
	// delete := flag.Int("delete", -1, "Delete a task on given ID")

	// Flags for updating a task
	// taskID := flag.Int("id", -1, "Task ID to update. Work with -description and -status")
	// newDescription := flag.String("description", "", "New description for the task")
	// status := flag.Int("status", -1, "Update the status on the task.: Use 0 for todo, 1 for in-progress, 2 for done")

	flag.Parse()
}
