package cmd

import (
	"TaskTracker/tasks"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var AddCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		desc := args[0]
		if err := tasks.Add(desc); err != nil {
			fmt.Println("Error adding task:", err)
			os.Exit(1)
		}
		fmt.Println("Task added successfully.")
	},
}

func init() {
	rootCmd.AddCommand(AddCmd)
}
