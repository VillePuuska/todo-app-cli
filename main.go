package main

import (
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

	fmt.Println(app_utils.Test(0, &TodoList))
	fmt.Println(app_utils.Test(3, &TodoList))
	fmt.Println(app_utils.Test(4, &TodoList))

	app_utils.DeleteItem(0, &TodoList)
	fmt.Println(printList(TodoList))
	app_utils.DeleteItem(6, &TodoList)
	fmt.Println(printList(TodoList))
	app_utils.DeleteItem(1, &TodoList)
	fmt.Println(printList(TodoList))
}

func printList(TodoList []app_utils.ListItem) string {
	res := ""
	for _, item := range TodoList {
		res += strconv.Itoa(item.Id) + " " + item.Description + " " + item.Status + " " + item.Added.String() + "\n"
	}
	return res
}
