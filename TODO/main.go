package main

import (
	"fmt"
	"os"
)

type Todo struct {
	Entry string `json:entry`
	Done  bool   `json:done`
}

// commands for my todo list:
// add(add a task), list(enumerate the tasks and status),
// done(correct the status of one task to done), delete(delete a task)
func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		return
	}
	var TaskList []Todo
	switch args[0] {
	case "add":
		task := args[1]
		var status bool
		if args[2] == "Done" {
			status = true
		} else {
			status = false
		}

		newTodo := Todo{Entry: task, Done: status}
		TaskList = append(TaskList, newTodo)

	default:
		fmt.Print("Fuck off \n")
	}

	fmt.Println(TaskList)

}
