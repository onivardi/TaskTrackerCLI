package task

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"
)

type Status int

const (
	Todo Status = iota
	Done
	InProgress
)

const fileName = "tasks.json"

// only update between done and inProgress
var validStatus = map[Status]bool{
	Todo:       false,
	Done:       true,
	InProgress: true,
}

type Task struct {
	Id          int
	Description string
	Status      Status
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
		Status:      Todo,
		CreatedAt:   time.Now(),
	}

	lt.Tasks[newID] = t
	// defer fmt.Printf("Task added successfully (ID: %d)", newID)
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
// FIXME: Not letting the user add tasks
// INFO: Fixed
func (lt *ListTask) GetAll(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, lt)
}

// INFO: Updated to return a new ListTask with the status selected, to work with Stringer Interface -> fmt.Print()
func (lt ListTask) GetTasksByStatus(status Status) (ListTask, error) {
	if !validStatus[status] {
		return ListTask{}, fmt.Errorf("invalid status; please provide a valid status")
	}
	tasks := make(map[int]Task)
	for id, task := range lt.Tasks {
		if task.Status == status {
			tasks[id] = task
		}
	}

	lt.Tasks = tasks
	return lt, nil
}

func (lt ListTask) Save(filename string) error {
	js, err := json.Marshal(lt)
	if err != nil {
		return fmt.Errorf("cannot save your tasks to file: %w", err)
	}
	return os.WriteFile(filename, js, 0644)
}

// TODO: Should update the updatedAt attribute too
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

// TODO: Should update the updatedAt attribute too
func (lt *ListTask) UpdateStatus(id int, status Status) error {
	if _, exists := lt.Tasks[id]; !exists {
		return fmt.Errorf("task ID %d not found; please provide a valid task ID", id)
	}

	if !validStatus[status] {
		return fmt.Errorf("invalid status; please provide a valid status")
	}

	t := lt.Tasks[id]
	t.Status = status
	lt.Tasks[id] = t
	return nil
}

func (t Task) GetStatus() Status {
	return t.Status
}

// String prints out a formatted list
// Implements the fmt.Stringer interface
// INFO: Every time the fmt.Print() is called, will be formatted
func (lt ListTask) String() string {
	formatted := ""
	for _, t := range lt.Tasks {

		prefix := ""

		switch t.Status {
		case Done:

			prefix = "✅ " // Symbol for completed tasks
		case InProgress:

			prefix = "🚧 " // Symbol for in progress tasks
		default:

			prefix = "⬜" // Symbol for incomplete tasks
		}

		// Adjust the item number k to print numbers starting from 1 instead of 0
		formatted += fmt.Sprintf("%s- %s (ID: %d)\n", prefix, t.Description, t.Id)
	}

	return formatted
}

func Main() int {
	add := flag.String("add", "", "Add a task")
	list := flag.Bool("list", false, "List all tasks")
	delete := flag.Int("delete", -1, "Delete a task on given ID")
	markInProgress := flag.Int("markInProgress", -1, "Mark a task on given ID as in progress")
	markDone := flag.Int("markDone", -1, "Mark a task on given ID as done")

	// Flags for updating a task
	// taskID := flag.Int("id", -1, "Task ID to update. Work with -description and -status")
	// newDescription := flag.String("description", "", "New description for the task")
	// status := flag.Int("status", -1, "Update the status on the task.: Use 0 for todo, 1 for in-progress, 2 for done")

	flag.Parse()

	lt := &ListTask{Tasks: make(map[int]Task)}
	if err := lt.GetAll(fileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	switch {
	case *list:
		// args := flag.Args()

		// if len(args) > 0 {
		// 	switch args[0] {
		// 	case "done":
		// 		tasks, _ := lt.GetTasksByStatus(Done)
		// 		fmt.Print(tasks)
		//
		// 	case "in-progress":
		// 		tasks, _ := lt.GetTasksByStatus(InProgress)
		// 		fmt.Print(tasks)
		//
		// 	default:
		// 		fmt.Println("Invalid list option. Use 'done' or 'inprogress'")
		// 	}
		// }

		// tasks, _ := lt.GetTasksByStatus(Done)
		// lt.Tasks = tasks
		fmt.Print(lt)

	case *add != "":
		lt.Add(*add)

		if err := lt.Save(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}

	case *delete > 0:
		lt.Delete(*delete)

		if err := lt.Save(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}

	case *markInProgress > 0:
		lt.UpdateStatus(*markInProgress, InProgress)

		if err := lt.Save(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}

	case *markDone > 0:
		lt.UpdateStatus(*markDone, Done)

		if err := lt.Save(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}

	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		return 1
	}

	return 0
}
