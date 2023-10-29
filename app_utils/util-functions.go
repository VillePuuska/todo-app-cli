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
