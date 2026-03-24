package cmd

import (
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

func main() {}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	Run: func(cmd *cobra.Command, args []string) {
		task.UpdateTask()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
