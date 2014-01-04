/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package todotxt

import (
	"errors"
	"sort"
	"time"
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
		return errors.New("unrecognized sort option")
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
		if order == SORT_PRIORITY_ASC { // ASC
			if t1.HasPriority() && t2.HasPriority() {
				return t1.Priority < t2.Priority
			}
			return t1.HasPriority()
		}
		// DESC
		if t1.HasPriority() && t2.HasPriority() {
			return t1.Priority > t2.Priority
		}
		return !t1.HasPriority()
	})
	return tasklist
}

func sortByDate(asc bool, hasDate1, hasDate2 bool, date1, date2 time.Time) bool {
	if asc { // ASC
		if hasDate1 && hasDate2 {
			return date1.Before(date2)
		}
		return hasDate2
	}
	// DESC
	if hasDate1 && hasDate2 {
		return date1.After(date2)
	}
	return !hasDate2
}

func (tasklist *TaskList) sortByCreatedDate(order int) *TaskList {
	tasklist.sortBy(func(t1, t2 *Task) bool {
		return sortByDate(order == SORT_CREATED_DATE_ASC, t1.HasCreatedDate(), t2.HasCreatedDate(), t1.CreatedDate, t2.CreatedDate)
	})
	return tasklist
}

func (tasklist *TaskList) sortByCompletedDate(order int) *TaskList {
	tasklist.sortBy(func(t1, t2 *Task) bool {
		return sortByDate(order == SORT_COMPLETED_DATE_ASC, t1.HasCompletedDate(), t2.HasCompletedDate(), t1.CompletedDate, t2.CompletedDate)
	})
	return tasklist
}

func (tasklist *TaskList) sortByDueDate(order int) *TaskList {
	tasklist.sortBy(func(t1, t2 *Task) bool {
		return sortByDate(order == SORT_DUE_DATE_ASC, t1.HasDueDate(), t2.HasDueDate(), t1.DueDate, t2.DueDate)
	})
	return tasklist
}
