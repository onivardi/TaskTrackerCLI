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

func TestAdd(t *testing.T) {
	t.Parallel()
	l := task.ListTask{}

	want := "How to Become Sofware Engineer"

	l.Add(want)

	got := l.Tasks[0].Description

	if want != got {
		t.Errorf("Want %s, Got %s", want, got)
	}
}
