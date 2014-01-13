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
	"strings"
)

// TaskList represents a list of todo.txt task entries.
// It is usually loaded from a whole todo.txt file.
type TaskList []Task

// IgnoreComments can be set to 'false', in order to revert to a more standard todo.txt behaviour.
// The todo.txt format does not define comments.
var (
	// IgnoreComments is used to switch ignoring of comments (lines starting with "#").
	// If this is set to 'false', then lines starting with "#" will be parsed as tasks.
	IgnoreComments = true
)

// String returns a complete tasklist string in todo.txt format.
func (tasklist TaskList) String() (text string) {
	for _, task := range tasklist {
		text += fmt.Sprintf("%s\n", task.String())
	}
	return text
}

// AddTask appends a task to the current TaskList, and takes care to set the Task.Id correctly, modifying the Task by the given pointer!
func (tasklist *TaskList) AddTask(task *Task) (err error) {
	task.Id = 0
	for _, t := range *tasklist {
		if t.Id > task.Id {
			task.Id = t.Id
		}
	}
	task.Id += 1

	*tasklist = append(*tasklist, *task)
	return
}

// LoadFromFile loads a TaskList from *os.File.
//
// Using *os.File instead of a filename allows to also use os.Stdin.
//
// Note: This will clear the current TaskList and overwrite it's contents with whatever is in *os.File.
func (tasklist *TaskList) LoadFromFile(file *os.File) error {
	*tasklist = []Task{} // Empty tasklist

	taskId := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Trim(scanner.Text(), "\t\n\r ") // Read line

		// Ignore blank or comment lines
		if text == "" || (IgnoreComments && strings.HasPrefix(text, "#")) {
			continue
		}

		task, err := ParseTask(text)
		if err != nil {
			return err
		}
		task.Id = taskId

		*tasklist = append(*tasklist, *task)
		taskId++
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
	return ioutil.WriteFile(filename, []byte(tasklist.String()), 0640)
}

// LoadFromFile loads and returns a TaskList from *os.File.
//
// Using *os.File instead of a filename allows to also use os.Stdin.
func LoadFromFile(file *os.File) (TaskList, error) {
	tasklist := TaskList{}
	if err := tasklist.LoadFromFile(file); err != nil {
		return nil, err
	}
	return tasklist, nil
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
	if err := tasklist.LoadFromFilename(filename); err != nil {
		return nil, err
	}
	return tasklist, nil
}

// WriteToFilename writes a TaskList to the specified file (most likely called "todo.txt").
func WriteToFilename(tasklist *TaskList, filename string) error {
	return tasklist.WriteToFilename(filename)
}
