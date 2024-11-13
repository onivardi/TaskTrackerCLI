package task_test

import (
	"encoding/json"
	"log"
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

func TestGet(t *testing.T) {
	t.Parallel()

	l := task.ListTask{
		Tasks: make(map[int]task.Task),
	}

	ts := task.Task{
		Id:          1,
		Description: "build todo list cli",
	}

	l.Tasks[ts.Id] = ts

	// Simulating a saving json file
	tempFile, err := os.CreateTemp("", "testData.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	jsonData, err := json.MarshalIndent(l.Tasks, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	_, err = tempFile.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
	// ------------------------------

	err = l.Get(tempFile.Name())
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

func TestGetFileNotExist(t *testing.T) {
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
