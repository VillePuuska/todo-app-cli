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
	new_id := len(*TodoList)
	new_item := ListItem{
		Id:          new_id,
		Description: Description,
		Status:      "backlog",
		Added:       time.Now(),
		Started:     time.Time{},
		Finished:    time.Time{},
	}
	*TodoList = append(*TodoList, new_item)
}

func DeleteItem(id int, TodoList *[]ListItem) {
	deletion_index := findIndex(id, TodoList)
	if deletion_index == -1 {
		return
	}
	*TodoList = append((*TodoList)[:deletion_index], (*TodoList)[deletion_index+1:]...)
	for i := 0; i < len(*TodoList); i++ {
		if (*TodoList)[i].Id > id {
			(*TodoList)[i].Id--
		}
	}
}

func Test(find_id int, TodoList *[]ListItem) int {
	return findIndex(find_id, TodoList)
}
