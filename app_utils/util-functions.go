package app_utils

import "time"

type ListItem struct {
	Id          int       // number for ordering and referring to tasks when editing and deleting
	Description string    // description of the task,
	Status      string    // backlog, working on, done
	Added       time.Time // added timestamp,
	Started     time.Time // started work timestamp,
	Finished    time.Time // finished timestamp.
}

func findIndex(find_id int, TodoList *[]ListItem) int {
	for i := 0; i < len(*TodoList); i++ {
		if (*TodoList)[i].Id == find_id {
			return i
		}
	}
	return -1
}

func AddItem(Description string, TodoList *[]ListItem) {

}

func Test(i ...int) ListItem {
	id := 1
	if len(i) > 0 {
		id = i[0]
	}
	return ListItem{
		Id:          id,
		Description: "Test list item",
		Status:      "backlog",
		Added:       time.Now(),
		Started:     time.Time{},
		Finished:    time.Time{},
	}
}

func Test2(find_id int, TodoList *[]ListItem) int {
	return findIndex(find_id, TodoList)
}
