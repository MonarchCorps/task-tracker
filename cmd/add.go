package cmd

import (
	"fmt"
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

var status string

var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]

		parsedStatus, ok := task.ParseStatus(status)
		if !ok {
			fmt.Println("Invalid status. Use: todo, in-progress, done")
			return
		}

		task.AddTaskWithParams(description, parsedStatus)
	},
}

func init() {
	addCmd.Flags().StringVarP(&status, "status", "s", "todo", "Task status")
	rootCmd.AddCommand(addCmd)
}
