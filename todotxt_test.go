/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package todotxt

import (
	"io/ioutil"
	"os"
	"testing"
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
		testExpected := string(data)
		testGot := testTasklist.String()
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
		testExpected := string(data)
		testGot := testTasklist.String()
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
	testExpected := string(data)
	testGot := testTasklist.String()
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
	testExpected := string(data)
	testGot := testTasklist.String()
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
	testExpected := string(data)
	testGot := testTasklist.String()
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
	testExpected := string(data)
	testGot := testTasklist.String()
	if testGot != testExpected {
		t.Errorf("Expected TaskList to be [%s], but got [%s]", testExpected, testGot)
	}
}

func TestTaskListCount(t *testing.T) {
	if err := testTasklist.LoadFromFilename(testInputTasklist); err != nil {
		t.Fatal(err)
	}

	testExpected := 63
	testGot := len(testTasklist)
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
