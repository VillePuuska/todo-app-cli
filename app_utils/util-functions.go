package app_utils

import (
	"encoding/json"
	"log"
	"os"
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

func SaveList(TodoList *[]ListItem, filepath string) {
	/*if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(filepath)
		if err != nil {
			fmt.Println("what is even going on")
			log.Fatal(err)
		}
		f.Close()
	}*/
	marshaled, _ := json.Marshal(&TodoList)
	err := os.WriteFile(filepath, []byte(marshaled), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func Test(find_id int, TodoList *[]ListItem) int {
	return findIndex(find_id, TodoList)
}
