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
	Statuses []Status // Now accepts a list, e.g., []Status{Todo, InProgress, Done}
}

func main() {
	options := printMenu()

	// ---> THIS IS THE OUTER LOOP <---
	// It keeps the whole program running until we hit 'break' at the bottom

	for {
		choice := 0

		// 1. Print the menu

		for i, option := range options {
			fmt.Printf("%d. %s \n", i+1, option)
		}

		// 2. Validate user input

		for {
			fmt.Print("Pick a choice: ")
			count, err := fmt.Scanln(&choice)

			if err == nil && count > 0 && choice <= len(options) {
				fmt.Println()
				break
			}

			fmt.Println("Invalid choice. Please try again.")
			fmt.Println()
			bufio.NewReader(os.Stdin).ReadString('\n')
		}

		// 3. Command Map Routing

		commands := map[int]Command{
			1: addTask,
			2: updateTask,
			3: deleteTask,
			4: listAllTasks,
			5: listCompletedTasks,
			6: listUncompletedTasks,
			7: listTasksInProgress,
		}

		action, exists := commands[choice]

		if exists {
			action()

			var cont string
			fmt.Print("\nDo you want to continue? (y/n): ")
			fmt.Scanln(&cont)

			if strings.ToLower(cont) != "y" {
				fmt.Println("Goodbye!")
				break
			}
			fmt.Println()

		}

	}
}
