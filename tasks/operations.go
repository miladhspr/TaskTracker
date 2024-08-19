package tasks

import (
	"fmt"
	"strconv"
	"time"
)

func Add(desc string) error {
	tasks, err := LoadTasks()
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
	return StoreTasks(tasks)
}

func Update(taskId, column string, content string) ([]Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return nil, fmt.Errorf("unable to update task: %v", err)
	}

	id, err := strconv.Atoi(taskId)
	if err != nil {
		return nil, fmt.Errorf("invalid task ID: %v", err)
	}

	task, err := FindTaskByID(tasks, id)
	if err != nil {
		return nil, err
	}

	switch column {
	case "status":
		task.Status = content
	case "description":
		task.Description = content
	default:
		return nil, fmt.Errorf("Invalid column name: %s", column)
	}
	task.UpdatedAt = time.Now().String()

	for i, t := range tasks {
		if t.ID == id {
			tasks[i] = task
			break
		}
	}

	if err := StoreTasks(tasks); err != nil {
		return nil, fmt.Errorf("unable to save updated tasks: %v", err)
	}

	return tasks, nil
}

func Delete(taskId string) error {
	tasks, err := LoadTasks()
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

	return StoreTasks(tasks)
}

func All(status string) ([]Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return nil, fmt.Errorf("unable to load tasks: %v", err)
	}

	if status != "" {
		tasks = FilterTasksByStatus(tasks, status)
	}

	return tasks, nil
}

func FindTaskByID(tasks []Task, id int) (Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return Task{}, fmt.Errorf("task with ID %d not found", id)
}

func FilterTasksByStatus(tasks []Task, status string) []Task {
	var filteredTasks []Task
	for _, task := range tasks {
		if task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}
