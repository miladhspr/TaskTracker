// cmd/add.go
package cmd

import (
	"TaskTracker/tasks"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var ListCmd = &cobra.Command{
	Use:   "list [status]",
	Short: "list all tasks and filter by status",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

func init() {
	rootCmd.AddCommand(ListCmd)
}
