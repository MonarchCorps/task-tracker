package cmd

import (
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

var listInProgressCmd = &cobra.Command{
	Use:   "list-inprogress",
	Short: "List tasks in progress",
	Run: func(cmd *cobra.Command, args []string) {
		task.ListTasksByStatus([]task.Status{
			task.InProgress,
		})
	},
}

func init() {
	rootCmd.AddCommand(listInProgressCmd)
}
