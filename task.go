/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package todotxt

import (
	"fmt"
	"sort"
	"time"
)

// Task represents a todo.txt task entry.
type Task struct {
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
