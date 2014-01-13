/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package todotxt

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

var (
	// DateLayout is used for formatting time.Time into todo.txt date format and vice-versa.
	DateLayout = "2006-01-02"

	priorityRx = regexp.MustCompile(`^(x|x \d{4}-\d{2}-\d{2}|)\s*\(([A-Z])\)\s+`) // Match priority: '(A) ...' or 'x (A) ...' or 'x 2012-12-12 (A) ...'
	// Match created date: '(A) 2012-12-12 ...' or 'x 2012-12-12 (A) 2012-12-12 ...' or 'x (A) 2012-12-12 ...'or 'x 2012-12-12 2012-12-12 ...' or '2012-12-12 ...'
	createdDateRx   = regexp.MustCompile(`^(\([A-Z]\)|x \d{4}-\d{2}-\d{2} \([A-Z]\)|x \([A-Z]\)|x \d{4}-\d{2}-\d{2}|)\s*(\d{4}-\d{2}-\d{2})\s+`)
	completedRx     = regexp.MustCompile(`^x\s+`)                       // Match completed: 'x ...'
	completedDateRx = regexp.MustCompile(`^x\s*(\d{4}-\d{2}-\d{2})\s+`) // Match completed date: 'x 2012-12-12 ...'
	addonTagRx      = regexp.MustCompile(`(^|\s+)([\w-]+):(\S+)`)       // Match additional tags date: '... due:2012-12-12 ...'
	contextRx       = regexp.MustCompile(`(^|\s+)@(\S+)`)               // Match contexts: '@Context ...' or '... @Context ...'
	projectRx       = regexp.MustCompile(`(^|\s+)\+(\S+)`)              // Match projects: '+Project...' or '... +Project ...')
)

// Task represents a todo.txt task entry.
type Task struct {
	Id             int    // Internal task id.
	Original       string // Original raw task text.
	Todo           string // Todo part of task text.
	Priority       string
	Projects       []string
	Contexts       []string
	AdditionalTags map[string]string // Addon tags will be available here.
	CreatedDate    time.Time
	DueDate        time.Time
	CompletedDate  time.Time
	Completed      bool
}

// String returns a complete task string in todo.txt format.
//
// Contexts,  Projects and additional tags are alphabetically sorted,
// and appendend at the end in the following order:
// Contexts, Projects, Tags
//
// For example:
//  "(A) 2013-07-23 Call Dad @Home @Phone +Family due:2013-07-31 customTag1:Important!"
func (task Task) String() string {
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
		sort.Strings(task.Contexts)
		for _, context := range task.Contexts {
			text += fmt.Sprintf(" @%s", context)
		}
	}

	if len(task.Projects) > 0 {
		sort.Strings(task.Projects)
		for _, project := range task.Projects {
			text += fmt.Sprintf(" +%s", project)
		}
	}

	if len(task.AdditionalTags) > 0 {
		// Sort map alphabetically by keys
		keys := make([]string, 0, len(task.AdditionalTags))
		for key := range task.AdditionalTags {
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

// NewTask creates a new empty Task with default values. (CreatedDate is set to Now())
func NewTask() Task {
	task := Task{}
	task.CreatedDate = time.Now()
	return task
}

// ParseTask parses the input text string into a Task struct.
func ParseTask(text string) (*Task, error) {
	var err error

	task := Task{}
	task.Original = strings.Trim(text, "\t\n\r ")
	task.Todo = task.Original

	// Check for completed
	if completedRx.MatchString(task.Original) {
		task.Completed = true
		// Check for completed date
		if completedDateRx.MatchString(task.Original) {
			if date, err := time.Parse(DateLayout, completedDateRx.FindStringSubmatch(task.Original)[1]); err == nil {
				task.CompletedDate = date
			} else {
				return nil, err
			}
		}

		// Remove from Todo text
		task.Todo = completedDateRx.ReplaceAllString(task.Todo, "") // Strip CompletedDate first, otherwise it wouldn't match anymore (^x date...)
		task.Todo = completedRx.ReplaceAllString(task.Todo, "")     // Strip 'x '
	}

	// Check for priority
	if priorityRx.MatchString(task.Original) {
		task.Priority = priorityRx.FindStringSubmatch(task.Original)[2]
		task.Todo = priorityRx.ReplaceAllString(task.Todo, "") // Remove from Todo text
	}

	// Check for created date
	if createdDateRx.MatchString(task.Original) {
		if date, err := time.Parse(DateLayout, createdDateRx.FindStringSubmatch(task.Original)[2]); err == nil {
			task.CreatedDate = date
			task.Todo = createdDateRx.ReplaceAllString(task.Todo, "") // Remove from Todo text
		} else {
			return nil, err
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
		task.Todo = contextRx.ReplaceAllString(task.Todo, "") // Remove from Todo text
	}

	// Check for projects
	if projectRx.MatchString(task.Original) {
		task.Projects = getSlice(projectRx)
		task.Todo = projectRx.ReplaceAllString(task.Todo, "") // Remove from Todo text
	}

	// Check for additional tags
	if addonTagRx.MatchString(task.Original) {
		matches := addonTagRx.FindAllStringSubmatch(task.Original, -1)
		tags := make(map[string]string, len(matches))
		for _, match := range matches {
			key, value := match[2], match[3]
			if key == "due" { // due date is a known addon tag, it has its own struct field
				if date, err := time.Parse(DateLayout, value); err == nil {
					task.DueDate = date
				} else {
					return nil, err
				}
			} else if key != "" && value != "" {
				tags[key] = value
			}
		}
		task.AdditionalTags = tags
		task.Todo = addonTagRx.ReplaceAllString(task.Todo, "") // Remove from Todo text
	}

	// Trim any remaining whitespaces from Todo text
	task.Todo = strings.Trim(task.Todo, "\t\n\r\f ")

	return &task, err
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
	return !task.CompletedDate.IsZero() && task.Completed
}

// Complete sets Task.Completed to 'true' if the task was not already completed.
// Also sets Task.CompletedDate to time.Now()
func (task *Task) Complete() {
	if !task.Completed {
		task.Completed = true
		task.CompletedDate = time.Now()
	}
}

// Reopen sets Task.Completed to 'false' if the task was completed.
// Also resets Task.CompletedDate.
func (task *Task) Reopen() {
	if task.Completed {
		task.Completed = false
		task.CompletedDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC) // time.IsZero() value
	}
}

// IsOverdue returns true if due date is in the past.
//
// This function does not take the Completed flag into consideration.
// You should check Task.Completed first if needed.
func (task *Task) IsOverdue() bool {
	if task.HasDueDate() {
		return task.DueDate.Before(time.Now())
	}
	return false
}

// Due returns the duration passed since due date, or until due date from now.
// Check with IsOverdue() if the task is overdue or not.
//
// Just as with IsOverdue(), this function does also not take the Completed flag into consideration.
// You should check Task.Completed first if needed.
func (task *Task) Due() time.Duration {
	if task.IsOverdue() {
		return time.Now().Sub(task.DueDate)
	}
	return task.DueDate.Sub(time.Now())
}
