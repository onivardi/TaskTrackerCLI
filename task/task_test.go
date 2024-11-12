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

	err := l.Add(want)
	if err != nil {
		t.Fatal(err)
	}

	got := l.Tasks[0].Description

	if want != got {
		t.Errorf("Want %s, Got %s", want, got)
	}
}

func TestAddInvalidInput(t *testing.T) {
	t.Parallel()
	l := task.ListTask{}

	desc := ""

	err := l.Add(desc)

	if err == nil {
		t.Fatal("want error for invalid input description, got nil")
	}
}

func TestDelete(t *testing.T) {
	t.Parallel()
	l := task.ListTask{}

	tasks := []string{
		"build todo list cli",
		"learn how to test in go",
	}
	for _, t := range tasks {
		_ = l.Add(t)
	}
	err := l.Delete(2)
	if err != nil {
		t.Fatal(err)
	}

	wantCount := 1
	gotCount := len(l.Tasks)
	if wantCount != gotCount {
		t.Errorf("expected %d tasks after deletion, got %d", wantCount, gotCount)
	}
}

func Test(t *testing.T) {
	t.Parallel()
	l := task.ListTask{}

	tasks := []string{
		"build todo list cli",
		"learn how to test in go",
	}
	for _, t := range tasks {
		_ = l.Add(t)
	}
	err := l.Delete(0)
	if err == nil {
		t.Fatal("want error for invalid id, got nil")
	}
}
