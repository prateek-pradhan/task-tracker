package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var deleteTaskCli = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing task",
	Long:  `Delete an existing task in task tracker`,
	Run: func(cli *cobra.Command, args []string) {
		fmt.Println(deleteTask(args[0]))
	},
}

func init() {
	rootCli.AddCommand(deleteTaskCli)
}
