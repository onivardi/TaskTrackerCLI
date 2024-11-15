package task_test

import (
	"os"
	"reflect"
	"testing"

	task "github.com/onivardi/TaskTrackerCLI"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	l := task.ListTask{
		Tasks: make(map[int]task.Task),
	}

	want := "How to Become Sofware Engineer"

	err := l.Add(want)
	if err != nil {
		t.Fatal(err)
	}

	got := l.Tasks[1].Description

	if want != got {
		t.Errorf("Want %s, Got %s", want, got)
	}
}

func TestAddInvalidInput(t *testing.T) {
	t.Parallel()
	l := task.ListTask{
		Tasks: make(map[int]task.Task),
	}

	desc := ""

	err := l.Add(desc)

	if err == nil {
		t.Fatal("want error for invalid input description, got nil")
	}
}

func TestDelete(t *testing.T) {
	t.Parallel()
	l := task.ListTask{
		Tasks: make(map[int]task.Task),
	}

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

func TestDeleteInvalidInputID(t *testing.T) {
	t.Parallel()
	l := task.ListTask{
		Tasks: make(map[int]task.Task),
	}

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

func TestSaveAndGetAll(t *testing.T) {
	t.Parallel()

	l := task.ListTask{
		Tasks: make(map[int]task.Task),
	}
	ts := task.Task{
		Id:          1,
		Description: "build todo list cli",
	}
	l.Tasks[ts.Id] = ts

	tempFile, err := os.CreateTemp("", "testData.json")
	if err != nil {
		t.Fatal(err)
	}

	err = l.Save(tempFile.Name())
	if err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}
	err = l.GetAll(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	expected := map[int]task.Task{
		1: {Id: 1, Description: "build todo list cli"},
	}
	got := l.Tasks
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()
	l := task.ListTask{
		Tasks: make(map[int]task.Task),
	}

	tasks := []string{
		"build todo list cli",
		"learn how to test in go",
	}
	for _, t := range tasks {
		_ = l.Add(t)
	}

	err := l.Update(1, "update description")
	if err != nil {
		t.Fatal(err)
	}

	got := l.Tasks[1].Description
	want := "update description"
	if got != want {
		t.Errorf("Want %s, Got %s", want, got)
	}
}

func TestUpdateInvalidInput(t *testing.T) {
	t.Parallel()

	l := task.ListTask{
		Tasks: make(map[int]task.Task),
	}

	tasks := []string{
		"build todo list cli",
		"learn how to t in go",
	}
	for _, t := range tasks {
		_ = l.Add(t)
	}

	// wrong id
	err := l.Update(3, "this id does not exist")
	if err == nil {
		t.Fatal("want error for invalid id, got nil")
	}
	// empty description
	err = l.Update(2, "")
	if err == nil {
		t.Fatal("want error for invalid description, got nil")
	}
}

func TestUpdateStatus(t *testing.T) {
	t.Parallel()

	l := task.ListTask{Tasks: make(map[int]task.Task)}
	l.Add("test")

	err := l.UpdateStatus(1, 2)
	if err != nil {
		t.Fatal(err)
	}

	want := task.InProgress
	got := l.Tasks[1].Status
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestUpdateStatusInvalidInput(t *testing.T) {
	t.Parallel()
	l := task.ListTask{Tasks: make(map[int]task.Task)}
	l.Add("test")
	err := l.UpdateStatus(0, 1)
	if err == nil {
		t.Fatal("want error for invalid id, got nil")
	}

	err = l.UpdateStatus(1, 999)
	if err == nil {
		t.Fatal("want error for invalid status, got nil")
	}
}

// FIXME: its returning createdAt and updatedAt
// INFO: Fixed
func TestGetTaskByStatus(t *testing.T) {
	t.Parallel()

	l := task.ListTask{Tasks: map[int]task.Task{
		1: {Id: 1, Description: "test", Status: 1},
		2: {Id: 2, Description: "test2", Status: 1},
		3: {Id: 3, Description: "test3", Status: 0},
		4: {Id: 4, Description: "test4", Status: 2},
	}}

	want := map[int]task.Task{
		1: {Id: 1, Description: "test", Status: 1},
		2: {Id: 2, Description: "test2", Status: 1},
	}
	newLT, err := l.GetTasksByStatus(1)
	if err != nil {
		t.Fatal(err)
	}

	got := newLT.Tasks

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

// FIXME: its returning createdAt and updatedAt
// INFO: Fixed
func TestGetTaskByStatusInvalidInput(t *testing.T) {
	t.Parallel()

	l := task.ListTask{Tasks: map[int]task.Task{
		1: {Id: 1, Description: "test", Status: 1},
	}}
	// l.Add("test")
	// l.UpdateStatus(1, 1)

	_, err := l.GetTasksByStatus(999)
	if err == nil {
		t.Fatal("want error for invalid status, got nil")
	}
}
