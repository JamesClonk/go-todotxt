/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package todotxt

import (
	"testing"
)

var (
	testInputSort = "testdata/sort_todo.txt"
)

func TestTaskSortByPriority(t *testing.T) {
	testTasklist.LoadFromFilename(testInputSort)
	taskId := 0

	testTasklist = testTasklist[taskId : taskId+6]

	if err := testTasklist.Sort(SORT_PRIORITY_ASC); err != nil {
		t.Fatal(err)
	}

	testExpected = "(A) 2012-01-30 Call Mom @Call @Phone +Family"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "(D) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "2013-02-22 Pick up milk @GroceryStore"
	testGot = testTasklist[4].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[5] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05"
	testGot = testTasklist[5].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[6] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	if err := testTasklist.Sort(SORT_PRIORITY_DESC); err != nil {
		t.Fatal(err)
	}

	testExpected = "x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "2013-02-22 Pick up milk @GroceryStore"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "(D) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[4].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[5] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "(A) 2012-01-30 Call Mom @Call @Phone +Family"
	testGot = testTasklist[5].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[6] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}
}

func TestTaskSortByCreatedDate(t *testing.T) {
	testTasklist.LoadFromFilename(testInputSort)
	taskId := 6

	testTasklist = testTasklist[taskId : taskId+5]

	if err := testTasklist.Sort(SORT_CREATED_DATE_ASC); err != nil {
		t.Fatal(err)
	}

	testExpected = "x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "(A) Call Mom @Call @Phone +Family"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "2013-02-22 Pick up milk @GroceryStore"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[4].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[5] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	if err := testTasklist.Sort(SORT_CREATED_DATE_DESC); err != nil {
		t.Fatal(err)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "2013-02-22 Pick up milk @GroceryStore"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "(A) Call Mom @Call @Phone +Family"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05"
	testGot = testTasklist[4].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[5] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}
}

func TestTaskSortByCompletedDate(t *testing.T) {
	testTasklist.LoadFromFilename(testInputSort)
	taskId := 11

	testTasklist = testTasklist[taskId : taskId+6]

	if err := testTasklist.Sort(SORT_COMPLETED_DATE_ASC); err != nil {
		t.Fatal(err)
	}

	testExpected = "x Download Todo.txt mobile app @Phone"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "2013-02-22 Pick up milk @GroceryStore"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05"
	testGot = testTasklist[4].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[5] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-04 2014-01-01 Create some more golang library test cases @Go +go-todotxt"
	testGot = testTasklist[5].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[6] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	if err := testTasklist.Sort(SORT_COMPLETED_DATE_DESC); err != nil {
		t.Fatal(err)
	}

	testExpected = "x 2014-01-04 2014-01-01 Create some more golang library test cases @Go +go-todotxt"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "2013-02-22 Pick up milk @GroceryStore"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[4].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[5] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x Download Todo.txt mobile app @Phone"
	testGot = testTasklist[5].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[6] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}
}

func TestTaskSortByDueDate(t *testing.T) {
	testTasklist.LoadFromFilename(testInputSort)
	taskId := 17

	testTasklist = testTasklist[taskId : taskId+4]

	if err := testTasklist.Sort(SORT_DUE_DATE_ASC); err != nil {
		t.Fatal(err)
	}

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	if err := testTasklist.Sort(SORT_DUE_DATE_DESC); err != nil {
		t.Fatal(err)
	}

	testExpected = "(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}
}

func TestTaskSortError(t *testing.T) {
	testTasklist.LoadFromFilename(testInputSort)

	if err := testTasklist.Sort(123); err == nil {
		t.Errorf("Expected Sort() to fail because of unrecognized sort option, but it didn't!")
	} else if err.Error() != "unrecognized sort option" {
		t.Error(err)
	}
}
