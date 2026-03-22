package main

import (
	"fmt"
	"strings"
)

func getNextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func addTask() {
	description := getStringInput("Enter description: ")
	if strings.TrimSpace(description) == "" {
		fmt.Println("Description cannot be empty")
		return
	}

	var status Status
	for {
		input := getStringInput("Enter status (todo, in-progress, done): ")
		parsed, ok := parseStatus(input)
		if ok {
			status = parsed
			break
		}
		fmt.Println("Invalid status. Try again.")
	}

	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	newTask := Task{
		ID:          getNextID(tasks),
		Description: description,
		Status:      status,
	}

	tasks = append(tasks, newTask)

	if err := saveTasks(tasks); err != nil {
		fmt.Println("Error saving:", err)
		return
	}

	fmt.Println("✅ Task added.")
}

func updateTask() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	printFormattedTasks(tasks)

	choice := getIntInput("Select task to update: ")
	if choice < 1 || choice > len(tasks) {
		fmt.Println("Invalid selection")
		return
	}

	description := getStringInput("Enter new description: ")
	if strings.TrimSpace(description) == "" {
		fmt.Println("Description cannot be empty")
		return
	}

	var status Status
	for {
		input := getStringInput("Enter new status: ")
		parsed, ok := parseStatus(input)
		if ok {
			status = parsed
			break
		}
		fmt.Println("Invalid status. Try again.")
	}

	tasks[choice-1].Description = description
	tasks[choice-1].Status = status

	if err := saveTasks(tasks); err != nil {
		fmt.Println("Error saving:", err)
		return
	}

	fmt.Println("✅ Task updated.")
}

func deleteTask() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	printFormattedTasks(tasks)

	choice := getIntInput("Select task to delete: ")
	if choice < 1 || choice > len(tasks) {
		fmt.Println("Invalid selection")
		return
	}

	tasks = append(tasks[:choice-1], tasks[choice:]...)

	if err := saveTasks(tasks); err != nil {
		fmt.Println("Error saving:", err)
		return
	}

	fmt.Println("🗑️ Task deleted.")
}

func listAllTasks() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	printFormattedTasks(tasks)
}

func listTasksByStatus(statuses []Status) {
	filter := TaskFilter{Statuses: statuses}

	tasks, err := filterTasks(filter)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	printFormattedTasks(tasks)
}
