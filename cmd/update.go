package cmd

import (
	"fmt"
	"strconv"
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

var updateStatus string
var description string

var updateCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update a task",
	Args:  cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		parsedStatus, ok := task.ParseStatus(updateStatus)
		if !ok {
			fmt.Println("Invalid status")
			return
		}

		task.UpdateTask(id, description, parsedStatus)
	},
}

func init() {
	updateCmd.Flags().StringVarP(&updateStatus, "status", "s", "todo", "New status")
	updateCmd.Flags().StringVarP(&description, "desc", "d", "", "New description")

	rootCmd.AddCommand(updateCmd)
}
