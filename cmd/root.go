// cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "TaskTracker",
	Short: "TaskTracker is a CLI for managing tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TaskTracker CLI")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
