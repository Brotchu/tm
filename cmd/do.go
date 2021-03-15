package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Brotchu/tasks/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "mark a task completed",
	Run: func(cmd *cobra.Command, args []string) {
		var id []int
		for _, arg := range args {
			taskId, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Unable to parse %s\n", arg)
			} else {
				id = append(id, taskId)
			}
		}

		if len(id) == 0 {
			fmt.Printf("provide task id to mark done.\n")
			os.Exit(0)
		}

		for _, taskId := range id {
			err := db.DeleteTask(taskId)
			if err != nil {
				fmt.Printf("Could not do %d\n", taskId)
			} else {
				fmt.Printf("Task %d done.. \n", taskId)
			}
		}
		// fmt.Printf("tasks %v done\n", id)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
