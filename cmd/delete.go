// cmd/add.go
package cmd

import (
	"TaskTracker/tasks"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "delete a task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskId := os.Args[2]
		if err := tasks.Delete(taskId); err != nil {
			fmt.Println("Error deleting task:", err)
			os.Exit(1)
		}
		fmt.Println("Task deleted successfully.")
	},
}

func init() {
	rootCmd.AddCommand(DeleteCmd)
}
