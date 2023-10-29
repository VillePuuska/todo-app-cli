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
		TodoList = append(TodoList, app_utils.Test())
		if i != 3 {
			time.Sleep(time.Second)
		}
	}
	fmt.Println(printList(TodoList))

	TodoList = []app_utils.ListItem{}
	for i := 0; i < 4; i++ {
		TodoList = append(TodoList, app_utils.Test(i+1))
		if i != 3 {
			time.Sleep(time.Second)
		}
	}
	fmt.Println(printList(TodoList))
}

func printList(TodoList []app_utils.ListItem) string {
	res := ""
	for _, item := range TodoList {
		res += strconv.Itoa(item.Id) + " " + item.Description + " " + item.Status + " " + item.Added.String() + "\n"
	}
	return res
}
