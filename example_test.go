/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package todotxt

import (
	"fmt"
	"log"
)

func ExampleLoadFromFilename() {
	if tasklist, err := LoadFromFilename("todo.txt"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(tasklist) // String representation of TaskList works as expected.
	}

	// Output:
	// (A) Call Mom @Phone +Family
	// (A) Schedule annual checkup +Health
	// (B) Outline chapter 5 @Computer +Novel
	// (C) Add cover sheets @Office +TPSReports
	// Plan backyard herb garden @Home
	// Pick up milk @GroceryStore
	// Research self-publishing services @Computer +Novel
	// x Download Todo.txt mobile app @Phone
}

func ExampleTaskList_LoadFromFilename() {
	var tasklist TaskList

	// This will overwrite whatever was in the tasklist before.
	// Irrelevant here since the list is still empty.
	if err := tasklist.LoadFromFilename("todo.txt"); err != nil {
		log.Fatal(err)
	}

	fmt.Println(tasklist[0].Todo)      // Text part of first task (Call Mom)
	fmt.Println(tasklist[2].Contexts)  // Slice of contexts from third task ([Computer])
	fmt.Println(tasklist[3].Priority)  // Priority of fourth task (C)
	fmt.Println(tasklist[7].Completed) // Completed flag of eigth task (true)
	// Output:
	// Call Mom
	// [Computer]
	// C
	// true
}
