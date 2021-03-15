package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Brotchu/tasks/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "to add a task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		taskid, err := db.CreateTask(task)
		if err != nil {
			fmt.Printf("[Err] \n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("task %d \"%s\" added\n", taskid, task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
