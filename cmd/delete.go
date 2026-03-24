package cmd

import (
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		task.DeleteTask()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
