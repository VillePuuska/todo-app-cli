package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/VillePuuska/todo-app-cli/app_utils"
)

func main() {
	fmt.Println("main called")
	var TodoList []app_utils.ListItem

	for i := 0; i < 4; i++ {
		app_utils.AddItem("Test list item "+strconv.Itoa(i), &TodoList)
		if i != 3 {
			time.Sleep(time.Second)
		}
	}
	fmt.Println(printList(TodoList))

	// Marshal & Unmarshal testing
	marshaled, _ := json.Marshal(&TodoList)
	fmt.Println(string(marshaled))
	var unmarshaled []app_utils.ListItem
	json.Unmarshal([]byte(marshaled), &unmarshaled)
	fmt.Println(printList(unmarshaled))

	app_utils.SaveList(&TodoList, "projects/test_project.json")

	/*
		fmt.Println(app_utils.Test(0, &TodoList))
		fmt.Println(app_utils.Test(3, &TodoList))
		fmt.Println(app_utils.Test(4, &TodoList))

		app_utils.UpdateStatus(0, "asd", &TodoList)
		fmt.Println(printList(TodoList))
		app_utils.UpdateStatus(0, "done", &TodoList)
		fmt.Println(printList(TodoList))
		app_utils.UpdateStatus(1, "working on", &TodoList)
		fmt.Println(printList(TodoList))
		time.Sleep(time.Second)
		app_utils.UpdateStatus(1, "done", &TodoList)
		fmt.Println(printList(TodoList))
		time.Sleep(time.Second)
		app_utils.UpdateStatus(1, "backlog", &TodoList)
		fmt.Println(printList(TodoList))

		app_utils.DeleteItem(0, &TodoList)
		fmt.Println(printList(TodoList))
		app_utils.DeleteItem(6, &TodoList)
		fmt.Println(printList(TodoList))
		app_utils.DeleteItem(1, &TodoList)
		fmt.Println(printList(TodoList))
	*/
}

func printList(TodoList []app_utils.ListItem) string {
	res := ""
	for _, item := range TodoList {
		res += strconv.Itoa(item.Id) + " " + item.Description + " " + item.Status + " " + item.Added.String() + " " + item.Started.String() + " " + item.Finished.String() + "\n"
	}
	return res
}
