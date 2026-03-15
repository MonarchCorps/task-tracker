package main

import (
	"bufio"
	"fmt"
	"os"
)

func addTask() {
	// Create a scanner that reads from the keyboard (os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a description: ")
	scanner.Scan() // Reads the whole line until you press Enter
	description := scanner.Text()

	var status Status

	for {
		fmt.Print("Enter a status (Todo, In-Progress, Done): ")
		scanner.Scan()
		input := scanner.Text()

		parsedStatus, ok := parseStatus(input)

		if ok {
			status = parsedStatus
			break
		}

		fmt.Println("Invalid status. Try again.")

	}

	loadedTasks, _ := loadTasks()

	newTask := Task{
		ID:          len(loadedTasks) + 1,
		Description: description,
		Status:      status,
	}

	tasks := append(loadedTasks, newTask)

	err := saveTasks(tasks)

	if err != nil {
		fmt.Print("Error saving tasks: ", err)
		return
	}

	fmt.Println("Tasks saved.")
}

func updateTask() {
	choice := 0
	input := ""
	var (
		description string
		status      Status
	)

	tasks, err := loadTasks()
	if err != nil {
		fmt.Print("Error loading tasks: ", err)
		return
	}

	fmt.Println("Here are your lists of tasks: ")
	printFormattedTasks(tasks)

	for {
		fmt.Println()
		fmt.Print("Select a task to update: ")

		_, choiceError := fmt.Scanln(&choice)

		if choiceError != nil {
			fmt.Print("Error scanning tasks: ", choiceError)
			return
		} else if choice < 1 || choice > len(tasks) {
			fmt.Print("Invalid task count. Try again.")
		} else {
			break
		}
	}

	fmt.Print("Enter a description: ")
	fmt.Scanln(&description)

	fmt.Print("Enter a status (Todo, In-Progress, Done): ")
	fmt.Scanln(&input)
	parsedStatus, ok := parseStatus(input)

	if ok {
		status = parsedStatus

		tasks[choice-1] = Task{
			ID:          tasks[choice-1].ID,
			Description: description,
			Status:      status,
		}

		saveTasks(tasks)

		fmt.Println("Updated task: ", tasks)

	}

}

func deleteTask() {
	choice := 0

	tasks, err := loadTasks()
	if err != nil {
		fmt.Print("Error loading tasks: ", err)
		return
	}

	fmt.Println("Here are your lists of tasks: ")
	printFormattedTasks(tasks)

	for {
		fmt.Println()
		fmt.Print("Select a task to delete: ")

		_, choiceError := fmt.Scanln(&choice)

		if choiceError != nil {
			fmt.Print("Error scanning tasks: ", choiceError)
			return
		} else if choice < 1 || choice > len(tasks) {
			fmt.Print("Invalid task count. Try again.")
		} else {
			break
		}
	}

	var filteredTasks []Task
	//
	for i, task := range tasks {
		if i != choice-1 {
			filteredTasks = append(filteredTasks, task)
		}
	}

	savedToFileError := saveTasks(filteredTasks)
	if savedToFileError != nil {
		fmt.Print("Error saving tasks: ", savedToFileError)
		return
	}

	fmt.Println("Tasks saved.")
}

func listAllTasks() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Print("Error loading tasks: ", err)
	}

	fmt.Println(tasks)
}

func listCompletedTasks() {
	filter := TaskFilter{
		Statuses: []Status{Done},
	}
	tasks, err := filterTasks(filter)
	if err != nil {
		fmt.Print("Error finding tasks: ", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println(tasks)
}

func listUncompletedTasks() {
	filter := TaskFilter{
		Statuses: []Status{InProgress, Todo},
	}
	tasks, err := filterTasks(filter)
	if err != nil {
		fmt.Print("Error finding tasks: ", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println(tasks)
}

func listTasksInProgress() {
	filter := TaskFilter{
		Statuses: []Status{InProgress},
	}
	tasks, err := filterTasks(filter)
	if err != nil {
		fmt.Print("Error finding tasks: ", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println(tasks)
}
