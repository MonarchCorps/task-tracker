package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command func()

type Status string

const (
	Todo       Status = "todo"
	InProgress Status = "in-progress"
	Done       Status = "done"
)

type Task struct {
	ID          int
	Description string
	Status      Status
}

type TaskFilter struct {
	Statuses []Status
}

var reader = bufio.NewScanner(os.Stdin)

func main() {
	commands := map[int]Command{
		1: addTask,
		2: updateTask,
		3: deleteTask,
		4: listAllTasks,
		5: func() { listTasksByStatus([]Status{Done}) },
		6: func() { listTasksByStatus([]Status{Todo, InProgress}) },
		7: func() { listTasksByStatus([]Status{InProgress}) },
	}

	for {
		printMenu()

		choice := getIntInput("Pick a choice: ")

		action, exists := commands[choice]
		if !exists {
			fmt.Println("Invalid choice.")
			continue
		}

		action()

		cont := getStringInput("\nDo you want to continue? (y/n): ")
		if strings.ToLower(cont) != "y" {
			fmt.Println("Goodbye!")
			break
		}
	}
}
