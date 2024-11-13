package task_test

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/onivardi/TaskTrackerCLI/task"
)

func TestTask(t *testing.T) {
	t.Parallel()
	_ = task.Task{
		Id:          1,
		Description: "My first task",
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}
}

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

func TestGetFileNotExist(t *testing.T) {
	t.Parallel()
	l := task.ListTask{
		Tasks: make(map[int]task.Task),
	}

	// Simulating a saving json file
	tempFile, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	err = l.Get(tempFile.Name())
	if err == nil {
		t.Fatal("want file does not exist, got nil")
	}
}

func TestSaveAndGet(t *testing.T) {
	t.Parallel()

	l := task.ListTask{
		Tasks: make(map[int]task.Task),
	}
	ts := task.Task{
		Id:          1,
		Description: "build todo list cli",
	}
	l.Tasks[ts.Id] = ts

	err := l.Save("testData.json")
	if err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}
	err = l.Get("testData.json")
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

	l.Update(1, "update description")
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
