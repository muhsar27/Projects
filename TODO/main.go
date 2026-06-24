package main

import (
	"fmt"
	"os"
)

type Todo struct {
	Task string
	Done bool
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Print("Usage: todo command [arguments]")
		return
	}

	var taskList []Todo
	command := args[1]
	taskName := args[2]

	if command == "add" {
		taskList.Task = taskName
	}

}
