package cmd

import (
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

var listCompletedCmd = &cobra.Command{
	Use:   "list-completed",
	Short: "List completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		task.ListTasksByStatus([]task.Status{task.Done})
	},
}

func init() {
	rootCmd.AddCommand(listCompletedCmd)
}
