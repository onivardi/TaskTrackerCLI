package task_test

import (
	"testing"
	"time"

	"github.com/onivardi/TaskTrackerCLI/task"
)

func TestTask(t *testing.T) {
	t.Parallel()
	_ = task.Task{
		Id:          1,
		Description: "My first task",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}
}
