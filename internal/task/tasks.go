package task

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

func AddTask() {
	description := GetStringInput("Enter description: ")
	if strings.TrimSpace(description) == "" {
		fmt.Println("Description cannot be empty")
		return
	}

	var status Status
	for {
		input := GetStringInput("Enter status (todo, in-progress, done): ")
		parsed, ok := ParseStatus(input)
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

func UpdateTask(id int, description string, status Status) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	found := false

	for i, t := range tasks {
		if t.ID == id {
			if description != "" {
				tasks[i].Description = description
			}
			tasks[i].Status = status
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Task not found.")
		return
	}

	if err := saveTasks(tasks); err != nil {
		fmt.Println("Error saving:", err)
		return
	}

	fmt.Println("✅ Task updated.")
}

func DeleteTask(id int) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	var updated []Task
	found := false

	for _, t := range tasks {
		if t.ID != id {
			updated = append(updated, t)
		} else {
			found = true
		}
	}

	if !found {
		fmt.Println("Task not found.")
		return
	}

	if err := saveTasks(updated); err != nil {
		fmt.Println("Error saving:", err)
		return
	}

	fmt.Println("🗑️ Task deleted.")
}

func ListAllTasks() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No task found.")
		return
	}

	printFormattedTasks(tasks)
}

func ListTasksByStatus(statuses []Status) {
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

func AddTaskWithParams(description string, status Status) {
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
