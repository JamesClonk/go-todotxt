package main

import (
	"fmt"
	"github.com/JamesClonk/go-todotxt"
	"log"
)

func main() {
	todotxt.IgnoreComments = false

	tasklist, err := todotxt.LoadFromFilename("../todo.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Task 2, todo: %v\n", tasklist[1].Todo)
	fmt.Printf("Task 3: %v\n", tasklist[2])
	fmt.Printf("Task 4, has priority: %v\n\n", tasklist[3].HasPriority())
	fmt.Print(tasklist)
}
