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

	instruction := args[1]

	var taskList []string

	if instruction == "add" {
		task := args[2]
		taskList = append(taskList, task)
	}

	if instruction == "list" {
		fmt.Println(taskList)
	}

}
