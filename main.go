package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/VillePuuska/todo-app-cli/app_utils"
)

const projectpath = "todo-app-cli-projects"

func main() {
	projectname := "test_project2"
	fileextension := ".json"

	_, err := os.Stat(projectpath)
	if os.IsNotExist(err) {
		fmt.Println("Missing folder for projects. Creating it.")
		err := os.Mkdir(projectpath, 0750)
		if err != nil {
			log.Fatal(err)
		}
		absolute_path, err := filepath.Abs(projectpath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Path to folder:", absolute_path)
	} else if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(filepath.Join(projectpath, projectname+fileextension))
	if os.IsNotExist(err) {
		fmt.Println("Project does not exist. Add it first.")
		return
	} else if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Adding items to list.")
	var TodoList []app_utils.ListItem

	for i := 0; i < 4; i++ {
		app_utils.AddItem("Test list item "+strconv.Itoa(i), &TodoList)
		if i != 3 {
			time.Sleep(time.Second)
		}
	}
	fmt.Println(listToString(&TodoList))

	fmt.Println("Marshaling and Unmarshaling the list.")
	marshaled, _ := json.Marshal(&TodoList)
	fmt.Println(string(marshaled))
	var unmarshaled []app_utils.ListItem
	json.Unmarshal([]byte(marshaled), &unmarshaled)
	fmt.Println(listToString(&unmarshaled))

	fmt.Println("Testing updating status.")
	app_utils.UpdateStatus(0, "asd", &TodoList)
	fmt.Println(listToString(&TodoList))
	app_utils.UpdateStatus(0, "done", &TodoList)
	fmt.Println(listToString(&TodoList))
	app_utils.UpdateStatus(1, "working on", &TodoList)
	fmt.Println(listToString(&TodoList))
	app_utils.UpdateStatus(2, "done", &TodoList)
	fmt.Println(listToString(&TodoList))

	fmt.Println("Testing deleting an item.")
	app_utils.DeleteItem(0, &TodoList)
	fmt.Println(listToString(&TodoList))

	fmt.Println("Saving and loading the list.")
	app_utils.SaveList(&TodoList, filepath.Join(projectpath, projectname+fileextension))
	readList := app_utils.ReadList(filepath.Join(projectpath, projectname+fileextension))
	fmt.Println(listToString(readList))

	fmt.Println("Testing changing id.")
	app_utils.ChangeId(0, 4, readList)
	fmt.Println(listToString(readList))
	app_utils.ChangeId(0, 2, readList)
	fmt.Println(listToString(readList))
	app_utils.ChangeId(2, 1, readList)
	fmt.Println(listToString(readList))

	fmt.Println("Testing sorting.")
	app_utils.OrderList("asd", readList)
	fmt.Println(listToString(readList))
	app_utils.OrderList("id", readList)
	fmt.Println(listToString(readList))
	app_utils.OrderList("description", readList)
	fmt.Println(listToString(readList))
	app_utils.OrderList("status", readList)
	fmt.Println(listToString(readList))
	app_utils.OrderList("added", readList)
	fmt.Println(listToString(readList))
	app_utils.OrderList("finished", readList)
	fmt.Println(listToString(readList))
}

func listToString(TodoList *[]app_utils.ListItem) string {
	res := ""
	for _, item := range *TodoList {
		res += strconv.Itoa(item.Id) + " " + item.Description + " " + item.Status + " " + item.Added.String() + " " + item.Started.String() + " " + item.Finished.String() + "\n"
	}
	return res
}
