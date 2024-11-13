package task

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
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
	// defer fmt.Println("Task added successfully (ID:%i)", newID)
	return nil
}

func (lt *ListTask) Delete(id int) error {
	if _, exists := lt.Tasks[id]; !exists {
		return fmt.Errorf("task ID %d not found; please provide a valid task ID", id)
	}
	delete(lt.Tasks, id)

	return nil
}

// Read a json file and load to the ListTask map
func (lt *ListTask) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("file does not exist, create one on current directory")
		}
		return err
	}
	if len(file) == 0 {
		return fmt.Errorf("There is no task added")
	}
	return json.Unmarshal(file, lt)
}

func (lt ListTask) Save(filename string) error {
	js, err := json.Marshal(lt)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, js, 0644)
}

func (lt *ListTask) Update(id int, description string) error {
	if _, exists := lt.Tasks[id]; !exists {
		return fmt.Errorf("task ID %d not found; please provide a valid task ID", id)
	}

	if description == "" {
		return fmt.Errorf("description cannot be empty; please privide a valid description")
	}

	d := lt.Tasks[id]
	d.Description = description
	lt.Tasks[id] = d

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
