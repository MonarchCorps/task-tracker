package main

import "fmt"

func printMenu() []string {

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

	return options
}
