package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/VillePuuska/todo-app-cli/app_utils"
)

const projectpath = "todo-app-cli-projects"
const fileextension = ".json"

func main() {
	var projectname string

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

	files := listProjects()
	projects := make([]string, 0)
	for _, project := range *files {
		if strings.HasSuffix(project, fileextension) {
			projects = append(projects, project)
			fmt.Println(len(projects)-1, project)
		}
	}

	var user_input string
	intCheck := regexp.MustCompile("^[0-9]+$")
	reader := bufio.NewReader(os.Stdin)
loop:
	for {
		fmt.Println("Choose a project with the corresponding integer: (\"add\" to create new blank project, \"stop\" to exit program)")
		user_input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		user_input = strings.TrimSuffix(strings.ToLower(user_input), "\n")
		if user_input == "stop" {
			return
		} else if user_input == "add" {
			fmt.Println("Sorry. This function is not yet implemented.")
		} else if intCheck.MatchString(user_input) {
			project_index, err := strconv.Atoi(user_input)
			if err != nil {
				log.Fatal(err)
			}
			if project_index >= 0 && project_index < len(projects) {
				projectname = projects[project_index]
				break loop
			}
		}
	}

	TodoList := app_utils.ReadList(filepath.Join(projectpath, projectname))
	var EmptyList []app_utils.ListItem
	for {
		fmt.Println("Commands: \"show\" will print the list, \"test\" will run tests, \"hi\" greet you back, \"stop\" will exit program, anything else will be echoed")
		user_input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		user_input = strings.TrimSuffix(strings.ToLower(user_input), "\n")
		switch user_input {
		case "stop":
			return
		case "show":
			fmt.Println(listToString(TodoList))
		case "hi":
			fmt.Print("Well hello!" + "\n\n")
		case "test":
			test(&EmptyList, projectname)
			EmptyList = make([]app_utils.ListItem, 0)
		default:
			fmt.Println(user_input)
		}
	}
}

func listProjects() *[]string {
	f, err := os.Open(projectpath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	res, err := f.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}
	return &res
}

func test(TodoList *[]app_utils.ListItem, projectname string) {
	fmt.Println("Loading chosen project.")
	fmt.Println(listToString(app_utils.ReadList(filepath.Join(projectpath, projectname))))

	fmt.Println("Adding items to list.")
	for i := 0; i < 4; i++ {
		app_utils.AddItem("Test list item "+strconv.Itoa(i), TodoList)
		if i != 3 {
			time.Sleep(time.Second)
		}
	}
	fmt.Println(listToString(TodoList))

	fmt.Println("Marshaling and Unmarshaling the list.")
	marshaled, _ := json.Marshal(TodoList)
	fmt.Println(string(marshaled))
	var unmarshaled []app_utils.ListItem
	json.Unmarshal([]byte(marshaled), &unmarshaled)
	fmt.Println(listToString(&unmarshaled))

	fmt.Println("Testing updating status.")
	app_utils.UpdateStatus(0, "asd", TodoList)
	fmt.Println(listToString(TodoList))
	app_utils.UpdateStatus(0, "done", TodoList)
	fmt.Println(listToString(TodoList))
	app_utils.UpdateStatus(1, "working on", TodoList)
	fmt.Println(listToString(TodoList))
	app_utils.UpdateStatus(2, "done", TodoList)
	fmt.Println(listToString(TodoList))

	fmt.Println("Testing deleting an item.")
	app_utils.DeleteItem(0, TodoList)
	fmt.Println(listToString(TodoList))

	fmt.Println("Saving and loading the list.")
	app_utils.SaveList(TodoList, filepath.Join(projectpath, projectname))
	readList := app_utils.ReadList(filepath.Join(projectpath, projectname))
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

	app_utils.SaveList(readList, filepath.Join(projectpath, "sorting_"+projectname))
}

func listToString(TodoList *[]app_utils.ListItem) string {
	res := ""
	for _, item := range *TodoList {
		res += strconv.Itoa(item.Id) + " " + item.Description + " " + item.Status + " " + item.Added.String() + " " + item.Started.String() + " " + item.Finished.String() + "\n"
	}
	return res
}
