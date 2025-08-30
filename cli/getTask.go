package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var getTaskCli = &cobra.Command{
	Use:   "get",
	Short: "Get an existing task",
	Long:  `Get an existing task from task tracker`,
	Run: func(cli *cobra.Command, args []string) {
		fmt.Println(getTask(args[0]))
	},
}

func init() {
	rootCli.AddCommand(getTaskCli)
}
