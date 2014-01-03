/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package todotxt

import (
	"bufio"
	"os"
	"regexp"
	"time"
)

type Task struct {
	Task          string // Raw text
	Todo          string // Only actual todo part of text
	Priority      string
	Projects      []string
	Contexts      []string
	CreatedDate   time.Time
	DueDate       time.Time
	CompletedDate time.Time
	Completed     bool
}

type TaskList []Task

var (
	priorityRx    = regexp.MustCompile(`^\(([A-Z])\)\s+`)                              // Match priority value: '(A) ...'
	createdDateRx = regexp.MustCompile(`^(\([A-Z]\)|)\s*([\d]{4}-[\d]{2}-[\d]{2})\s+`) // Match date value: '(A) 2012-12-12 ...' or '2012-12-12 ...'
)

// Return raw task text for String()
func (task *Task) String() string {
	return task.Task
}

func (task *Task) HasPriority() bool {
	return task.Priority != ""
}

func (task *Task) HasCreatedDate() bool {
	return !task.CreatedDate.IsZero()
}

func (task *Task) HasDueDate() bool {
	return !task.DueDate.IsZero()
}

func (task *Task) HasCompletedDate() bool {
	return !task.CompletedDate.IsZero()
}

// Loading from *os.File allows to also use os.Stdin instead of just actual files
func (tasklist *TaskList) LoadFromFile(file *os.File) error {
	*tasklist = []Task{} // Reset tasklist

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		task := Task{}
		task.Task = scanner.Text()

		// Check for priority
		if priorityRx.MatchString(task.Task) {
			task.Priority = priorityRx.FindStringSubmatch(task.Task)[1] // First match is priority value
		}

		// Check for created date
		if createdDateRx.MatchString(task.Task) {
			// Second match is created date value
			if date, err := time.Parse("2006-01-02", createdDateRx.FindStringSubmatch(task.Task)[2]); err != nil {
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

// Convenience method, since most of the time tasks will be loaded from an actual file, called "todo.txt" most likely ;)
func (tasklist *TaskList) LoadFromFilename(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return tasklist.LoadFromFile(file)
}
