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

	var tasklist TaskList
	if err := tasklist.LoadFromFile(file); err != nil {
		t.Fatal(err)
	}
	loadTest(t, tasklist)
}

func TestLoadFromFilename(t *testing.T) {
	var tasklist TaskList
	if err := tasklist.LoadFromFilename("todo.txt"); err != nil {
		t.Fatal(err)
	}
	loadTest(t, tasklist)
}

func loadTest(t *testing.T, tasklist TaskList) {
	var expected, got interface{}
	var err error

	expected = 8
	got = len(tasklist)
	if got != expected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", expected, got)
	}

	expected = "B"
	got = tasklist[2].Priority
	if got != expected {
		t.Errorf("Expected third task to have priority '%s', but got '%s'", expected, got)
	}

	if tasklist[4].HasPriority() {
		t.Errorf("Expected fifth task to have no priority, but got '%s'", tasklist[4].Priority)
	}

	expected, err = time.Parse("2006-01-02", "2012-01-30")
	if err != nil {
		t.Fatal(err)
	}
	got = tasklist[0].CreatedDate
	if got != expected {
		t.Errorf("Expected first task to have created date '%s', but got '%v'", expected, got)
	}

	expected, err = time.Parse("2006-01-02", "2013-02-22")
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
}
