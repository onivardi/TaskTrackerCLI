package task_test

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"

	task "github.com/onivardi/TaskTrackerCLI"
)

var validInput = "test add 1"

func createListTasksData() task.ListTask {
	return task.ListTask{
		Tasks: map[int]task.Task{
			1: {Id: 1, Description: "test", Status: 2},
			2: {Id: 2, Description: "test2", Status: 1},
			3: {Id: 3, Description: "test3", Status: 0},
			4: {Id: 4, Description: "test4", Status: 2},
		},
	}
}

func TestAdd_ATaskWithValidInput(t *testing.T) {
	t.Parallel()
	l := task.ListTask{
		Tasks: make(map[int]task.Task),
	}

	err := l.Add(validInput)
	if err != nil {
		t.Fatal(err)
	}

	wantCount := 1
	gotCount := len(l.Tasks)

	if wantCount != gotCount {
		t.Errorf("Want %d, Got %d", wantCount, gotCount)
	}
}

func TestAdd_ReturnsErrorForInvalidInput(t *testing.T) {
	lt := task.ListTask{Tasks: make(map[int]task.Task)}

	testCases := map[string]struct {
		invalidInput string
	}{
		"empty string": {
			invalidInput: "",
		},

		"more than 60 words": {
			invalidInput: strings.Repeat("test", 61),
		},
	}
	for name, tC := range testCases {
		t.Run(name, func(t *testing.T) {
			err := lt.Add(tC.invalidInput)
			if err == nil {
				t.Fatal("want error for invalid input description, got nil")
			}
		})
	}
}

func TestDelete_ATaskWithValidId(t *testing.T) {
	t.Parallel()

	lt := createListTasksData()
	beforeLen := len(lt.Tasks)
	t.Logf("before, data size: %d", beforeLen)

	err := lt.Delete(2)
	if err != nil {
		t.Fatal(err)
	}
	afterLen := len(lt.Tasks)
	t.Logf("After, data size: %d", afterLen)

	wantCount := beforeLen - 1
	gotCount := afterLen
	if wantCount != gotCount {
		t.Errorf("expected %d tasks after deletion, got %d", wantCount, gotCount)
	}
}

func TestDelete_ReturnsErrorWithInvalidId(t *testing.T) {
	t.Parallel()

	lt := createListTasksData()
	beforeLen := len(lt.Tasks)
	t.Logf("before, data size: %d", beforeLen)

	err := lt.Delete(0)
	afterLen := len(lt.Tasks)
	t.Logf("After, data size: %d", afterLen)

	if err == nil {
		t.Fatal("want error for invalid id, got nil")
	}
}

func TestSave_WriteToJsonFile(t *testing.T) {
	t.Parallel()

	lt := createListTasksData()

	tempFile, err := os.CreateTemp("", "testData.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	err = lt.Save(tempFile.Name())
	if err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	var savedLT task.ListTask
	data, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(data, &savedLT)
	if err != nil {
		t.Fatal(err)
	}

	want := lt.Tasks
	got := savedLT.Tasks

	if !reflect.DeepEqual(want, got) {
		t.Fatal("want saved list tasks to be equal to the original list tasks")
	}
}

func TestGetAll_LoadJsonFile(t *testing.T) {
	lt := &task.ListTask{
		Tasks: make(map[int]task.Task),
	}

	tmpFile, err := os.CreateTemp("", "testData.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	testData := []byte(`{"Tasks": {"1": {"Id":1, "Description":"Task 1", "Status": 0}, "2": {"Id":2, "Description":"Task 2", "Status":2}}}`)
	_, err = tmpFile.Write(testData)
	if err != nil {
		t.Fatal(err)
	}

	err = lt.GetAll(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	want := map[int]task.Task{
		1: {Id: 1, Description: "Task 1", Status: task.Todo},
		2: {Id: 2, Description: "Task 2", Status: task.InProgress},
	}

	if !reflect.DeepEqual(want, lt.Tasks) {
		t.Errorf("want loaded list tasks to be equal to the test data, got %+v vs %+v", lt.Tasks, want)
	}
}

func TestUpdate_DescriptionWithValidInput(t *testing.T) {
	t.Parallel()
	lt := createListTasksData()

	err := lt.Update(1, "this is a new description")
	if err != nil {
		t.Fatal(err)
	}

	got := lt.Tasks[1].Description
	want := "this is a new description"
	if got != want {
		t.Errorf("Want %s, Got %s", want, got)
	}
}

func TestUpdate_ReturnsErrorWithInvalidInput(t *testing.T) {
	lt := createListTasksData()

	testCases := map[string]struct {
		id          int
		description string
	}{
		"invalid id": {
			id:          50,
			description: "wrong id",
		},
		"empty description": {
			id:          2,
			description: "",
		},
		"both invalid": {
			id:          0,
			description: "",
		},
		"more than 60 words": {
			id:          1,
			description: strings.Repeat("test", 61),
		},
	}

	for name, tC := range testCases {
		t.Run(name, func(t *testing.T) {
			err := lt.Update(tC.id, tC.description)
			if err == nil {
				t.Fatal("want error for invalid input, got nil")
			}
		})
	}
}

func TestUpdateStatus_WithValidInputId(t *testing.T) {
	t.Parallel()

	lt := createListTasksData()

	err := lt.UpdateStatus(1, task.InProgress)
	if err != nil {
		t.Fatal(err)
	}

	want := task.InProgress
	got := lt.Tasks[1].Status
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestUpdateStatus_ReturnsErrorWithInvalidInput(t *testing.T) {
	lt := createListTasksData()

	testCases := map[string]struct {
		id     int
		status task.Status
	}{
		"invalid id": {
			id:     0,
			status: task.Done,
		},
		"invalid status": {
			id:     1,
			status: 999,
		},
		"invalid both": {
			id:     0,
			status: 999,
		},
		"status todo": {
			id:     1,
			status: task.Todo,
		},
	}
	for name, tC := range testCases {
		t.Run(name, func(t *testing.T) {
			err := lt.UpdateStatus(tC.id, tC.status)
			if err == nil {
				t.Fatal("want error for invalid input, got nil")
			}
		})
	}
}

func TestGetTasksByStatus_WithValidInput(t *testing.T) {
	t.Parallel()

	lt := createListTasksData()

	newLT, err := lt.GetTasksByStatus(2)
	if err != nil {
		t.Fatal(err)
	}

	want := map[int]task.Task{
		1: {Id: 1, Description: "test", Status: 2},
		4: {Id: 4, Description: "test4", Status: 2},
	}
	got := newLT.Tasks

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestGetTaskByStatus_ReturnsErrorWithInvalidInput(t *testing.T) {
	t.Parallel()

	lt := createListTasksData()

	_, err := lt.GetTasksByStatus(999)
	if err == nil {
		t.Fatal("want error for invalid status, got nil")
	}
}
