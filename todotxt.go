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
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

// TaskList represents a list of todo.txt task entries.
// It is usually loaded from a whole todo.txt file.
type TaskList []Task

// IgnoreComments can be set to 'false', in order to revert to more standard todo.txt behaviour.
// The todo.txt format does not define comments.
var (
	// Used for formatting time.Time into todo.txt date format and vice-versa.
	DateLayout = "2006-01-02"
	// Ignores comments (Lines/Text starting with "#").
	IgnoreComments = true

	// unexported vars
	priorityRx = regexp.MustCompile(`^(x|x \d{4}-\d{2}-\d{2}|)\s*\(([A-Z])\)\s+`) // Match priority: '(A) ...' or 'x (A) ...' or 'x 2012-12-12 (A) ...'
	// Match created date: '(A) 2012-12-12 ...' or 'x 2012-12-12 (A) 2012-12-12 ...' or 'x (A) 2012-12-12 ...'or 'x 2012-12-12 2012-12-12 ...' or '2012-12-12 ...'
	createdDateRx   = regexp.MustCompile(`^(\([A-Z]\)|x \d{4}-\d{2}-\d{2} \([A-Z]\)|x \([A-Z]\)|x \d{4}-\d{2}-\d{2}|)\s*(\d{4}-\d{2}-\d{2})\s+`)
	completedRx     = regexp.MustCompile(`^x\s+`)                       // Match completed: 'x ...'
	completedDateRx = regexp.MustCompile(`^x\s*(\d{4}-\d{2}-\d{2})\s+`) // Match completed date: 'x 2012-12-12 ...'
	addonTagRx      = regexp.MustCompile(`(^|\s+)([\w-]+):(\S+)`)       // Match additional tags date: '... due:2012-12-12 ...'
	contextRx       = regexp.MustCompile(`(^|\s+)@(\S+)`)               // Match contexts: '@Context ...' or '... @Context ...'
	projectRx       = regexp.MustCompile(`(^|\s+)\+(\S+)`)              // Match projects: '+Project...' or '... +Project ...'
)

// String returns a complete tasklist string in todo.txt format.
func (tasklist TaskList) String() (text string) {
	for _, task := range tasklist {
		text += fmt.Sprintf("%s\n", task.String())
	}
	return text
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
		task.Original = strings.Trim(scanner.Text(), "\t\n\r ") // Read line
		task.Todo = task.Original

		// Ignore blank or comment lines
		if task.Todo == "" || (IgnoreComments && strings.HasPrefix(task.Todo, "#")) {
			continue
		}

		// Check for completed
		if completedRx.MatchString(task.Original) {
			task.Completed = true
			// Check for completed date
			if completedDateRx.MatchString(task.Original) {
				if date, err := time.Parse(DateLayout, completedDateRx.FindStringSubmatch(task.Original)[1]); err != nil {
					return err
				} else {
					task.CompletedDate = date
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
			if date, err := time.Parse(DateLayout, createdDateRx.FindStringSubmatch(task.Original)[2]); err != nil {
				return err
			} else {
				task.CreatedDate = date
				task.Todo = createdDateRx.ReplaceAllString(task.Todo, "") // Remove from Todo text
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
			task.Todo = addonTagRx.ReplaceAllString(task.Todo, "") // Remove from Todo text
		}

		// Trim any remaining whitespaces from Todo text
		task.Todo = strings.Trim(task.Todo, "\t\n\r\f ")

		*tasklist = append(*tasklist, task)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// WriteToFile writes a TaskList to *os.File.
//
// Using *os.File instead of a filename allows to also use os.Stdout.
func (tasklist *TaskList) WriteToFile(file *os.File) error {
	writer := bufio.NewWriter(file)
	_, err := writer.WriteString(tasklist.String())
	writer.Flush()
	return err
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

// WriteToFilename writes a TaskList to the specified file (most likely called "todo.txt").
func (tasklist *TaskList) WriteToFilename(filename string) error {
	return ioutil.WriteFile(filename, []byte(tasklist.String()), 0644)
}

// LoadFromFile loads and returns a TaskList from *os.File.
//
// Using *os.File instead of a filename allows to also use os.Stdin.
func LoadFromFile(file *os.File) (TaskList, error) {
	tasklist := TaskList{}
	err := tasklist.LoadFromFile(file)
	return tasklist, err
}

// WriteToFile writes a TaskList to *os.File.
//
// Using *os.File instead of a filename allows to also use os.Stdout.
func WriteToFile(tasklist *TaskList, file *os.File) error {
	return tasklist.WriteToFile(file)
}

// LoadFromFilename loads and returns a TaskList from a file (most likely called "todo.txt").
func LoadFromFilename(filename string) (TaskList, error) {
	tasklist := TaskList{}
	err := tasklist.LoadFromFilename(filename)
	return tasklist, err
}

// WriteToFilename writes a TaskList to the specified file (most likely called "todo.txt").
func WriteToFilename(tasklist *TaskList, filename string) error {
	return tasklist.WriteToFilename(filename)
}
