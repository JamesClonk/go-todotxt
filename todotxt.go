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
	Original       string // Original raw task text
	Todo           string // Todo part of task text
	Priority       string
	Projects       []string
	Contexts       []string
	AdditionalTags map[string]string // Addon tags will be available here
	CreatedDate    time.Time
	DueDate        time.Time
	CompletedDate  time.Time
	Completed      bool
}

// TaskList represents a list of todo.txt task entries.
// It is usually loaded from a whole todo.txt file.
type TaskList []Task

var (
	// Used for formatting time.Time into todo.txt date format.
	DateLayout = "2006-01-02"

	// unexported vars
	priorityRx = regexp.MustCompile(`^(x|x \d{4}-\d{2}-\d{2}|)\s*\(([A-Z])\)\s+`) // Match priority: '(A) ...' or 'x (A) ...' or 'x 2012-12-12 (A) ...'
	// Match created date: '(A) 2012-12-12 ...' or 'x 2012-12-12 (A) 2012-12-12 ...' or 'x 2012-12-12 2012-12-12 ...' or '2012-12-12 ...'
	createdDateRx   = regexp.MustCompile(`^(\([A-Z]\)|x \d{4}-\d{2}-\d{2} \([A-Z]\)|x \d{4}-\d{2}-\d{2}|)\s*(\d{4}-\d{2}-\d{2})\s+`)
	completedDateRx = regexp.MustCompile(`^x\s*(\d{4}-\d{2}-\d{2})\s+`) // Match completed date: 'x 2012-12-12 ...'
	addonTagRx      = regexp.MustCompile(`(^|\s+)([\w-]+):(\S+)`)       // Match additional tags date: '... due:2012-12-12 ...'
	contextRx       = regexp.MustCompile(`(^|\s+)@(\S+)`)               // Match contexts: '@Context ...' or '... @Context ...'
	projectRx       = regexp.MustCompile(`(^|\s+)\+(\S+)`)              // Match projects: '+Project...' or '... +Project ...'
)

// String returns a complete tasklist string in todo.txt format.
func (tasklist *TaskList) String() (text string) {
	for _, task := range *tasklist {
		text += fmt.Sprintf("%s\n", task.String())
	}
	return text
}

// String returns a complete task string in todo.txt format.
//
// Contexts,  Projects and additional tags are alphabetically sorted,
// and appendend at the end in the following order:
// Contexts, Projects, Tags
//
// For example:
//  "(A) 2013-07-23 Call Dad @Home @Phone +Family due:2013-07-31 customTag1:Important!"
func (task *Task) String() string {
	var text string

	if task.Completed {
		text += "x "
		if task.HasCompletedDate() {
			text += fmt.Sprintf("%s ", task.CompletedDate.Format(DateLayout))
		}
	}

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

	if len(task.AdditionalTags) > 0 {
		// Sort map alphabetically by keys
		keys := make([]string, 0, len(task.AdditionalTags))
		for key, _ := range task.AdditionalTags {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			text += fmt.Sprintf(" %s:%s", key, task.AdditionalTags[key])
		}
	}

	if task.HasDueDate() {
		text += fmt.Sprintf(" due:%s", task.DueDate.Format(DateLayout))
	}

	return text
}

// Task returns a complete task string in todo.txt format.
// See *Task.String() for further information.
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

		// Ignore blank or comment lines
		if strings.Trim(task.Original, "\t\n\r ") == "" || strings.HasPrefix(task.Original, "#") {
			continue
		}

		// Check for completed
		if strings.HasPrefix(task.Original, "x ") {
			task.Completed = true
			// Check for completed date
			if completedDateRx.MatchString(task.Original) {
				if date, err := time.Parse(DateLayout, completedDateRx.FindStringSubmatch(task.Original)[1]); err != nil {
					return err
				} else {
					task.CompletedDate = date
				}
			}
		}

		// Check for priority
		if priorityRx.MatchString(task.Original) {
			task.Priority = priorityRx.FindStringSubmatch(task.Original)[2]
		}

		// Check for created date
		if createdDateRx.MatchString(task.Original) {
			if date, err := time.Parse(DateLayout, createdDateRx.FindStringSubmatch(task.Original)[2]); err != nil {
				return err
			} else {
				task.CreatedDate = date
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

		// Check for additional tags
		if addonTagRx.MatchString(task.Original) {
			matches := addonTagRx.FindAllStringSubmatch(task.Original, -1)
			tags := make(map[string]string, len(matches))
			for _, match := range matches {
				key, value := match[2], match[3]
				if key == "due" { // due date is a known addon tag, it has its own struct field
					if date, err := time.Parse(DateLayout, value); err != nil {
						return err
					} else {
						task.DueDate = date
					}
				} else if key != "" && value != "" {
					tags[key] = value
				}
			}
			task.AdditionalTags = tags
		}

		text := task.Original
		if task.Completed {
			text = text[2:] // Strip 'x '
			text = completedDateRx.ReplaceAllString(text, "")
		}
		// Remove all matching regular expressions from text, so only the actual todo text is left
		text = priorityRx.ReplaceAllString(text, "")
		text = createdDateRx.ReplaceAllString(text, "")
		text = contextRx.ReplaceAllString(text, "")
		text = projectRx.ReplaceAllString(text, "")
		text = addonTagRx.ReplaceAllString(text, "")
		task.Todo = strings.Trim(text, "\t\n\r\f ")

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
