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
	"sort"
	"strings"
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
	priorityRx    = regexp.MustCompile(`^\(([A-Z])\)\s+`)                              // Match priority: '(A) ...'
	createdDateRx = regexp.MustCompile(`^(\([A-Z]\)|)\s*([\d]{4}-[\d]{2}-[\d]{2})\s+`) // Match date: '(A) 2012-12-12 ...' or '2012-12-12 ...'
	dueDateRx     = regexp.MustCompile(`\s+due:([\d]{4}-[\d]{2}-[\d]{2})`)             // Match due date: '... due:2012-12-12 ...'
	contextRx     = regexp.MustCompile(`(^|\s+)@([[:word:]]+)`)                        // Match contexts: '@Context ...' or '... @Context ...'
	projectRx     = regexp.MustCompile(`(^|\s+)\+([[:word:]]+)`)                       // Match projects: '+Project...' or '... +Project ...'
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

	if len(task.Contexts) > 0 {
		for _, context := range task.Contexts {
			text += fmt.Sprintf(" @%s", context)
		}
	}

	if len(task.Projects) > 0 {
		for _, project := range task.Projects {
			text += fmt.Sprintf(" +%s", project)
		}
	}

	if task.HasDueDate() {
		text += fmt.Sprintf(" due:%s", task.DueDate.Format(DateLayout))
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
			task.Priority = priorityRx.FindStringSubmatch(task.Original)[1]
		}

		// Check for created date
		if createdDateRx.MatchString(task.Original) {
			if date, err := time.Parse(DateLayout, createdDateRx.FindStringSubmatch(task.Original)[2]); err != nil {
				return err
			} else {
				task.CreatedDate = date
			}
		}

		// Check for due date
		if dueDateRx.MatchString(task.Original) {
			if date, err := time.Parse(DateLayout, dueDateRx.FindStringSubmatch(task.Original)[1]); err != nil {
				return err
			} else {
				task.DueDate = date
			}
		}

		// function for collecting projects/contexts as slices from text
		getSlice := func(rx *regexp.Regexp) []string {
			matches := rx.FindAllStringSubmatch(task.Original, -1)
			slice := make([]string, 0, len(matches))
			seen := make(map[string]bool, len(matches))
			for _, match := range matches {
				word := strings.Trim(match[2], "\t\n\r ")
				if _, found := seen[word]; !found {
					slice = append(slice, word)
					seen[word] = true
				}
			}
			sort.Strings(slice)
			return slice
		}

		// Check for contexts
		if contextRx.MatchString(task.Original) {
			task.Contexts = getSlice(contextRx)
		}

		// Check for projects
		if projectRx.MatchString(task.Original) {
			task.Projects = getSlice(projectRx)
		}

		// Todo text
		// use replacer function here.. strip all other fields from task.Original, then left+right trim --> todo text
		text := strings.Replace(task.Original, " ", "", -1)
		task.Todo = text

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
