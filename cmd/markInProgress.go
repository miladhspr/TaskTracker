// cmd/add.go
package cmd

import (
	"TaskTracker/tasks"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var MarkInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress [id]",
	Short: "mark task as in-progress",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newStatus := "in-progress"
		taskId := os.Args[2]
		if _, err := tasks.Update(taskId, "status", newStatus); err != nil {
			fmt.Println("Error updating task:", err)
			os.Exit(1)
		}
		fmt.Println("Task marked as in-progress")
	},
}

func init() {
	rootCmd.AddCommand(MarkInProgressCmd)
}
