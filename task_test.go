package todotxt

import (
	"testing"
	"time"
)

var (
	testInputTask = "testdata/task_todo.txt"
)

func TestTaskString(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskId := 1

	testExpected = "2013-02-22 Pick up milk @GroceryStore"
	testGot = testTasklist[taskId-1].String()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = "x Download Todo.txt mobile app @Phone"
	testGot = testTasklist[taskId-1].String()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = "(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
	testGot = testTasklist[taskId-1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[taskId-1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = "x 2014-01-03 2014-01-01 Create some more golang library test cases @Go +go-todotxt"
	testGot = testTasklist[taskId-1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
	taskId++
}

func TestTaskPriority(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskId := 6

	testExpected = "B"
	testGot = testTasklist[taskId-1].Priority
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have priority '%s', but got '%s'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = "C"
	testGot = testTasklist[taskId-1].Priority
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have priority '%s', but got '%s'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = "B"
	testGot = testTasklist[taskId-1].Priority
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have priority '%s', but got '%s'", taskId, testExpected, testGot)
	}
	taskId++

	if testTasklist[taskId-1].HasPriority() {
		t.Errorf("Expected Task[%d] to have no priority, but got '%s'", taskId, testTasklist[4].Priority)
	}
	taskId++
}

func TestTaskCreatedDate(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskId := 10

	testExpected, err := time.Parse(DateLayout, "2012-01-30")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CreatedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected, err = time.Parse(DateLayout, "2013-02-22")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CreatedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected, err = time.Parse(DateLayout, "2014-01-01")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CreatedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected, err = time.Parse(DateLayout, "2013-12-30")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CreatedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected, err = time.Parse(DateLayout, "2014-01-01")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CreatedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	if testTasklist[taskId-1].HasCreatedDate() {
		t.Errorf("Expected Task[%d] to have no created date, but got '%v'", taskId, testTasklist[4].CreatedDate)
	}
	taskId++
}

func TestTaskContexts(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskId := 16

	testExpected = []string{"Call", "Phone"}
	testGot = testTasklist[taskId-1].Contexts
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have contexts '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = []string{"Office"}
	testGot = testTasklist[taskId-1].Contexts
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have contexts '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = []string{"Electricity", "Home", "Of_Super-Importance", "Television"}
	testGot = testTasklist[taskId-1].Contexts
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have contexts '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = []string{}
	testGot = testTasklist[taskId-1].Contexts
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have no contexts, but got '%v'", taskId, testGot)
	}
	taskId++
}

func TestTasksProjects(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskId := 20

	testExpected = []string{"Gardening", "Improving", "Planning", "Relaxing-Work"}
	testGot = testTasklist[taskId-1].Projects
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have projects '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = []string{"Novel"}
	testGot = testTasklist[taskId-1].Projects
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have projects '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = []string{}
	testGot = testTasklist[taskId-1].Projects
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have no projects, but got '%v'", taskId, testGot)
	}
	taskId++
}

func TestTaskDueDate(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskId := 23

	testExpected, err := time.Parse(DateLayout, "2014-02-17")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].DueDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have due date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	if testTasklist[taskId-1].HasDueDate() {
		t.Errorf("Expected Task[%d] to have no due date, but got '%v'", taskId, testTasklist[taskId-1].DueDate)
	}
	taskId++
}

func TestTaskAddonTags(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskId := 25

	testExpected = map[string]string{"Level": "5", "private": "false"}
	testGot = testTasklist[taskId-1].AdditionalTags
	if len(testGot.(map[string]string)) != 2 ||
		!compareMaps(testGot.(map[string]string), testExpected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have addon tags '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = map[string]string{"Importance": "Very!"}
	testGot = testTasklist[taskId-1].AdditionalTags
	if len(testGot.(map[string]string)) != 1 ||
		!compareMaps(testGot.(map[string]string), testExpected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have projects '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = map[string]string{}
	testGot = testTasklist[taskId-1].AdditionalTags
	if len(testGot.(map[string]string)) != 0 ||
		!compareMaps(testGot.(map[string]string), testExpected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have no additional tags, but got '%v'", taskId, testGot)
	}
	taskId++

	testExpected = map[string]string{}
	testGot = testTasklist[taskId-1].AdditionalTags
	if len(testGot.(map[string]string)) != 0 ||
		!compareMaps(testGot.(map[string]string), testExpected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have no additional tags, but got '%v'", taskId, testGot)
	}
	taskId++
}

func TestTaskCompleted(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskId := 29

	testExpected = true
	testGot = testTasklist[taskId-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] not to be completed, but got '%v'", taskId, testGot)
	}
	taskId++

	testExpected = true
	testGot = testTasklist[taskId-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] not to be completed, but got '%v'", taskId, testGot)
	}
	taskId++

	testExpected = true
	testGot = testTasklist[taskId-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] not to be completed, but got '%v'", taskId, testGot)
	}
	taskId++

	testExpected = false
	testGot = testTasklist[taskId-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] not to be completed, but got '%v'", taskId, testGot)
	}
	taskId++

	testExpected = false
	testGot = testTasklist[taskId-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] not to be completed, but got '%v'", taskId, testGot)
	}
	taskId++
}

func TestTaskCompletedDate(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskId := 34

	if testTasklist[taskId-1].HasCompletedDate() {
		t.Errorf("Expected Task[%d] to not have a completed date, but got '%v'", taskId, testTasklist[taskId-1].CompletedDate)
	}
	taskId++

	testExpected, err := time.Parse(DateLayout, "2014-01-03")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CompletedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have completed date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	if testTasklist[taskId-1].HasCompletedDate() {
		t.Errorf("Expected Task[%d] to not have a completed date, but got '%v'", taskId, testTasklist[taskId-1].CompletedDate)
	}
	taskId++

	testExpected, err = time.Parse(DateLayout, "2014-01-02")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CompletedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have completed date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected, err = time.Parse(DateLayout, "2014-01-03")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CompletedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have completed date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	if testTasklist[taskId-1].HasCompletedDate() {
		t.Errorf("Expected Task[%d] to not have a completed date, but got '%v'", taskId, testTasklist[taskId-1].CompletedDate)
	}
	taskId++
}

func TestTaskIsOverdue(t *testing.T) {
	testTasklist.LoadFromFilename(testInputTask)
	taskId := 40

	testGot = testTasklist[taskId-1].IsOverdue()
	if !testGot.(bool) {
		t.Errorf("Expected Task[%d] to be overdue, but got '%v'", taskId, testGot)
	}
	taskId++

	testGot = testTasklist[taskId-1].IsOverdue()
	if testGot.(bool) {
		t.Errorf("Expected Task[%d] not to be overdue, but got '%v'", taskId, testGot)
	}
	testTasklist[taskId-1].DueDate = time.Now().AddDate(0, 0, 1)
	testGot = testTasklist[taskId-1].Due()
	if testGot.(time.Duration).Hours() < 23 ||
		testGot.(time.Duration).Hours() > 25 {
		t.Errorf("Expected Task[%d] to be due in 24 hours, but got '%v'", taskId, testGot)
	}
	taskId++

	testGot = testTasklist[taskId-1].IsOverdue()
	if !testGot.(bool) {
		t.Errorf("Expected Task[%d] to be overdue, but got '%v'", taskId, testGot)
	}
	testTasklist[taskId-1].DueDate = time.Now().AddDate(0, 0, -3)
	testGot = testTasklist[taskId-1].Due()
	if testGot.(time.Duration).Hours() < 71 ||
		testGot.(time.Duration).Hours() > 73 {
		t.Errorf("Expected Task[%d] to be due since 72 hours, but got '%v'", taskId, testGot)
	}
	taskId++
}

func compareSlices(list1 []string, list2 []string) bool {
	if len(list1) != len(list2) {
		return false
	}

	for i := range list1 {
		if list1[i] != list2[i] {
			return false
		}
	}

	return true
}

func compareMaps(map1 map[string]string, map2 map[string]string) bool {
	if len(map1) != len(map2) {
		return false
	}

	compare := func(map1 map[string]string, map2 map[string]string) bool {
		for key, value := range map1 {
			if value2, found := map2[key]; !found {
				return false
			} else if value != value2 {
				return false
			}
		}
		return true
	}

	return compare(map1, map2) && compare(map2, map1)
}
