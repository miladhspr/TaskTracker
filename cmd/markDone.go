// cmd/add.go
package cmd

import (
	"TaskTracker/tasks"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var MarkDoneCmd = &cobra.Command{
	Use:   "mark-done [id]",
	Short: "update a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newStatus := "done"
		taskId := os.Args[2]
		if _, err := tasks.Update(taskId, "status", newStatus); err != nil {
			fmt.Println("Error updating task:", err)
			os.Exit(1)
		}
		fmt.Println("Task marked as done")
	},
}

func init() {
	rootCmd.AddCommand(MarkDoneCmd)
}
