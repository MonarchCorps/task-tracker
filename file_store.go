package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const filePath = "data/tasks.json"

func saveTasks(tasks []Task) error {
	dir := filepath.Dir(filePath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(tasks)
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
	err = json.NewDecoder(file).Decode(&tasks)

	return tasks, err
}

func filterTasks(filter TaskFilter) ([]Task, error) {
	tasks, err := loadTasks()
	if err != nil {
		return nil, err
	}

	var result []Task
	for _, task := range tasks {
		for _, s := range filter.Statuses {
			if task.Status == s {
				result = append(result, task)
				break
			}
		}
	}

	return result, nil
}
