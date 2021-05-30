package cmd

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
	"strings"

	"github.com/Brotchu/tasks/db"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "to publish tasks to epaper. arguments - task numbers",
	Run: func(cmd *cobra.Command, args []string) {
		var id []int
		for _, arg := range args {
			taskId, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("unable to parse %s\n", arg)
			} else {
				id = append(id, taskId)
			}
		}

		if len(id) == 0 {
			fmt.Printf("provide task ids to publish\n")
			os.Exit(0)
		}

		var resString strings.Builder
		for _, taskId := range id {
			t, err := db.GetTasks(taskId)
			if err != nil {
				fmt.Printf("Couldnt find task %d\n", taskId)
			} else {
				resString.Write([]byte(strconv.Itoa(t.Key) + "-"))
				resString.WriteString(t.Value)
				resString.WriteString("\n")
			}
		}
		res := resString.String()
		res = res[:len(res)-1]
		fmt.Println(res)
		pi_addr := os.Getenv("PI_ADDRESS")

		client, err := rpc.DialHTTP("tcp", pi_addr+":4040")
		if err != nil {
			log.Printf("couldnt connect to %s at 4040", pi_addr)
			log.Printf("check env variable PI_ADDRESS (set to local PI address)")
			os.Exit(1)
		}

		var reply string
		err = client.Call("Display.EPrint", res, &reply)
		if err != nil {
			log.Println("Error in RPC ", err.Error())
			os.Exit(1)
		}
		// log.Printf(reply)
	},
}

func init() {
	RootCmd.AddCommand(publishCmd)
}
