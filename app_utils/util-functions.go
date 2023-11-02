package app_utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

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

func UpdateStatus(id int, new_status string, TodoList *[]ListItem) {
	index := findIndex(id, TodoList)
	if index == -1 {
		return
	}
	if new_status != "backlog" && new_status != "working on" && new_status != "done" {
		return
	}
	(*TodoList)[index].Status = new_status
	if new_status == "backlog" {
		(*TodoList)[index].Added = time.Now()
		(*TodoList)[index].Started = time.Time{}
		(*TodoList)[index].Finished = time.Time{}
	} else if new_status == "working on" {
		(*TodoList)[index].Started = time.Now()
		(*TodoList)[index].Finished = time.Time{}
	} else if new_status == "done" {
		(*TodoList)[index].Finished = time.Now()
		if (*TodoList)[index].Started == (time.Time{}) {
			(*TodoList)[index].Started = (*TodoList)[index].Finished
		}
	}
}

func ChangeId(old_id, new_id int, TodoList *[]ListItem) {
	if old_id < 0 || old_id >= len(*TodoList) || new_id < 0 || new_id >= len(*TodoList) {
		return
	}
	for i := 0; i < len(*TodoList); i++ {
		if (*TodoList)[i].Id == old_id {
			(*TodoList)[i].Id = new_id
		} else if (*TodoList)[i].Id > old_id && (*TodoList)[i].Id <= new_id {
			(*TodoList)[i].Id--
		} else if (*TodoList)[i].Id < old_id && (*TodoList)[i].Id >= new_id {
			(*TodoList)[i].Id++
		}
	}
}

func OrderList(attribute string, TodoList *[]ListItem) {
	var cmp func(a, b ListItem) int
	switch strings.ToUpper(attribute) {
	case "ID":
		cmp = func(a, b ListItem) int {
			if a.Id < b.Id {
				return -1
			} else if a.Id > b.Id {
				return 1
			}
			return 0
		}
	case "DESCRIPTION":
		cmp = func(a, b ListItem) int {
			if a.Description < b.Description {
				return -1
			} else if a.Description > b.Description {
				return 1
			}
			return 0
		}
	case "STATUS":
		cmp = func(a, b ListItem) int {
			if a.Status == b.Status {
				return 0
			} else if a.Status == "backlog" || (a.Status == "working on" && b.Status == "done") {
				return -1
			}
			return 1
		}
	case "ADDED":
		cmp = func(a, b ListItem) int {
			if b.Added.After(a.Added) {
				return -1
			} else if a.Added.After(b.Added) {
				return 1
			}
			return 0
		}
	case "STARTED":
		cmp = func(a, b ListItem) int {
			if b.Started.After(a.Started) {
				return -1
			} else if a.Started.After(b.Started) {
				return 1
			}
			return 0
		}
	case "FINISHED":
		cmp = func(a, b ListItem) int {
			if b.Finished.After(a.Finished) {
				return -1
			} else if a.Finished.After(b.Finished) {
				return 1
			}
			return 0
		}
	default:
		fmt.Println("Incorrect attribute name.")
		return
	}
	slices.SortStableFunc(*TodoList, cmp)
}

func SaveList(TodoList *[]ListItem, filepath string) {
	marshaled, _ := json.Marshal(&TodoList)
	err := os.WriteFile(filepath, []byte(marshaled), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadList(filepath string) *[]ListItem {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	var unmarshaled []ListItem
	json.Unmarshal(data, &unmarshaled)
	return &unmarshaled
}
