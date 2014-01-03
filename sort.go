/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package todotxt

import (
	"errors"
	"sort"
)

// Flags for defining sort element and order.
const (
	SORT_PRIORITY_ASC = iota
	SORT_PRIORITY_DESC
	SORT_CREATED_DATE_ASC
	SORT_CREATED_DATE_DESC
	SORT_COMPLETED_DATE_ASC
	SORT_COMPLETED_DATE_DESC
	SORT_DUE_DATE_ASC
	SORT_DUE_DATE_DESC
)

// Sort allows a TaskList to be sorted by certain predefined fields.
// See constants SORT_* for fields and sort order.
func (tasklist *TaskList) Sort(sortFlag int) error {
	switch sortFlag {
	case SORT_PRIORITY_ASC, SORT_PRIORITY_DESC:
		tasklist.sortByPriority(sortFlag)
	case SORT_CREATED_DATE_ASC, SORT_CREATED_DATE_DESC:
		tasklist.sortByCreatedDate(sortFlag)
	case SORT_COMPLETED_DATE_ASC, SORT_COMPLETED_DATE_DESC:
		tasklist.sortByCompletedDate(sortFlag)
	case SORT_DUE_DATE_ASC, SORT_DUE_DATE_DESC:
		tasklist.sortByDueDate(sortFlag)
	default:
		return errors.New("Unrecognized sort option")
	}

	return nil
}

type tasklistSort struct {
	tasklists TaskList
	by        func(t1, t2 *Task) bool
}

func (ts *tasklistSort) Len() int {
	return len(ts.tasklists)
}

func (ts *tasklistSort) Swap(l, r int) {
	ts.tasklists[l], ts.tasklists[r] = ts.tasklists[r], ts.tasklists[l]
}

func (ts *tasklistSort) Less(l, r int) bool {
	return ts.by(&ts.tasklists[l], &ts.tasklists[r])
}

func (tasklist *TaskList) sortBy(by func(t1, t2 *Task) bool) *TaskList {
	ts := &tasklistSort{
		tasklists: *tasklist,
		by:        by,
	}
	sort.Sort(ts)
	return tasklist
}

func (tasklist *TaskList) sortByPriority(order int) *TaskList {
	tasklist.sortBy(func(t1, t2 *Task) bool {
		if order == SORT_PRIORITY_DESC { // DESC
			if t1.HasPriority() && t2.HasPriority() {
				return t1.Priority > t2.Priority
			} else {
				return !t1.HasPriority()
			}
		} else { // ASC
			if t1.HasPriority() && t2.HasPriority() {
				return t1.Priority < t2.Priority
			} else {
				return t1.HasPriority()
			}
		}
	})
	return tasklist
}

func (tasklist *TaskList) sortByCreatedDate(order int) *TaskList {
	tasklist.sortBy(func(t1, t2 *Task) bool {
		if order == SORT_CREATED_DATE_DESC { // DESC
			if t1.HasCreatedDate() && t2.HasCreatedDate() {
				return t1.CreatedDate.After(t2.CreatedDate)
			} else {
				return !t1.HasCreatedDate()
			}
		} else { // ASC
			if t1.HasCreatedDate() && t2.HasCreatedDate() {
				return t1.CreatedDate.Before(t2.CreatedDate)
			} else {
				return t1.HasCreatedDate()
			}
		}
	})
	return tasklist
}

func (tasklist *TaskList) sortByCompletedDate(order int) *TaskList {
	tasklist.sortBy(func(t1, t2 *Task) bool {
		if order == SORT_COMPLETED_DATE_DESC { // DESC
			if t1.HasCompletedDate() && t2.HasCompletedDate() {
				return t1.CompletedDate.After(t2.CompletedDate)
			} else {
				return !t1.HasCompletedDate()
			}
		} else { // ASC
			if t1.HasCompletedDate() && t2.HasCompletedDate() {
				return t1.CompletedDate.Before(t2.CompletedDate)
			} else {
				return t1.HasCompletedDate()
			}
		}
	})
	return tasklist
}

func (tasklist *TaskList) sortByDueDate(order int) *TaskList {
	tasklist.sortBy(func(t1, t2 *Task) bool {
		if order == SORT_DUE_DATE_DESC { // DESC
			if t1.HasDueDate() && t2.HasDueDate() {
				return t1.DueDate.After(t2.DueDate)
			} else {
				return !t1.HasDueDate()
			}
		} else { // ASC
			if t1.HasDueDate() && t2.HasDueDate() {
				return t1.DueDate.Before(t2.DueDate)
			} else {
				return t1.HasDueDate()
			}
		}
	})
	return tasklist
}
