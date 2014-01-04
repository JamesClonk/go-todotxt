go-todotxt
==========

A Go todo.txt library.  

[![GoDoc](https://godoc.org/github.com/JamesClonk/go-todotxt?status.png)](https://godoc.org/github.com/JamesClonk/go-todotxt) [![Build Status](https://travis-ci.org/JamesClonk/go-todotxt.png?branch=master)](https://travis-ci.org/JamesClonk/go-todotxt) [![Codebot](https://codebot.io/badge/github.com/JamesClonk/go-todotxt.png)](http://codebot.io/doc/pkg/github.com/JamesClonk/go-todotxt "Codebot") [![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/JamesClonk/go-todotxt/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

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

		fmt.Printf("Task 2, todo: %v\n", tasklist[1].Todo)
		fmt.Printf("Task 3: %v\n", tasklist[2])
		fmt.Printf("Task 4, has priority: %v\n\n", tasklist[3].HasPriority())
		fmt.Print(tasklist)
	}
```

## Documentation

See [GoDoc - Documentation](https://godoc.org/github.com/JamesClonk/go-todotxt) for further documentation.

## License

The source files are distributed under the [Mozilla Public License, version 2.0](http://mozilla.org/MPL/2.0/), unless otherwise noted.  
Please read the [FAQ](http://www.mozilla.org/MPL/2.0/FAQ.html) if you have further questions regarding the license. 
