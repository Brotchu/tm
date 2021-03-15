package cmd

import (
	"fmt"
	"os"

	"github.com/Brotchu/tasks/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks ",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("[Err] ", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You do not have anything to do :)")
			os.Exit(0)
		}
		fmt.Println("Here is what you have to do: ")
		for _, task := range tasks {
			fmt.Printf("%d. %s\n", task.Key, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
