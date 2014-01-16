/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package todotxt

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

var (
	testInputTasklist                   = "testdata/tasklist_todo.txt"
	testInputTasklistCreatedDateError   = "testdata/tasklist_createdDate_error.txt"
	testInputTasklistDueDateError       = "testdata/tasklist_dueDate_error.txt"
	testInputTasklistCompletedDateError = "testdata/tasklist_completedDate_error.txt"
	testInputTasklistScannerError       = "testdata/tasklist_scanner_error.txt"
	testOutput                          = "testdata/ouput_todo.txt"
	testExpectedOutput                  = "testdata/expected_todo.txt"
	testTasklist                        TaskList
	testExpected                        interface{}
	testGot                             interface{}
)

func TestLoadFromFile(t *testing.T) {
	file, err := os.Open(testInputTasklist)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if testTasklist, err := LoadFromFile(file); err != nil {
		t.Fatal(err)
	} else {
		data, err := ioutil.ReadFile(testExpectedOutput)
		if err != nil {
			t.Fatal(err)
		}
		testExpected = string(data)
		testGot = testTasklist.String()
		if testGot != testExpected {
			t.Errorf("Expected TaskList to be [%s], but got [%s]", testExpected, testGot)
		}
	}

	if testTasklist, err := LoadFromFile(nil); testTasklist != nil || err == nil {
		t.Errorf("Expected LoadFromFile to fail, but got TaskList back: [%s]", testTasklist)
	}
}

func TestLoadFromFilename(t *testing.T) {
	if testTasklist, err := LoadFromFilename(testInputTasklist); err != nil {
		t.Fatal(err)
	} else {
		data, err := ioutil.ReadFile(testExpectedOutput)
		if err != nil {
			t.Fatal(err)
		}
		testExpected = string(data)
		testGot = testTasklist.String()
		if testGot != testExpected {
			t.Errorf("Expected TaskList to be [%s], but got [%s]", testExpected, testGot)
		}
	}

	if testTasklist, err := LoadFromFilename("some_file_that_does_not_exists.txt"); testTasklist != nil || err == nil {
		t.Errorf("Expected LoadFromFilename to fail, but got TaskList back: [%s]", testTasklist)
	}
}

func TestWriteFile(t *testing.T) {
	os.Remove(testOutput)
	os.Create(testOutput)
	var err error

	fileInput, err := os.Open(testInputTasklist)
	if err != nil {
		t.Fatal(err)
	}
	defer fileInput.Close()
	fileOutput, err := os.OpenFile(testOutput, os.O_RDWR, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer fileInput.Close()

	if testTasklist, err = LoadFromFile(fileInput); err != nil {
		t.Fatal(err)
	}
	if err = WriteToFile(&testTasklist, fileOutput); err != nil {
		t.Fatal(err)
	}
	fileInput.Close()
	fileOutput, err = os.Open(testOutput)
	if err != nil {
		t.Fatal(err)
	}
	if testTasklist, err = LoadFromFile(fileOutput); err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadFile(testExpectedOutput)
	if err != nil {
		t.Fatal(err)
	}
	testExpected = string(data)
	testGot = testTasklist.String()
	if testGot != testExpected {
		t.Errorf("Expected TaskList to be [%s], but got [%s]", testExpected, testGot)
	}
}

func TestTaskListWriteFile(t *testing.T) {
	os.Remove(testOutput)
	os.Create(testOutput)
	testTasklist := TaskList{}

	fileInput, err := os.Open(testInputTasklist)
	if err != nil {
		t.Fatal(err)
	}
	defer fileInput.Close()
	fileOutput, err := os.OpenFile(testOutput, os.O_RDWR, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer fileInput.Close()

	if err := testTasklist.LoadFromFile(fileInput); err != nil {
		t.Fatal(err)
	}
	if err := testTasklist.WriteToFile(fileOutput); err != nil {
		t.Fatal(err)
	}
	fileInput.Close()
	fileOutput, err = os.Open(testOutput)
	if err != nil {
		t.Fatal(err)
	}
	if err := testTasklist.LoadFromFile(fileOutput); err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadFile(testExpectedOutput)
	if err != nil {
		t.Fatal(err)
	}
	testExpected = string(data)
	testGot = testTasklist.String()
	if testGot != testExpected {
		t.Errorf("Expected TaskList to be [%s], but got [%s]", testExpected, testGot)
	}
}

func TestWriteFilename(t *testing.T) {
	os.Remove(testOutput)
	var err error

	if testTasklist, err = LoadFromFilename(testInputTasklist); err != nil {
		t.Fatal(err)
	}
	if err = WriteToFilename(&testTasklist, testOutput); err != nil {
		t.Fatal(err)
	}
	if testTasklist, err = LoadFromFilename(testOutput); err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadFile(testExpectedOutput)
	if err != nil {
		t.Fatal(err)
	}
	testExpected = string(data)
	testGot = testTasklist.String()
	if testGot != testExpected {
		t.Errorf("Expected TaskList to be [%s], but got [%s]", testExpected, testGot)
	}
}

func TestTaskListWriteFilename(t *testing.T) {
	os.Remove(testOutput)
	testTasklist := TaskList{}

	if err := testTasklist.LoadFromFilename(testInputTasklist); err != nil {
		t.Fatal(err)
	}
	if err := testTasklist.WriteToFilename(testOutput); err != nil {
		t.Fatal(err)
	}
	if err := testTasklist.LoadFromFilename(testOutput); err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadFile(testExpectedOutput)
	if err != nil {
		t.Fatal(err)
	}
	testExpected = string(data)
	testGot = testTasklist.String()
	if testGot != testExpected {
		t.Errorf("Expected TaskList to be [%s], but got [%s]", testExpected, testGot)
	}
}

func TestNewTaskList(t *testing.T) {
	testTasklist := NewTaskList()

	testExpected = 0
	testGot = len(testTasklist)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}
}

func TestTaskListCount(t *testing.T) {
	if err := testTasklist.LoadFromFilename(testInputTasklist); err != nil {
		t.Fatal(err)
	}

	testExpected = 63
	testGot = len(testTasklist)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}
}

func TestTaskListAddTask(t *testing.T) {
	if err := testTasklist.LoadFromFilename(testInputTasklist); err != nil {
		t.Fatal(err)
	}

	// add new empty task
	task := NewTask()
	testTasklist.AddTask(&task)

	testExpected = 64
	testGot = len(testTasklist)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}

	taskId := 64
	testExpected = time.Now().Format(DateLayout) + " " // tasks created by NewTask() have their created date set
	testGot = testTasklist[taskId-1].String()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
	testExpected = 64
	testGot = testTasklist[taskId-1].Id
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%d], but got [%d]", taskId, testExpected, testGot)
	}
	taskId++

	// add parsed task
	parsed, err := ParseTask("x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12")
	if err != nil {
		t.Error(err)
	}
	testTasklist.AddTask(parsed)

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[taskId-1].String()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
	testExpected = 65
	testGot = testTasklist[taskId-1].Id
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%d], but got [%d]", taskId, testExpected, testGot)
	}
	taskId++

	// add selfmade task
	createdDate := time.Now()
	testTasklist.AddTask(&Task{
		CreatedDate: createdDate,
		Todo:        "Go shopping..",
		Contexts:    []string{"GroceryStore"},
	})

	testExpected = createdDate.Format(DateLayout) + " Go shopping.. @GroceryStore"
	testGot = testTasklist[taskId-1].String()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
	testExpected = 66
	testGot = testTasklist[taskId-1].Id
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%d], but got [%d]", taskId, testExpected, testGot)
	}
	taskId++

	// add task with explicit Id, AddTask() should ignore this!
	testTasklist.AddTask(&Task{
		Id: 101,
	})

	testExpected = 67
	testGot = testTasklist[taskId-1].Id
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%d], but got [%d]", taskId, testExpected, testGot)
	}
	taskId++
}

func TestTaskListGetTask(t *testing.T) {
	if err := testTasklist.LoadFromFilename(testInputTasklist); err != nil {
		t.Fatal(err)
	}

	taskId := 3
	task, err := testTasklist.GetTask(taskId)
	if err != nil {
		t.Error(err)
	}
	testExpected = "(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
	testGot = task.String()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
	testExpected = 3
	testGot = testTasklist[taskId-1].Id
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%d], but got [%d]", taskId, testExpected, testGot)
	}
	taskId++
}

func TestTaskListUpdateTask(t *testing.T) {
	if err := testTasklist.LoadFromFilename(testInputTasklist); err != nil {
		t.Fatal(err)
	}

	taskId := 3
	task, err := testTasklist.GetTask(taskId)
	if err != nil {
		t.Error(err)
	}
	testExpected = "(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
	testGot = task.String()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
	testExpected = 3
	testGot = testTasklist[taskId-1].Id
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%d], but got [%d]", taskId, testExpected, testGot)
	}

	task.Priority = "C"
	task.Todo = "Go home!"
	date, err := time.Parse(DateLayout, "2011-11-11")
	if err != nil {
		t.Error(err)
	}
	task.DueDate = date
	testGot := task

	os.Remove(testOutput)
	if err := testTasklist.WriteToFilename(testOutput); err != nil {
		t.Fatal(err)
	}
	if err := testTasklist.LoadFromFilename(testOutput); err != nil {
		t.Fatal(err)
	}
	testExpected, err := testTasklist.GetTask(taskId)
	if err != nil {
		t.Error(err)
	}
	if testGot.Task() != testExpected.Task() {
		t.Errorf("Expected Task to be [%v]\n, but got [%v]", testExpected, testGot)
	}
}

func TestTaskListRemoveTaskById(t *testing.T) {
	if err := testTasklist.LoadFromFilename(testInputTasklist); err != nil {
		t.Fatal(err)
	}

	taskId := 10
	if err := testTasklist.RemoveTaskById(taskId); err != nil {
		t.Error(err)
	}
	testExpected = 62
	testGot = len(testTasklist)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}
	task, err := testTasklist.GetTask(taskId)
	if err == nil || task != nil {
		t.Errorf("Expected no Task to be found anymore, but got %v", task)
	}

	taskId = 27
	if err := testTasklist.RemoveTaskById(taskId); err != nil {
		t.Error(err)
	}
	testExpected = 61
	testGot = len(testTasklist)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}
	task, err = testTasklist.GetTask(taskId)
	if err == nil || task != nil {
		t.Errorf("Expected no Task to be found anymore, but got %v", task)
	}

	taskId = 99
	if err := testTasklist.RemoveTaskById(taskId); err == nil {
		t.Errorf("Expected no Task to be found for removal")
	}
}

func TestTaskListRemoveTask(t *testing.T) {
	if err := testTasklist.LoadFromFilename(testInputTasklist); err != nil {
		t.Fatal(err)
	}

	taskId := 52 // Is "unique" in tasklist
	task, err := testTasklist.GetTask(taskId)
	if err != nil {
		t.Error(err)
	}

	if err := testTasklist.RemoveTask(*task); err != nil {
		t.Error(err)
	}
	testExpected = 62
	testGot = len(testTasklist)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}
	task, err = testTasklist.GetTask(taskId)
	if err == nil || task != nil {
		t.Errorf("Expected no Task to be found anymore, but got %v", task)
	}

	taskId = 2 // Exists 3 times in tasklist
	task, err = testTasklist.GetTask(taskId)
	if err != nil {
		t.Error(err)
	}

	if err := testTasklist.RemoveTask(*task); err != nil {
		t.Error(err)
	}
	testExpected = 59
	testGot = len(testTasklist)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}
	task, err = testTasklist.GetTask(taskId)
	if err == nil || task != nil {
		t.Errorf("Expected no Task to be found anymore, but got %v", task)
	}

	if err := testTasklist.RemoveTask(NewTask()); err == nil {
		t.Errorf("Expected no Task to be found for removal")
	}
}

func TestTaskListFilter(t *testing.T) {
	if err := testTasklist.LoadFromFilename(testInputTasklist); err != nil {
		t.Fatal(err)
	}

	// Filter list to get only completed tasks
	completedList := testTasklist.Filter(func(t Task) bool { return t.Completed })
	testExpected = 33
	testGot = len(*completedList)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}

	// Filter list to get only tasks with a due date
	dueDateList := testTasklist.Filter(func(t Task) bool { return t.HasDueDate() })
	testExpected = 26
	testGot = len(*dueDateList)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}

	// Filter list to get only tasks with "B" priority
	prioBList := testTasklist.Filter(func(t Task) bool {
		return t.HasPriority() && t.Priority == "B"
	})
	testExpected = 17
	testGot = len(*prioBList)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}
}

func TestTaskListReadErrors(t *testing.T) {
	if testTasklist, err := LoadFromFilename(testInputTasklistCreatedDateError); testTasklist != nil || err == nil {
		t.Errorf("Expected LoadFromFilename to fail because of invalid created date, but got TaskList back: [%s]", testTasklist)
	} else if err.Error() != `parsing time "2013-13-01": month out of range` {
		t.Error(err)
	}

	if testTasklist, err := LoadFromFilename(testInputTasklistDueDateError); testTasklist != nil || err == nil {
		t.Errorf("Expected LoadFromFilename to fail because of invalid due date, but got TaskList back: [%s]", testTasklist)
	} else if err.Error() != `parsing time "2014-02-32": day out of range` {
		t.Error(err)
	}

	if testTasklist, err := LoadFromFilename(testInputTasklistCompletedDateError); testTasklist != nil || err == nil {
		t.Errorf("Expected LoadFromFilename to fail because of invalid completed date, but got TaskList back: [%s]", testTasklist)
	} else if err.Error() != `parsing time "2014-25-04": month out of range` {
		t.Error(err)
	}

	// really silly test
	if testTasklist, err := LoadFromFilename(testInputTasklistScannerError); testTasklist != nil || err == nil {
		t.Errorf("Expected LoadFromFilename to fail because of invalid file, but got TaskList back: [%s]", testTasklist)
	} else if err.Error() != `bufio.Scanner: token too long` {
		t.Error(err)
	}
}
