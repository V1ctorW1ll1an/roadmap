package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func ReadJson(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var taskList TaskList
	if err := json.Unmarshal(file, &taskList); err != nil {
		panic(err)
	}

	for _, task := range taskList.Tasks {
		fmt.Printf("Id: %d\nDescription: %s\nStatus: %t\nCreatedAt: %s\nUpdatedAt: %s\n\n",
			task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	}
}

// -- tasks

// Add tasks
func AddTask(parameter string) {
}

// Update tasks
// Delete tasks
// Mark a task as in progress or done
// List all tasks
// List all tasks that are done
// List all tasks that are not done
// List all tasks that are in progress
func ShowHelp() {
	helpMessage := `
task-cli - Just a simple tool for learning Go

Usage:
  task-cli <commands> [arguments]

Available Commands:
  help          Displays this help message
  version       Shows the program version

Examples:
  task-cli help
  task-cli version
`
	fmt.Println(helpMessage)
}

func ShowVersion() {
	fmt.Println("task-cli v1.0.0")
}
func main() {

	if len(os.Args) < 2 {
		ShowHelp()
		os.Exit(1)
	}
	var parameter string

	command := os.Args[1]

	if len(os.Args) >= 3 {
		parameter = os.Args[2]
	}

	_ = parameter

	switch command {
	case "help":
		ShowHelp()
	case "version":
		ShowVersion()
	case "add":
		ReadJson("tasks.json")
	default:
		fmt.Printf("unknow command: %s\n", command)
		ShowHelp()
		os.Exit(1)
	}
}
