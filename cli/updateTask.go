package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var updateTaskCli = &cobra.Command{
	Use:   "update",
	Short: "Update an existing task",
	Long:  `Update an existing task in task tracker`,
	Run: func(cli *cobra.Command, args []string) {
		fmt.Println(updateTask(args[0], args[1]))
	},
}

func init() {
	rootCli.AddCommand(updateTaskCli)
}
