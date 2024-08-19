package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int    `json:id`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func main() {
	errorCheck("At least one argument is needed", 2)

	action := os.Args[1]
	switch action {
	case "add":
		errorCheck("Provide a task description", 3)
		taskDesc := os.Args[2]
		if err := add(taskDesc); err != nil {
			fmt.Println("Error adding task:", err)
			os.Exit(1)
		}
		fmt.Println("Task added successfully.")
	case "update":
		errorCheck("Please provide a task ID and description.", 4)
		taskId := os.Args[2]
		newDescription := os.Args[3]
		if _, err := update(taskId, newDescription); err != nil {
			fmt.Println("Error updating task:", err)
			os.Exit(1)
		}
		fmt.Println("Task description updated successfully.")
	case "delete":
		errorCheck("Please provide a task ID.", 3)
		taskId := os.Args[2]
		if err := delete(taskId); err != nil {
			fmt.Println("Error deleting task:", err)
			os.Exit(1)
		}
		fmt.Println("Task deleted successfully.")
	case "list":
		var status string
		if len(os.Args) > 2 {
			status = os.Args[2]
		}
		tasks, err := all(status)
		if err != nil {
			fmt.Println("Error listing tasks:", err)
			os.Exit(1)
		}
		for _, task := range tasks {
			fmt.Printf("ID: %d, Description: %s, Status: %s\n", task.ID, task.Description, task.Status)
		}
	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}

func errorCheck(msg string, lenOfArgs int) {
	if len(os.Args) < lenOfArgs {
		fmt.Println(msg)
		os.Exit(1)
	}
}

func add(desc string) error {
	tasks, err := loadTasks()
	if err != nil {
		return fmt.Errorf("unable to add task: %v", err)
	}

	newTask := Task{
		ID:          len(tasks) + 1,
		Description: desc,
		Status:      "todo",
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}
	tasks = append(tasks, newTask)
	return storeTasks(tasks)
}

func update(taskId, content string) ([]Task, error) {
	tasks, err := loadTasks()
	if err != nil {
		return nil, fmt.Errorf("unable to update task: %v", err)
	}

	id, err := strconv.Atoi(taskId)
	if err != nil {
		return nil, fmt.Errorf("invalid task ID: %v", err)
	}

	task, err := findTaskByID(tasks, id)
	if err != nil {
		return nil, err
	}

	task.Description = content
	task.UpdatedAt = time.Now().String()

	for i, t := range tasks {
		if t.ID == id {
			tasks[i] = task
			break
		}
	}

	if err := storeTasks(tasks); err != nil {
		return nil, fmt.Errorf("unable to save updated tasks: %v", err)
	}

	return tasks, nil
}

func delete(taskId string) error {
	tasks, err := loadTasks()
	if err != nil {
		return fmt.Errorf("unable to delete task: %v", err)
	}

	id, err := strconv.Atoi(taskId)
	if err != nil {
		return fmt.Errorf("invalid task ID: %v", err)
	}

	index := -1
	for key, task := range tasks {
		if task.ID == id {
			index = key
			break
		}
	}

	tasks = append(tasks[:index], tasks[index+1:]...)

	return storeTasks(tasks)
}

func all(status string) ([]Task, error) {
	tasks, err := loadTasks()
	if err != nil {
		return nil, fmt.Errorf("unable to load tasks: %v", err)
	}

	if status != "" {
		tasks = filterTasksByStatus(tasks, status)
	}

	return tasks, nil
}

func findTaskByID(tasks []Task, id int) (Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return Task{}, fmt.Errorf("task with ID %d not found", id)
}

func loadTasks() ([]Task, error) {
	file, err := os.OpenFile("tasks.json", os.O_CREATE|os.O_RDWR, 0755)
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

func storeTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks: %v", err)
	}
	if err := os.WriteFile("tasks.json", data, 0644); err != nil {
		return fmt.Errorf("failed to write tasks to file: %v", err)
	}
	return nil
}

func filterTasksByStatus(tasks []Task, status string) []Task {
	var filteredTasks []Task
	for _, task := range tasks {
		if task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}
