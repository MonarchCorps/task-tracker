package cmd

import (
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		task.AddTask()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
