/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

// Package todotxt is a Go client library for Gina Trapani's todo.txt files.
// It allows for parsing and manipulating of task lists and tasks in the todo.txt format.
//
// Source code and project home: https://github.com/JamesClonk/go-todotxt
package todotxt

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

// Task represents a todo.txt task entry.
type Task struct {
	Original      string // Original raw task text
	Todo          string // Todo part of task text
	Priority      string
	Projects      []string
	Contexts      []string
	CreatedDate   time.Time
	DueDate       time.Time
	CompletedDate time.Time
	Completed     bool
}

// TaskList represents a list of todo.txt task entries.
// It is usually loaded from a whole todo.txt file.
type TaskList []Task

var (
	// Used for formatting time.Time into todo.txt date format.
	DateLayout = "2006-01-02"

	// unexported vars
	priorityRx    = regexp.MustCompile(`^\(([A-Z])\)\s+`)                              // Match priority value: '(A) ...'
	createdDateRx = regexp.MustCompile(`^(\([A-Z]\)|)\s*([\d]{4}-[\d]{2}-[\d]{2})\s+`) // Match date value: '(A) 2012-12-12 ...' or '2012-12-12 ...'
)

// String returns a complete task string in todo.txt format.
//
// For example:
//  "(A) 2013-07-23 Call Dad @Phone +Family due:2013-07-31"
func (task *Task) String() string {
	var text string
	if task.HasPriority() {
		text += fmt.Sprintf("(%s) ", task.Priority)
	}
	if task.HasCreatedDate() {
		text += fmt.Sprintf("%s ", task.CreatedDate.Format(DateLayout))
	}
	text += task.Todo
	if task.HasDueDate() {
		text += fmt.Sprintf(" %s", task.DueDate.Format(DateLayout))
	}
	return text
}

// Task returns a complete task string in todo.txt format.
// The same as *Task.String().
func (task *Task) Task() string {
	return task.String()
}

// HasPriority returns true if the task has a priority.
func (task *Task) HasPriority() bool {
	return task.Priority != ""
}

// HasCreatedDate returns true if the task has a created date.
func (task *Task) HasCreatedDate() bool {
	return !task.CreatedDate.IsZero()
}

// HasDueDate returns true if the task has a due date.
func (task *Task) HasDueDate() bool {
	return !task.DueDate.IsZero()
}

// HasCompletedDate returns true if the task has a completed date.
func (task *Task) HasCompletedDate() bool {
	return !task.CompletedDate.IsZero()
}

// LoadFromFile loads a TaskList from *os.File.
//
// Using *os.File instead of a filename allows to also use os.Stdin.
//
// Note: This will clear the current TaskList and overwrite it's contents with whatever is in *os.File.
func (tasklist *TaskList) LoadFromFile(file *os.File) error {
	*tasklist = []Task{} // Empty tasklist

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		task := Task{}
		task.Original = scanner.Text()

		// Check for priority
		if priorityRx.MatchString(task.Original) {
			task.Priority = priorityRx.FindStringSubmatch(task.Original)[1] // First match is priority value
		}

		// Check for created date
		if createdDateRx.MatchString(task.Original) {
			// Second match is created date value
			if date, err := time.Parse(DateLayout, createdDateRx.FindStringSubmatch(task.Original)[2]); err != nil {
				return err
			} else {
				task.CreatedDate = date
			}
		}

		*tasklist = append(*tasklist, task)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// LoadFromFilename loads a TaskList from a file (most likely called "todo.txt").
//
// Note: This will clear the current TaskList and overwrite it's contents with whatever is in the file.
func (tasklist *TaskList) LoadFromFilename(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return tasklist.LoadFromFile(file)
}

// LoadFromFile loads and returns a TaskList from *os.File.
//
// Using *os.File instead of a filename allows to also use os.Stdin.
func LoadFromFile(file *os.File) (*TaskList, error) {
	tasklist := &TaskList{}
	err := tasklist.LoadFromFile(file)
	return tasklist, err
}

// LoadFromFilename loads and returns a TaskList from a file (most likely called "todo.txt").
func LoadFromFilename(filename string) (*TaskList, error) {
	tasklist := &TaskList{}
	err := tasklist.LoadFromFilename(filename)
	return tasklist, err
}
