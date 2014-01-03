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

func TestLoadFromFile(t *testing.T) {
	file, err := os.Open("testdata/input_todo.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if tasklist, err := LoadFromFile(file); err != nil {
		t.Fatal(err)
	} else {
		loadTest(t, *tasklist)
	}
}

func TestLoadFromFilename(t *testing.T) {
	if tasklist, err := LoadFromFilename("testdata/input_todo.txt"); err != nil {
		t.Fatal(err)
	} else {
		loadTest(t, *tasklist)
	}
}

func loadTest(t *testing.T, tasklist TaskList) {
	taskId := 1
	var expected, got interface{}
	var err error

	// -------------------------------------------------------------------------------------
	// complete tasklist string
	data, err := ioutil.ReadFile("testdata/expected_todo.txt")
	if err != nil {
		t.Fatal(err)
	}
	expected = string(data)
	got = tasklist.String()
	if got != expected {
		//t.Errorf("Expected TaskList to be [%s], but got [%s]", expected, got)--------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	}

	// -------------------------------------------------------------------------------------
	// count tasks
	expected = 10
	got = len(tasklist)
	if got != expected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", expected, got)
	}

	// -------------------------------------------------------------------------------------
	// complete task strings
	expected = "2013-02-22 Pick up milk @GroceryStore"
	got = tasklist[taskId-1].String()
	if got != expected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, expected, got)
	}
	taskId++

	expected = "x Download Todo.txt mobile app @Phone"
	got = tasklist[taskId-1].String()
	if got != expected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, expected, got)
	}
	taskId++

	expected = "(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
	got = tasklist[taskId-1].Task()
	if got != expected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, expected, got)
	}
	taskId++

	expected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	got = tasklist[taskId-1].Task()
	if got != expected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, expected, got)
	}
	expected = "1"
	got = tasklist[taskId-1].Todo
	if got != expected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, expected, got)
	}
	taskId++

	expected = "x 2014-01-03 2014-01-01 Create some more golang library test cases @Go +go-todotxt"
	got = tasklist[taskId-1].Task()
	if got != expected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, expected, got)
	}
	taskId++

	// -------------------------------------------------------------------------------------
	// task priority
	expected = "B"
	got = tasklist[taskId-1].Priority
	if got != expected {
		t.Errorf("Expected Task[%d] to have priority '%s', but got '%s'", taskId, expected, got)
	}
	taskId++

	expected = "C"
	got = tasklist[taskId-1].Priority
	if got != expected {
		t.Errorf("Expected Task[%d] to have priority '%s', but got '%s'", taskId, expected, got)
	}
	taskId++

	expected = "B"
	got = tasklist[taskId-1].Priority
	if got != expected {
		t.Errorf("Expected Task[%d] to have priority '%s', but got '%s'", taskId, expected, got)
	}
	taskId++

	if tasklist[taskId-1].HasPriority() {
		t.Errorf("Expected Task[%d] to have no priority, but got '%s'", taskId, tasklist[4].Priority)
	}
	taskId++

	// -------------------------------------------------------------------------------------
	// task created date
	expected, err = time.Parse(DateLayout, "2012-01-30")
	if err != nil {
		t.Fatal(err)
	}
	got = tasklist[taskId-1].CreatedDate
	if got != expected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskId, expected, got)
	}
	taskId++

	expected, err = time.Parse(DateLayout, "2013-02-22")
	if err != nil {
		t.Fatal(err)
	}
	got = tasklist[taskId-1].CreatedDate
	if got != expected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskId, expected, got)
	}
	taskId++

	expected, err = time.Parse(DateLayout, "2013-12-30")
	if err != nil {
		t.Fatal(err)
	}
	got = tasklist[taskId-1].CreatedDate
	if got != expected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskId, expected, got)
	}
	taskId++

	expected, err = time.Parse(DateLayout, "2014-01-01")
	if err != nil {
		t.Fatal(err)
	}
	got = tasklist[taskId-1].CreatedDate
	if got != expected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskId, expected, got)
	}
	taskId++

	if tasklist[taskId-1].HasCreatedDate() {
		t.Errorf("Expected Task[%d] to have no created date, but got '%v'", taskId, tasklist[4].CreatedDate)
	}
	taskId++

	// -------------------------------------------------------------------------------------
	// task contexts
	expected = []string{"Call", "Phone"}
	got = tasklist[taskId-1].Contexts
	if !compareSlices(got.([]string), expected.([]string)) {
		t.Errorf("Expected Task[%d] to have contexts '%v', but got '%v'", taskId, expected, got)
	}
	taskId++

	expected = []string{"Office"}
	got = tasklist[taskId-1].Contexts
	if !compareSlices(got.([]string), expected.([]string)) {
		t.Errorf("Expected Task[%d] to have contexts '%v', but got '%v'", taskId, expected, got)
	}
	taskId++

	expected = []string{"Electricity", "Home", "Television"}
	got = tasklist[taskId-1].Contexts
	if !compareSlices(got.([]string), expected.([]string)) {
		t.Errorf("Expected Task[%d] to have contexts '%v', but got '%v'", taskId, expected, got)
	}
	taskId++

	expected = []string{}
	got = tasklist[taskId-1].Contexts
	if !compareSlices(got.([]string), expected.([]string)) {
		t.Errorf("Expected Task[%d] to have no contexts, but got '%v'", taskId, got)
	}
	taskId++

	// -------------------------------------------------------------------------------------
	// task projects
	expected = []string{"Gardening", "Improving", "Planning"}
	got = tasklist[taskId-1].Projects
	if !compareSlices(got.([]string), expected.([]string)) {
		t.Errorf("Expected Task[%d] to have projects '%v', but got '%v'", taskId, expected, got)
	}
	taskId++

	expected = []string{"Novel"}
	got = tasklist[taskId-1].Projects
	if !compareSlices(got.([]string), expected.([]string)) {
		t.Errorf("Expected Task[%d] to have projects '%v', but got '%v'", taskId, expected, got)
	}
	taskId++

	expected = []string{}
	got = tasklist[taskId-1].Projects
	if !compareSlices(got.([]string), expected.([]string)) {
		t.Errorf("Expected Task[%d] to have no projects, but got '%v'", taskId, got)
	}
	taskId++

	// -------------------------------------------------------------------------------------
	// task due date
	expected, err = time.Parse(DateLayout, "2014-02-17")
	if err != nil {
		t.Fatal(err)
	}
	got = tasklist[taskId-1].DueDate
	if got != expected {
		t.Errorf("Expected Task[%d] to have due date '%s', but got '%v'", taskId, expected, got)
	}
	taskId++

	if tasklist[taskId-1].HasDueDate() {
		t.Errorf("Expected Task[%d] to have no due date, but got '%v'", taskId, tasklist[taskId-1].DueDate)
	}
	taskId++

	// -------------------------------------------------------------------------------------
	// task addon tags
	expected = map[string]string{"Level": "5", "private": "false"}
	got = tasklist[taskId-1].AdditionalTags
	if len(got.(map[string]string)) != 2 ||
		!compareMaps(got.(map[string]string), expected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have addon tags '%v', but got '%v'", taskId, expected, got)
	}
	taskId++

	expected = map[string]string{"Importance": "Very!"}
	got = tasklist[taskId-1].AdditionalTags
	if len(got.(map[string]string)) != 1 ||
		!compareMaps(got.(map[string]string), expected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have projects '%v', but got '%v'", taskId, expected, got)
	}
	taskId++

	expected = map[string]string{}
	got = tasklist[taskId-1].AdditionalTags
	if len(got.(map[string]string)) != 0 ||
		!compareMaps(got.(map[string]string), expected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have no additional tags, but got '%v'", taskId, got)
	}
	taskId++

	expected = map[string]string{}
	got = tasklist[taskId-1].AdditionalTags
	if len(got.(map[string]string)) != 0 ||
		!compareMaps(got.(map[string]string), expected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have no additional tags, but got '%v'", taskId, got)
	}
	taskId++

	// -------------------------------------------------------------------------------------
	// task completed
	expected = true
	got = tasklist[taskId-1].Completed
	if got != expected {
		t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskId, got)
	}
	taskId++

	expected = true
	got = tasklist[taskId-1].Completed
	if got != expected {
		t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskId, got)
	}
	taskId++

	expected = true
	got = tasklist[taskId-1].Completed
	if got != expected {
		t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskId, got)
	}
	taskId++

	expected = false
	got = tasklist[taskId-1].Completed
	if got != expected {
		t.Errorf("Expected Task[%d] to not be completed, but got '%v'", taskId, got)
	}
	taskId++

	expected = false
	got = tasklist[taskId-1].Completed
	if got != expected {
		t.Errorf("Expected Task[%d] to not be completed, but got '%v'", taskId, got)
	}
	taskId++
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
