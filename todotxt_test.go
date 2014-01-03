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
	testInput          = "testdata/input_todo.txt"
	testOutput         = "testdata/ouput_todo.txt"
	testExpectedOutput = "testdata/expected_todo.txt"
	testTasklist       TaskList
	testExpected       interface{}
	testGot            interface{}
)

func TestLoadFromFile(t *testing.T) {
	file, err := os.Open(testInput)
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
}

func TestLoadFromFilename(t *testing.T) {
	if testTasklist, err := LoadFromFilename(testInput); err != nil {
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
}

func TestWriteFile(t *testing.T) {
	os.Remove(testOutput)
	os.Create(testOutput)
	var err error

	fileInput, err := os.Open(testInput)
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

	fileInput, err := os.Open(testInput)
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

	if testTasklist, err = LoadFromFilename(testInput); err != nil {
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

	if err := testTasklist.LoadFromFilename(testInput); err != nil {
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
	testTasklist.LoadFromFilename(testInput)

	testExpected := 44
	testGot := len(testTasklist)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}
}

func TestTaskString(t *testing.T) {
	testTasklist.LoadFromFilename(testInput)
	taskId := 1

	testExpected = "2013-02-22 Pick up milk @GroceryStore"
	testGot = testTasklist[taskId-1].String()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = "x Download Todo.txt mobile app @Phone"
	testGot = testTasklist[taskId-1].String()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = "(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
	testGot = testTasklist[taskId-1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[taskId-1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = "x 2014-01-03 2014-01-01 Create some more golang library test cases @Go +go-todotxt"
	testGot = testTasklist[taskId-1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
	taskId++
}

func TestTaskPriority(t *testing.T) {
	testTasklist.LoadFromFilename(testInput)
	taskId := 6

	testExpected = "B"
	testGot = testTasklist[taskId-1].Priority
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have priority '%s', but got '%s'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = "C"
	testGot = testTasklist[taskId-1].Priority
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have priority '%s', but got '%s'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = "B"
	testGot = testTasklist[taskId-1].Priority
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have priority '%s', but got '%s'", taskId, testExpected, testGot)
	}
	taskId++

	if testTasklist[taskId-1].HasPriority() {
		t.Errorf("Expected Task[%d] to have no priority, but got '%s'", taskId, testTasklist[4].Priority)
	}
	taskId++
}

func TestTaskCreatedDate(t *testing.T) {
	testTasklist.LoadFromFilename(testInput)
	taskId := 10

	testExpected, err := time.Parse(DateLayout, "2012-01-30")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CreatedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected, err = time.Parse(DateLayout, "2013-02-22")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CreatedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected, err = time.Parse(DateLayout, "2013-12-30")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CreatedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected, err = time.Parse(DateLayout, "2014-01-01")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CreatedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	if testTasklist[taskId-1].HasCreatedDate() {
		t.Errorf("Expected Task[%d] to have no created date, but got '%v'", taskId, testTasklist[4].CreatedDate)
	}
	taskId++
}

func TestTaskContexts(t *testing.T) {
	testTasklist.LoadFromFilename(testInput)
	taskId := 15

	testExpected = []string{"Call", "Phone"}
	testGot = testTasklist[taskId-1].Contexts
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have contexts '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = []string{"Office"}
	testGot = testTasklist[taskId-1].Contexts
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have contexts '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = []string{"Electricity", "Home", "Of_Super-Importance", "Television"}
	testGot = testTasklist[taskId-1].Contexts
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have contexts '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = []string{}
	testGot = testTasklist[taskId-1].Contexts
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have no contexts, but got '%v'", taskId, testGot)
	}
	taskId++
}

func TestTasksProjects(t *testing.T) {
	testTasklist.LoadFromFilename(testInput)
	taskId := 19

	testExpected = []string{"Gardening", "Improving", "Planning", "Relaxing-Work"}
	testGot = testTasklist[taskId-1].Projects
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have projects '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = []string{"Novel"}
	testGot = testTasklist[taskId-1].Projects
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have projects '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = []string{}
	testGot = testTasklist[taskId-1].Projects
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have no projects, but got '%v'", taskId, testGot)
	}
	taskId++
}

func TestTaskDueDate(t *testing.T) {
	testTasklist.LoadFromFilename(testInput)
	taskId := 22

	testExpected, err := time.Parse(DateLayout, "2014-02-17")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].DueDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have due date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	if testTasklist[taskId-1].HasDueDate() {
		t.Errorf("Expected Task[%d] to have no due date, but got '%v'", taskId, testTasklist[taskId-1].DueDate)
	}
	taskId++
}

func TestTaskAddonTags(t *testing.T) {
	testTasklist.LoadFromFilename(testInput)
	taskId := 24

	testExpected = map[string]string{"Level": "5", "private": "false"}
	testGot = testTasklist[taskId-1].AdditionalTags
	if len(testGot.(map[string]string)) != 2 ||
		!compareMaps(testGot.(map[string]string), testExpected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have addon tags '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = map[string]string{"Importance": "Very!"}
	testGot = testTasklist[taskId-1].AdditionalTags
	if len(testGot.(map[string]string)) != 1 ||
		!compareMaps(testGot.(map[string]string), testExpected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have projects '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = map[string]string{}
	testGot = testTasklist[taskId-1].AdditionalTags
	if len(testGot.(map[string]string)) != 0 ||
		!compareMaps(testGot.(map[string]string), testExpected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have no additional tags, but got '%v'", taskId, testGot)
	}
	taskId++

	testExpected = map[string]string{}
	testGot = testTasklist[taskId-1].AdditionalTags
	if len(testGot.(map[string]string)) != 0 ||
		!compareMaps(testGot.(map[string]string), testExpected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have no additional tags, but got '%v'", taskId, testGot)
	}
	taskId++
}

func TestTaskCompleted(t *testing.T) {
	testTasklist.LoadFromFilename(testInput)
	taskId := 28

	testExpected = true
	testGot = testTasklist[taskId-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] not to completed, but got '%v'", taskId, testGot)
	}
	taskId++

	testExpected = true
	testGot = testTasklist[taskId-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] not to completed, but got '%v'", taskId, testGot)
	}
	taskId++

	testExpected = true
	testGot = testTasklist[taskId-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] not to completed, but got '%v'", taskId, testGot)
	}
	taskId++

	testExpected = false
	testGot = testTasklist[taskId-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] not to be completed, but got '%v'", taskId, testGot)
	}
	taskId++

	testExpected = false
	testGot = testTasklist[taskId-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] not to be completed, but got '%v'", taskId, testGot)
	}
	taskId++
}

func TestTaskCompletedDate(t *testing.T) {
	testTasklist.LoadFromFilename(testInput)
	taskId := 33

	if testTasklist[taskId-1].HasCompletedDate() {
		t.Errorf("Expected Task[%d] to not have a completed date, but got '%v'", taskId, testTasklist[taskId-1].CompletedDate)
	}
	taskId++

	testExpected, err := time.Parse(DateLayout, "2014-01-03")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CompletedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have completed date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	if testTasklist[taskId-1].HasCompletedDate() {
		t.Errorf("Expected Task[%d] to not have a completed date, but got '%v'", taskId, testTasklist[taskId-1].CompletedDate)
	}
	taskId++

	testExpected, err = time.Parse(DateLayout, "2014-01-02")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CompletedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have completed date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected, err = time.Parse(DateLayout, "2014-01-03")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CompletedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have completed date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	if testTasklist[taskId-1].HasCompletedDate() {
		t.Errorf("Expected Task[%d] to not have a completed date, but got '%v'", taskId, testTasklist[taskId-1].CompletedDate)
	}
	taskId++
}

func TestTaskSortByPriority(t *testing.T) {
	testTasklist.LoadFromFilename(testInput)
	taskId := 38

	testTasklist = testTasklist[taskId:]

	testTasklist.Sort(SORT_PRIORITY_ASC)

	testExpected = "(A) 2012-01-30 Call Mom @Call @Phone +Family"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "(D) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "2013-02-22 Pick up milk @GroceryStore"
	testGot = testTasklist[4].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[5] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05"
	testGot = testTasklist[5].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[6] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testTasklist.Sort(SORT_PRIORITY_DESC)

	testExpected = "x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "2013-02-22 Pick up milk @GroceryStore"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "(D) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[4].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[5] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "(A) 2012-01-30 Call Mom @Call @Phone +Family"
	testGot = testTasklist[5].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[6] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}
}

func compareSlices(list1 []string, list2 []string) bool {
	if len(list1) != len(list2) {
		return false
	}

	for i := range list1 {
		if list1[i] != list2[i] {
			return false
		}
	}

	return true
}

func compareMaps(map1 map[string]string, map2 map[string]string) bool {
	if len(map1) != len(map2) {
		return false
	}

	compare := func(map1 map[string]string, map2 map[string]string) bool {
		for key, value := range map1 {
			if value2, found := map2[key]; !found {
				return false
			} else if value != value2 {
				return false
			}
		}
		return true
	}

	return compare(map1, map2) && compare(map2, map1)
}
