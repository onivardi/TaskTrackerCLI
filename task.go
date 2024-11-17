package task

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
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

	if len(description) > 60 {
		return fmt.Errorf("description cannot be more than 60 words; please privide a valid description")
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

func (lt *ListTask) Update(id int, description string) error {
	if _, exists := lt.Tasks[id]; !exists {
		return fmt.Errorf("task ID %d not found; please provide a valid task ID", id)
	}

	if description == "" {
		return fmt.Errorf("description cannot be empty; please privide a valid description")
	}

	t := lt.Tasks[id]
	t.Description = description
	t.UpdatedAt = time.Now()
	lt.Tasks[id] = t

	return nil
}

func (lt *ListTask) UpdateStatus(id int, status Status) error {
	if _, exists := lt.Tasks[id]; !exists {
		return fmt.Errorf("task ID %d not found; please provide a valid task ID", id)
	}

	if !validStatus[status] {
		return fmt.Errorf("invalid status; please provide a valid status")
	}

	t := lt.Tasks[id]
	t.Status = status
	t.UpdatedAt = time.Now()
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
	delete := flag.Int("delete", -1, "Delete a task on the given ID")
	markInProgress := flag.Int("markInProgress", -1, "Mark a task on the given ID as in progress")
	markDone := flag.Int("markDone", -1, "Mark a task on the given ID as done")
	update := flag.Bool("update", false, "Update a task on the given ID")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of task-cli:
  -add <task_description>       Add a new task with the given description.
  -list [done|in-progress|todo] List tasks filtered by status or all tasks if no filter is specified.
  -delete <task_id>             Delete the task with the specified ID.
  -markInProgress <task_id>     Mark the task with the specified ID as "In Progress".
  -markDone <task_id>           Mark the task with the specified ID as "Done".
  -update <task_id> <description> 
                                 Update the description of the task with the specified ID.

Examples:
  Add a task:
    task-cli -add "Buy groceries"

  List all tasks:
    task-cli -list

  List tasks by status:
    task-cli -list done
    task-cli -list in-progress
    task-cli -list todo

  Delete a task:
    task-cli -delete 3

  Update a task description:
    task-cli -update 2 "Pick up laundry"

  Mark a task as in progress:
    task-cli -markInProgress 5

  Mark a task as done:
    task-cli -markDone 6
`)
	}
	flag.Parse()
	lt := &ListTask{Tasks: make(map[int]Task)}

	// Load existing tasks from file
	if err := lt.GetAll(fileName); err != nil {
		fmt.Fprintln(os.Stderr, "Error loading tasks:", err)
		return 1
	}

	switch {
	case *list:
		args := flag.Args()
		if len(args) > 0 {
			switch args[0] {
			case "done":
				tasks, _ := lt.GetTasksByStatus(Done)
				fmt.Println("Tasks marked as Done:")
				fmt.Print(tasks)
			case "in-progress":
				tasks, _ := lt.GetTasksByStatus(InProgress)
				fmt.Println("Tasks in Progress:")
				fmt.Print(tasks)
			case "todo":
				validStatus[Todo] = true
				tasks, _ := lt.GetTasksByStatus(Todo)
				fmt.Println("Tasks To-Do:")
				fmt.Print(tasks)
				validStatus[Todo] = false
			default:
				fmt.Println("Invalid list option. Use 'done', 'in-progress', or 'todo'.")
			}
		} else {
			fmt.Println("All tasks:")
			fmt.Print(lt)
		}

	case *add != "":
		if err := lt.Add(*add); err != nil {
			fmt.Fprintln(os.Stderr, "Error adding task:", err)
			return 1
		}
		if err := lt.Save(fileName); err != nil {
			fmt.Fprintln(os.Stderr, "Error saving tasks:", err)
			return 1
		}
		fmt.Println("Task added successfully!")

	case *delete > 0:
		if err := lt.Delete(*delete); err != nil {
			fmt.Fprintln(os.Stderr, "Error deleting task:", err)
			return 1
		}
		if err := lt.Save(fileName); err != nil {
			fmt.Fprintln(os.Stderr, "Error saving tasks:", err)
			return 1
		}
		fmt.Println("Task deleted successfully!")

	case *update:
		args := flag.Args()
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "Invalid arguments. Usage: -update <ID> <new description>")
			return 1
		}
		taskID, err := strconv.Atoi(args[0])
		if err != nil || taskID <= 0 {
			fmt.Fprintln(os.Stderr, "Invalid task ID")
			return 1
		}
		description := args[1]
		if err := lt.Update(taskID, description); err != nil {
			fmt.Fprintln(os.Stderr, "Error updating task:", err)
			return 1
		}
		if err := lt.Save(fileName); err != nil {
			fmt.Fprintln(os.Stderr, "Error saving tasks:", err)
			return 1
		}
		fmt.Println("Task updated successfully!")

	case *markInProgress > 0:
		if err := lt.UpdateStatus(*markInProgress, InProgress); err != nil {
			fmt.Fprintln(os.Stderr, "Error marking task as in progress:", err)
			return 1
		}
		if err := lt.Save(fileName); err != nil {
			fmt.Fprintln(os.Stderr, "Error saving tasks:", err)
			return 1
		}
		fmt.Println("Task marked as In Progress!")

	case *markDone > 0:
		if err := lt.UpdateStatus(*markDone, Done); err != nil {
			fmt.Fprintln(os.Stderr, "Error marking task as done:", err)
			return 1
		}
		if err := lt.Save(fileName); err != nil {
			fmt.Fprintln(os.Stderr, "Error saving tasks:", err)
			return 1
		}
		fmt.Println("Task marked as Done!")

	default:
		fmt.Fprintln(os.Stderr, "Invalid option. Use -h for help.")
		return 1
	}

	return 0
}
