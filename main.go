package main

import (
	"TaskTracker/tasks"
	"fmt"
	"os"
)

func main() {
	tasks.ErrorArgsCheck("At least one argument is needed", 2)

	action := os.Args[1]
	switch action {
	case "add":
		tasks.ErrorArgsCheck("Provide a task description", 3)
		taskDesc := os.Args[2]
		if err := tasks.Add(taskDesc); err != nil {
			fmt.Println("Error adding task:", err)
			os.Exit(1)
		}
		fmt.Println("Task added successfully.")
	case "update":
		tasks.ErrorArgsCheck("Please provide a task ID and description.", 4)
		taskId := os.Args[2]
		newDescription := os.Args[3]
		if _, err := tasks.Update(taskId, newDescription); err != nil {
			fmt.Println("Error updating task:", err)
			os.Exit(1)
		}
		fmt.Println("Task description updated successfully.")
	case "delete":
		tasks.ErrorArgsCheck("Please provide a task ID.", 3)
		taskId := os.Args[2]
		if err := tasks.Delete(taskId); err != nil {
			fmt.Println("Error deleting task:", err)
			os.Exit(1)
		}
		fmt.Println("Task deleted successfully.")
	case "list":
		var status string
		if len(os.Args) > 2 {
			status = os.Args[2]
		}
		tasks, err := tasks.All(status)
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
