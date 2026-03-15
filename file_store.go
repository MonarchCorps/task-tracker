package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const filePath = "data/tasks.json"

func saveTasks(tasks []Task) error {
	// 1. Get the directory name from the path ("data")
	dir := filepath.Dir(filePath)

	// 2. Create the folder if it doesn't exist.
	// 0755 is the standard permission for "Read/Write"
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}

func loadTasks() ([]Task, error) {
	file, err := os.Open(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	defer file.Close()

	var tasks []Task

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)

	return tasks, err
}

func filterTasks(filter TaskFilter) ([]Task, error) {
	tasks, err := loadTasks()
	if err != nil {
		return nil, err
	}

	var filteredTasks []Task
	for _, task := range tasks {
		// Check if the task's status is in our "allowed" filter list
		for _, allowedStatus := range filter.Statuses {
			if task.Status == allowedStatus {
				filteredTasks = append(filteredTasks, task)
				break // Found a match, move to the next task
			}
		}
	}

	return filteredTasks, nil
}

func printFormattedTasks(tasks []Task) {
	for i, task := range tasks {
		fmt.Printf("%d. %s, status = %s \n", i+1, task.Description, task.Status)
	}
}
