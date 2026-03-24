package cmd

import (
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

var listUncompletedCmd = &cobra.Command{
	Use:   "list-uncompleted",
	Short: "List uncompleted tasks",
	Run: func(cmd *cobra.Command, args []string) {
		task.ListTasksByStatus([]task.Status{
			task.Todo,
			task.InProgress,
		})
	},
}

func init() {
	rootCmd.AddCommand(listUncompletedCmd)
}
