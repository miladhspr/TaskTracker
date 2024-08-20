// cmd/add.go
package cmd

import (
	"TaskTracker/tasks"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var UpdateCmd = &cobra.Command{
	Use:   "update [id] [description]",
	Short: "update a new task",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		taskId := os.Args[2]
		newDescription := os.Args[3]
		if _, err := tasks.Update(taskId, "description", newDescription); err != nil {
			fmt.Println("Error updating task:", err)
			os.Exit(1)
		}
		fmt.Println("Task description updated successfully.")
	},
}

func init() {
	rootCmd.AddCommand(UpdateCmd)
}
