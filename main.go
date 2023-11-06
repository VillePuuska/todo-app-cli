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
const commands = `Possible commands:
"add item" lets you add an item to the current project,
"change id" lets your change the id of an item in the current list, use this to manually order items,
"change project" lets you change the current project/todo-list or create a new blank one,
"delete item" lets you delete an item from the current project,
"delete project" lets you delete any of the projects in the project-folder,
"help" will print this message,
"reload" loads the current project from its saved file, use this if you want to undo your unsaved changes,
"save" saves your current project,
"show" prints the todo-list,
"sort" lets you order the list by any attribute,
"stop" or "quit" quits the app,
"test" runs the hardcoded tests,
"update item" lets you update the status of an item`

func main() {
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

	projectname := chooseProject()
	TodoList := app_utils.ReadList(filepath.Join(projectpath, projectname))

	var user_input string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Command: (\"help\" will print the possible commands)")
		user_input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		user_input = strings.TrimSuffix(strings.ToLower(user_input), "\n")
		switch user_input {
		case "add item":
			fmt.Println("Sorry. This function is not yet implemented.")
		case "change id":
			fmt.Println("Sorry. This function is not yet implemented.")
		case "change project":
			projectname = chooseProject()
			TodoList = app_utils.ReadList(filepath.Join(projectpath, projectname))
		case "delete item":
			fmt.Println("Sorry. This function is not yet implemented.")
		case "delete project":
			fmt.Println("Sorry. This function is not yet implemented.")
		case "help":
			fmt.Println(commands)
		case "reload":
			fmt.Println("Sorry. This function is not yet implemented.")
		case "save":
			fmt.Println("Sorry. This function is not yet implemented.")
		case "show":
			fmt.Println(listToString(TodoList))
		case "sort":
			fmt.Println("Sorry. This function is not yet implemented.")
		case "stop":
			return
		case "quit":
			return
		case "test":
			test(projectname)
		case "update item":
			fmt.Println("Sorry. This function is not yet implemented.")
		default:
			fmt.Println("Unrecognized command:")
			fmt.Println(user_input)
		}
	}
}

func chooseProject() string {
	projects := listProjects()
	for i, project := range projects {
		fmt.Println(i, project)
	}

	intCheck := regexp.MustCompile("^[0-9]+$")
	reader := bufio.NewReader(os.Stdin)
	var user_input string
	var err error
	for {
		fmt.Println("Choose a project with the corresponding integer: (\"add\" to create new blank project, \"stop\" to exit program)")
		user_input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		user_input = strings.TrimSuffix(strings.ToLower(user_input), "\n")
		if user_input == "stop" {
			os.Exit(0)
		} else if user_input == "add" {
			projectName := addProject()
			if projectName != "" {
				return projectName + fileextension
			}
		} else if intCheck.MatchString(user_input) {
			project_index, err := strconv.Atoi(user_input)
			if err != nil {
				log.Fatal(err)
			}
			if project_index >= 0 && project_index < len(projects) {
				return projects[project_index]
			}
		}
	}
}

func listProjects() []string {
	f, err := os.Open(projectpath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	res, err := f.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}
	projects := make([]string, 0)
	for _, project := range res {
		if strings.HasSuffix(project, fileextension) {
			projects = append(projects, project)
		}
	}
	return projects
}

func addProject() string {
	nameCheck := regexp.MustCompile("^[A-Za-z_]+$")
	reader := bufio.NewReader(os.Stdin)
	var user_input string
	var err error
	for {
		fmt.Println(`Choose a name for the new project (only characters A-Z, a-z and _ are allowed, "stop" returns without adding a new project):`)
		user_input, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		user_input = strings.TrimSuffix(strings.ToLower(user_input), "\n")
		if user_input == "stop" {
			return ""
		} else if nameCheck.MatchString(user_input) {
			_, err := os.Stat(filepath.Join(projectpath, user_input+fileextension))
			if err == nil {
				fmt.Println("Project already exists.")
			} else if os.IsNotExist(err) {
				emptyList := make([]app_utils.ListItem, 0)
				app_utils.SaveList(&emptyList, filepath.Join(projectpath, user_input+fileextension))
				return user_input
			} else {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Invalid name.")
		}
	}
}

func test(projectname string) {
	fmt.Println("Loading chosen project.")
	fmt.Println(listToString(app_utils.ReadList(filepath.Join(projectpath, projectname))))

	TodoList := make([]app_utils.ListItem, 0)
	fmt.Println("Adding items to list.")
	for i := 0; i < 4; i++ {
		app_utils.AddItem("Test list item "+strconv.Itoa(i), &TodoList)
		if i != 3 {
			time.Sleep(time.Second)
		}
	}
	fmt.Println(listToString(&TodoList))

	fmt.Println("Marshaling and Unmarshaling the list.")
	marshaled, _ := json.Marshal(TodoList)
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
	app_utils.SaveList(&TodoList, filepath.Join(projectpath, projectname))
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
