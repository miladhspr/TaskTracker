package tasks

import (
	"encoding/json"
	"fmt"
	"os"
)

const taskFile = "tasks.json"

func LoadTasks() ([]Task, error) {
	file, err := os.OpenFile(taskFile, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to open tasks.json: %v", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %v", err)
	}

	var tasks []Task
	if info.Size() > 0 {
		if err := json.NewDecoder(file).Decode(&tasks); err != nil {
			return nil, fmt.Errorf("failed to decode tasks.json: %v", err)
		}
	} else {
		tasks = []Task{}
	}

	return tasks, nil
}

func StoreTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks: %v", err)
	}
	if err := os.WriteFile(taskFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write tasks to file: %v", err)
	}
	return nil
}
