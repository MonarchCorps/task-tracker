package cmd

import (
	"fmt"
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

var filterStatus string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",

	Run: func(cmd *cobra.Command, args []string) {
		if filterStatus == "" {
			task.ListAllTasks()
			return
		}

		parsed, ok := task.ParseStatus(filterStatus)
		if !ok {
			fmt.Println("Invalid status")
			return
		}

		task.ListTasksByStatus([]task.Status{parsed})
	},
}

func init() {
	listCmd.Flags().StringVarP(&filterStatus, "status", "s", "", "Filter by status")
	rootCmd.AddCommand(listCmd)
}
