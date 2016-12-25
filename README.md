go-todotxt
==========

A Go todo.txt library.  

[![GoDoc](https://godoc.org/github.com/JamesClonk/go-todotxt?status.png)](https://godoc.org/github.com/JamesClonk/go-todotxt) [![Build Status](https://travis-ci.org/JamesClonk/go-todotxt.png?branch=master)](https://travis-ci.org/JamesClonk/go-todotxt)

The *todotxt* package is a Go client library for Gina Trapani's [todo.txt](https://github.com/ginatrapani/todo.txt-cli/) files.
It allows for parsing and manipulating of task lists and tasks in the todo.txt format.

## Installation

	$ go get github.com/JamesClonk/go-todotxt

## Requirements

go-todotxt requires Go1.1 or higher.

## Usage

```go
	package main

	import (
		"fmt"
		"github.com/JamesClonk/go-todotxt"
		"log"
	)

	func main() {
		todotxt.IgnoreComments = false

		tasklist, err := todotxt.LoadFromFilename("todo.txt")
		if err != nil {
			log.Fatal(err)
		}

		// tasklist now contains a slice of Tasks
		fmt.Printf("Task 2, todo: %v\n", tasklist[1].Todo)
		fmt.Printf("Task 3: %v\n", tasklist[2])
		fmt.Printf("Task 4, has priority: %v\n\n", tasklist[3].HasPriority())
		fmt.Print(tasklist)

		// Filter list to get only completed tasks
		completedList := tasklist.Filter(func(t Task) bool {
			return t.Completed 
		})
		fmt.Print(completedList)

		// Add a new empty Task to tasklist
		task := NewTask()
		tasklist.AddTask(&task)

		// Or a parsed Task from a string
		parsedTask, _ := ParseTask("x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12")
		tasklist.AddTask(parsed)

		// Update an existing task
		task, _ := tasklist.GetTask(2) // Task pointer
		task.Todo = "Do something different.."
		tasklist.WriteToFilename("todo.txt")
	}
```

## Documentation

See [GoDoc - Documentation](https://godoc.org/github.com/JamesClonk/go-todotxt) for further documentation.

## License

The source files are distributed under the [Mozilla Public License, version 2.0](http://mozilla.org/MPL/2.0/), unless otherwise noted.  
Please read the [FAQ](http://www.mozilla.org/MPL/2.0/FAQ.html) if you have further questions regarding the license. 
