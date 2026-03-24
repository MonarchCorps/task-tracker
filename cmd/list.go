package cmd

import (
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		task.ListAllTasks()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
