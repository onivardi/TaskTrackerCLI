package task

import (
	"flag"
	"fmt"
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
	Tasks map[int]Task
}

func (lt *ListTask) Add(description string) error {
	if description == "" {
		return fmt.Errorf("description cannot be empty; please privide a valid description")
	}
	newID := len(lt.Tasks) + 1
	t := Task{
		Id:          newID,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	lt.Tasks[newID] = t

	return nil
}

func (lt *ListTask) Delete(id int) error {
	if _, exists := lt.Tasks[id]; !exists {
		return fmt.Errorf("task ID %d not found; please provide a valid task ID", id)
	}
	delete(lt.Tasks, id)

	return nil
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
