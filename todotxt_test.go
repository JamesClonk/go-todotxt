/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package todotxt

import (
	"os"
	"testing"
	"time"
)

func TestLoadFromFile(t *testing.T) {
	file, err := os.Open("todo.txt")
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
	if tasklist, err := LoadFromFilename("todo.txt"); err != nil {
		t.Fatal(err)
	} else {
		loadTest(t, *tasklist)
	}
}

func loadTest(t *testing.T, tasklist TaskList) {
	var expected, got interface{}
	var err error

	// -------------------------------------------------------------------------------------
	// count tasks
	expected = 9
	got = len(tasklist)
	if got != expected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", expected, got)
	}

	// -------------------------------------------------------------------------------------
	// complete task strings
	expected = "x Download Todo.txt mobile app @Phone"
	got = tasklist[7].String()
	if got != expected {
		t.Errorf("Expected eight Task to be [%s], but got [%s]", expected, got)
	}

	expected = "(B) 2013-12-01 Outline chapter 5 @Computer +Novel due:2014-01-01"
	got = tasklist[2].Task()
	if got != expected {
		t.Errorf("Expected third Task to be [%s], but got [%s]", expected, got)
	}

	// -------------------------------------------------------------------------------------
	// task priority
	expected = "B"
	got = tasklist[2].Priority
	if got != expected {
		t.Errorf("Expected third task to have priority '%s', but got '%s'", expected, got)
	}

	if tasklist[4].HasPriority() {
		t.Errorf("Expected fifth task to have no priority, but got '%s'", tasklist[4].Priority)
	}

	// -------------------------------------------------------------------------------------
	// task created date
	expected, err = time.Parse(DateLayout, "2012-01-30")
	if err != nil {
		t.Fatal(err)
	}
	got = tasklist[0].CreatedDate
	if got != expected {
		t.Errorf("Expected first task to have created date '%s', but got '%v'", expected, got)
	}

	expected, err = time.Parse(DateLayout, "2013-02-22")
	if err != nil {
		t.Fatal(err)
	}
	got = tasklist[5].CreatedDate
	if got != expected {
		t.Errorf("Expected sixth task to have created date '%s', but got '%v'", expected, got)
	}

	if tasklist[4].HasCreatedDate() {
		t.Errorf("Expected fifth task to have no created date, but got '%v'", tasklist[4].CreatedDate)
	}

	// -------------------------------------------------------------------------------------
	// task contexts
	expected = []string{"Call", "Phone"}
	got = tasklist[0].Contexts
	if !compareSlices(got.([]string), expected.([]string)) {
		t.Errorf("Expected first task to have contexts '%v', but got '%v'", expected, got)
	}

	expected = []string{"Office"}
	got = tasklist[3].Contexts
	if !compareSlices(got.([]string), expected.([]string)) {
		t.Errorf("Expected fourth task to have contexts '%v', but got '%v'", expected, got)
	}

	expected = []string{"Electricity", "Home", "Television"}
	got = tasklist[8].Contexts
	if !compareSlices(got.([]string), expected.([]string)) {
		t.Errorf("Expected ninth task to have contexts '%v', but got '%v'", expected, got)
	}

	expected = []string{}
	got = tasklist[6].Contexts
	if !compareSlices(got.([]string), expected.([]string)) {
		t.Errorf("Expected seventh task to have no contexts, but got '%v'", got)
	}

	// -------------------------------------------------------------------------------------
	// task projects
	expected = []string{"Gardening", "Improving", "Planning"}
	got = tasklist[4].Projects
	if !compareSlices(got.([]string), expected.([]string)) {
		t.Errorf("Expected fifth task to have projects '%v', but got '%v'", expected, got)
	}

	expected = []string{"Novel"}
	got = tasklist[6].Projects
	if !compareSlices(got.([]string), expected.([]string)) {
		t.Errorf("Expected sixth task to have projects '%v', but got '%v'", expected, got)
	}

	expected = []string{}
	got = tasklist[8].Projects
	if !compareSlices(got.([]string), expected.([]string)) {
		t.Errorf("Expected ninth task to have no projects, but got '%v'", got)
	}

	// -------------------------------------------------------------------------------------
	// task due date
	expected, err = time.Parse(DateLayout, "2014-02-17")
	if err != nil {
		t.Fatal(err)
	}
	got = tasklist[2].DueDate
	if got != expected {
		t.Errorf("Expected third task to have due date '%s', but got '%v'", expected, got)
	}

	if tasklist[0].HasDueDate() {
		t.Errorf("Expected first task to have no due date, but got '%v'", tasklist[0].DueDate)
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
