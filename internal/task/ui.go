package task

import "fmt"

func printMenu() {
	fmt.Println("Hello there 👋🏻! Welcome to Task Tracker")
	fmt.Println("Choose from the list of options to get started:")
	fmt.Println()

	options := []string{
		"Add a new task",
		"Update a task",
		"Delete a task",
		"List all tasks",
		"List completed tasks",
		"List uncompleted tasks",
		"List tasks in progress",
	}

	for i, option := range options {
		fmt.Printf("[%d] %s\n", i+1, option)
	}
}

func printFormattedTasks(tasks []Task) {
	for i, task := range tasks {
		fmt.Printf("%d. %s, status = %s \n", i+1, task.Description, task.Status)
	}
}
